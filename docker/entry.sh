#!/bin/sh
set -e

if [ "$1" != "${1#-}" ]; then
    # if the first argument is an option like `--help` or `-h`
    exec botway "$@"
fi

case "$1" in
    docker | help | init | new | remove | tokens | version )
    exec botway "$@";;
esac

exec "$@"
