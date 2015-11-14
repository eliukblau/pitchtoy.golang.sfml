#!/bin/sh

echo "\nEjecutando el binario...\n"
cd src
go clean
go build -o "$$.tmp" && "./$$.tmp"
rm "$$.tmp"
cd ..
