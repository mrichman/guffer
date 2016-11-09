#!/bin/bash

# Usage: ./crosscompile.bash main.go guffer

# argument handling
test "$1" && target="$1" # .go file to build

if ! test "$target"
then
  echo "target file required"
  exit 1
fi

binary="" # default to default
test "$2" && binary="$2" # binary output

platforms="darwin/386 darwin/amd64 freebsd/386 freebsd/amd64 freebsd/arm linux/386 linux/amd64 linux/arm windows/386 windows/amd64"

for platform in ${platforms}
do
    split=(${platform//\// })
    goos=${split[0]}
    goarch=${split[1]}

    # ensure output file name
    output="$binary"
    test "$output" || output="$(basename $target | sed 's/\.go//')"

    # add exe to windows output
    [[ "windows" == "$goos" ]] && output="$output.exe"

    # set destination path for binary
    destination="$(dirname $target)/builds/$goos/$goarch/$output"

    echo "GOOS=$goos GOARCH=$goarch go build -x -o $destination $target"
    GOOS=$goos GOARCH=$goarch go build -x -o $destination $target
done
