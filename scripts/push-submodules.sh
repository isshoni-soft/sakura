#!/usr/bin/zsh

function commit() {
  cd $1 || return
  pwd
  git commit -a
  cd ..
}

commit kirito
commit roxxy

exit 0