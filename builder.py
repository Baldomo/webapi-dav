#!/usr/bin/env python3
import click
import errno
import os
import platform
import queue as q
import random as r
import shutil as sh
import subprocess
from typing import Any, Dict, List, Optional
from threading import Thread


out_folder: str = 'build/'
goreleaser_token: str = '7c71dfc5c3255c918b8ace1b9cc024b1a16acc50'
filenames: Dict[str, str] = {
    'windows': 'webapi-dav-windows_amd64.exe',
    'windows-nogui': 'webapi-dav-windows_amd64.exe',
    'linux': 'webapi-dav-linux_amd64',
    'darwin': 'webapi-dav-mac_amd64'
}
package: str = './cmd/webapi'
commands: Dict[str, str] = {
    'windows':
        'go build -ldflags="-s -w" -o ' + out_folder + filenames['windows'] + ' ' + package,
    'windows-nogui':
        'go build -ldflags="-s -w -H windowsgui" -o ' + out_folder + filenames['windows'] + ' ' + package,
    'linux':
        'go build -ldflags="-s -w" -o ' + out_folder + filenames['linux'] + ' ' + package,
    'darwin':
        'go build -ldflags="-s -w" -o ' + out_folder + filenames['darwin'] + ' ' + package
}


class AliasedGroup(click.Group):
    def get_command(self, ctx: click.Context, cmd_name: Any) -> Optional[Any]:
        rv = click.Group.get_command(self, ctx, cmd_name)
        if rv is not None:
            return rv
        matches = [x for x in self.list_commands(ctx)
                   if x.startswith(cmd_name)]
        if not matches:
            return None
        elif len(matches) == 1:
            return click.Group.get_command(self, ctx, matches[0])
        ctx.fail('Too many matches: %s' % ', '.join(sorted(matches)))


@click.group(chain=True, cls=AliasedGroup)
def cli() -> None:
    click.echo('')


@cli.command()
@click.argument('names', nargs=-1)
def build(names: List[str]) -> None:
    if not os.path.exists(out_folder):
        os.mkdir(out_folder)
    for name in names:
        if name in commands.keys():
            Builder({name: commands[name]}).start()
        elif name == 'all':
            Builder(commands).start()
        elif name == 'this':
            Builder({name: commands[current_platform]}).start()


@cli.command()
def pack() -> None:
    if not os.path.exists(out_folder):
        print('folder {} not found'.format(out_folder))
        exit(1)
    subprocess.run('cd {} && upx --brute webapi-*'.format(out_folder), shell=True).check_returncode()


class Builder:
    def __init__(self, args_dict: Dict[str, str] = None) -> None:
        if len(args_dict) == 0 or not args_dict:
            return
        self.args = args_dict
        self.format_commands()
        print(self.args if self.args else 'none')

        self.queue = q.Queue()
        for target, arg in self.args.items():
            self.queue.put((target, arg))

    def start(self) -> None:
        threads = [
            Thread(target=Builder.proc, args=(self.queue,)) for _ in self.args
        ]
        for t in threads:
            t.setDaemon(True)
            t.start()
        self.queue.join()
        while threads:
            threads.pop().join()

    @staticmethod
    def proc(queue: q.Queue) -> None:
        while 1:
            try:
                arg = queue.get(block=False)
            except q.Empty:
                break
            # print('Building {}'.format(filenames[arg[0]]))
            subprocess.Popen(arg[1], shell=True).wait()
            queue.task_done()
            break

    def format_commands(self) -> None:
        for target, cmd in self.args.items():
            self.args[target] = Builder.generate_env_flags(target) + cmd

    @staticmethod
    def generate_env_flags(target: str) -> str:
        if 'windows' in target:
            target = 'windows'
        elif 'linux' in target:
            target = 'linux'
        elif 'darwin' in target:
            target = 'darwin'
        elif 'this' in target:
            target = current_platform
        else:
            raise ValueError('only windows, linux and darwin are supported; use "this" to target current platform')

        if platform.system() == 'Windows':
            return 'set "GOOS={}" && set "GOARCH=amd64" && '.format(target)
        elif platform.system() == 'Linux' or platform.system() == 'Darwin':
            return 'env GOOS={} GOARCH=amd64 '.format(target)


# Utils
playground_folder: str = 'playground/'
current_platform: str = \
    'windows' if platform.system() == 'Windows' \
    else 'linux' if platform.system() == 'Linux' \
    else 'darwin'


@cli.command()
def clean() -> None:
    if os.path.exists(playground_folder):
        sh.rmtree(playground_folder)
    if os.path.exists(out_folder):
        sh.rmtree(out_folder)


def create_playground(exe: str, overwrite: bool) -> None:
    if os.path.exists(playground_folder):
        if overwrite:
            sh.rmtree(playground_folder)
        else:
            return
    sh.copytree('static', playground_folder + 'static')
    sh.copy2('config.toml', playground_folder)
    sh.copy2('orario.xml', playground_folder)
    sh.copy2(out_folder + exe, playground_folder)
    dummyfiles(100, [
        playground_folder + 'comunicati-studenti',
        playground_folder + 'comunicati-docenti',
        playground_folder + 'comunicati-genitori'
    ])


@cli.command()
@click.argument('target', nargs=1, type=str, required=False, default='this')
@click.option('--run', '-r', is_flag=True)
@click.option('--cleanup', '-c', is_flag=True)
@click.pass_context
def deploy(ctx: click.Context, target: str, run: bool, cleanup: bool) -> None:
    ctx.invoke(build, names=[target])
    if target == 'this':
        exe: str = filenames[current_platform]
    else:
        exe: str = filenames[target]
    create_playground(exe, True)
    if run:
        try:
            subprocess.run('cd {} && ./{}'.format(playground_folder, exe), shell=True).check_returncode()
        finally:
            if cleanup:
                ctx.invoke(clean)


@cli.command()
@click.option('--run', '-r', is_flag=True)
@click.option('--cleanup', '-c', is_flag=True)
@click.pass_context
def docker(ctx: click.Context, run: bool, cleanup: bool) -> None:
    ctx.invoke(build, names=['linux'])
    exe: str = filenames['linux']
    create_playground(exe, False)
    if cleanup:
        subprocess.run('docker rm webapi-dav', shell=True)
    subprocess.run('docker build -t webapi-dav .', shell=True).check_returncode()
    if run:
        subprocess.run('docker run -it webapi-dav /bin/bash', shell=True).check_returncode()


@cli.command()
def release() -> None:
    subprocess.run('set "GITHUB_TOKEN={}" goreleaser --rm-dist'.format(goreleaser_token), shell=True).check_returncode()


@cli.command()
def docker_clean() -> None:
    subprocess.run('docker rm webapi-dav && docker rmi $(docker ps -aq)', shell=True)


def dummyfiles(numfiles: int, dirs: List[str]):
    for reldir in dirs:
        absdir = os.path.abspath(reldir)
        for num in range(numfiles):
            filepath = os.path.join(absdir, '{}.txt'.format(num))
            if not os.path.exists(os.path.dirname(filepath)):
                try:
                    os.makedirs(os.path.dirname(filepath))
                except OSError as exc:  # Guard against race condition
                    if exc.errno != errno.EEXIST:
                        raise
            file = open(filepath, 'w+')
            file.truncate(r.randint(4e5, 4e6))


@cli.command()
def test():
    subprocess.run('go test -v -benchmem ./...', shell=True).check_returncode()


if __name__ == '__main__':
    cli()
