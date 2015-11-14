#!/bin/sh

if (( $# != 1 )); then
    echo "\nERROR: debe especificar el nombre del binario del App Bundle!\n"
    exit 1
fi

if [[ "$1" = *[[:space:]]* ]]; then
  echo "\nERROR: el nombre del binario del App Bundle no debe tener espacios!\n"
  exit 1
fi

echo "\nLimpiando Directorios..."
find . -name ".DS_Store" -depth -exec rm -f {} \;
rm -rf build
rm -rf dist

echo "\nGenerando el binario..."
cd src
go clean
go build -tags appbundle -o "$1"
cd ..

echo "\nCreando App Bundle (py2app)..."
touch dummy.py
python3 setup.py py2app --semi-standalone &> /dev/null
rm -f dummy.py

echo "\nAdaptando el App Bundle..."
rm -rf dist/dummy.app/Contents/MacOS/**
rm -rf dist/dummy.app/Contents/Resources/**

echo "\nNormalizando los Frameworks y dylibs..."
rm -rf dist/dummy.app/Contents/Frameworks/Python.framework
rm -rf dist/dummy.app/Contents/Frameworks/libcrypto.1.0.0.dylib
rm -rf dist/dummy.app/Contents/Frameworks/libssl.1.0.0.dylib
rm -rf dist/dummy.app/Contents/Frameworks/liblzma.5.dylib

echo "\nCorrigiendo estructura..."
cp -v bundle_res/App.icns dist/dummy.app/Contents/Resources/
mv -v "src/$1" "dist/dummy.app/Contents/MacOS/$1"
sed "s/_APP_EXE_/$1/g" bundle_res/Info.plist \
  > dist/dummy.app/Contents/Info.plist

echo "\nCorrigiendo Frameworks y dylibs del binario..."
install_name_tool -change \
  /usr/local/opt/csfml/lib/libcsfml-window.2.3.dylib \
  @executable_path/../Frameworks/libcsfml-window.2.3.dylib \
  "dist/dummy.app/Contents/MacOS/$1"

install_name_tool -change \
  /usr/local/opt/csfml/lib/libcsfml-graphics.2.3.dylib \
  @executable_path/../Frameworks/libcsfml-graphics.2.3.dylib \
  "dist/dummy.app/Contents/MacOS/$1"

install_name_tool -change \
  /usr/local/opt/csfml/lib/libcsfml-audio.2.3.dylib \
  @executable_path/../Frameworks/libcsfml-audio.2.3.dylib \
  "dist/dummy.app/Contents/MacOS/$1"

install_name_tool -change \
  /usr/local/opt/csfml/lib/libcsfml-system.2.3.dylib \
  @executable_path/../Frameworks/libcsfml-system.2.3.dylib \
  "dist/dummy.app/Contents/MacOS/$1"

echo "\nCopiando recursos..."
cp -rv src/gfx dist/dummy.app/Contents/Resources/
cp -rv src/sfx dist/dummy.app/Contents/Resources/

echo "\nEstableciendo nombre del App Bundle..."
mv -v dist/dummy.app "dist/$1.app"

echo "\nLimpiando resultados..."
rm -rf build

echo "\n*** App Bundle creado con exito! :) ***\n"
