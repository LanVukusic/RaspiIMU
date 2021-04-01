#!/bin/bash
OUT=${1:-"out"}
echo "Compiling to file: $OUT"
env GOOS=linux GOARCH=arm GOARM=6 go build -o "$OUT"
err="$?"
if (( $err != 0 )) ; then
    echo "Compilation error. Aborting."
    exit $err
fi
echo "Compilation done."
echo "Copying..."
cp "$OUT" /run/user/1000/gvfs/smb-share:server=192.168.1.41,share=dietpi
err="$?"
echo "error stat: $err"