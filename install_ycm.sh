#!/bin/bash
cd $HOME/.vim/bundle/YouCompleteMe
yes | apt install build-essential cmake python3-dev
$PYTHON install.py --clang-completer
