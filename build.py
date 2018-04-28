#!/usr/bin/python3
import click
import errno
import os
import platform
import queue as q
import random as r
import shutil as sh
import subprocess
from threading import Thread


# Builder
out_folder = 'build/'
filenames = {
    'windows': 'webapi-dav-windows_amd64.exe',
    'linux': 'webapi-dav-linux_amd64',
    'darwin': 'webapi-dav-mac_amd64'
}
commands = {
    'windows':
        'go build -ldflags="-s -w" -o ' + out_folder + filenames['windows'],
    'windows-nogui':
        'go build -ldflags="-s -w -H windowsgui" -o ' + out_folder + filenames['windows'],
    'linux':
        'go build -ldflags="-s -w" -o ' + out_folder + filenames['linux'],
    'darwin':
        'go build -ldflags="-s -w" -o ' + out_folder + filenames['darwin']
}


@click.group()
def cli():
    click.echo('')


@cli.command()
@click.argument('name', nargs=1)
def build(name):
    if name in commands.keys():
        Builder({name: commands[name]}).start()
    elif name == 'all':
        Builder(commands).start()


class Builder(object):
    def __init__(self, args_dict):
        if len(args_dict) == 0 or not args_dict:
            return
        self.args = self.format_commands(args_dict)
        print(self.args if self.args else 'none')

        self.queue = q.Queue()
        for arg in self.args:
            self.queue.put(arg)

    def start(self):
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
    def proc(queue):
        while 1:
            try:
                arg = queue.get(block=False)
            except q.Empty:
                break
            # print('Building {}'.format(arg.split()[5]))
            subprocess.Popen(arg, shell=True).wait()
            queue.task_done()
            break

    @staticmethod
    def format_commands(args_dict):
        fargs = []
        for target, cmd in args_dict.items():
            fargs.append(Builder.generate_env_flags(target) + cmd)
        return fargs

    @staticmethod
    def generate_env_flags(target):
        if 'windows' in target:
            target = 'windows'
        elif 'linux' in target:
            target = 'linux'
        elif 'darwin' in target:
            target = 'darwin'
        else:
            raise ValueError('only windows, linux and darwin are supported')

        if platform.system() == 'Windows':
            return 'set "GOOS={}" && set "GOARCH=amd64" && '.format(target)
        elif platform.system() == 'Linux' or platform.system() == 'Darwin':
            return 'env GOOS={} GOARCH=amd64 '.format(target)


# Utils
playground_folder = 'playground/'
current_platform = \
    'windows' if platform.system() == 'Windows' \
    else 'linux' if platform.system() == 'Linux' \
    else 'darwin'


@cli.command()
def clean():
    if os.path.exists(playground_folder):
        sh.rmtree(playground_folder)
    if os.path.exists(out_folder):
        sh.rmtree(out_folder)


@cli.command()
def deploy():
    if os.path.exists(playground_folder):
        sh.rmtree(playground_folder)
    sh.copytree('static', playground_folder + 'static')
    sh.copy2('config.toml', playground_folder)
    sh.copy2('orario.xml', playground_folder)
    exe = filenames[current_platform]
    sh.copy2(out_folder + exe, playground_folder)
    dummyfiles(100, [
        playground_folder + 'comunicati-studenti',
        playground_folder + 'comunicati-docenti',
        playground_folder + 'comunicati-genitori'
    ])
    subprocess.run('cd {} && {}'.format(playground_folder, exe), shell=True).check_returncode()


def dummyfiles(numfiles, dirs):
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


if __name__ == '__main__':
    cli()
