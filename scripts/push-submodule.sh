#!/usr/bin/zsh

cd "$1" || return
pwd
git commit -a
cd ..

exit 0