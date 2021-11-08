#!/usr/bin/zsh

cd "$1" || return
pwd
git commit -a
git push
cd ..

exit 0