#!/usr/bin/python3

import argparse
import os
import random as r
import errno


def generate():
    parser = argparse.ArgumentParser()
    parser.add_argument('-n', action='store', dest='numfiles', default=80, help='Indica quanti file generare')
    parser.add_argument('-d', '--dir', action='append', dest='dirs', default=[],
                        help='Specifica cartelle dove generare file')
    result = parser.parse_args()
    if not result.dirs:
        exit(0)
    for reldir in result.dirs:
        absdir = os.path.abspath(reldir)
        for num in range(int(result.numfiles)):
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
    generate()
