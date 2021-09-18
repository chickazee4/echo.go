#!/bin/sh
for x in `go tool dist list`; do
    if [ ! -d dist ]; then
        mkdir dist
    fi
    GOOS=`echo ${x} | cut -d / -f 1`
    GOARCH=`echo ${x} | cut -d / -f 2`
    if [ $GOOS == "windows" ]; then
        go build -o dist/win_${GOARCH}-echo.exe src/main.go
    else
        go build -o dist/${GOOS}_${GOARCH}-echo src/main.go
    fi
done