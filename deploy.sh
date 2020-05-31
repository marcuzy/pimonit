#!/usr/bin/env bash

env GOOS=linux GOARCH=arm go build
rsync -r . pi:~/pimonit/.deploy
ssh pi "
pkill pimonit
cd ~/pimonit
cp -r .deploy/* .
./run.sh
echo 'Done.'
"