#!/bin/sh

echo "\nEjecutando el binario...\n"
go clean
go build -o "$$.tmp" && "./$$.tmp"
rm "$$.tmp"
