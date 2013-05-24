#!/bin/bash

if [ ! -d "$GOPATH/assets" ]; then
    mkdir $GOPATH/assets/
fi

if [ -d "$GOPATH/assets/irrigation" ]; then
    echo "Removing old assets in irrigation/"
    rm -rf $GOPATH/assets/irrigation
fi

mkdir $GOPATH/assets/irrigation

cp -r $GOPATH/src/irrigation/views $GOPATH/assets/irrigation/
cp -r $GOPATH/src/irrigation/assets $GOPATH/assets/irrigation/
cp $GOPATH/src/irrigation/config.yml $GOPATH/assets/irrigation/
