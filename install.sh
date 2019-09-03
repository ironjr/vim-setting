#!/bin/bash
PYTHON=python

# copy .vimrc
cp .vimrc $HOME

# install VimPlug
curl -fLo $HOME/.vim/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

# install Vundle
mkdir $HOME/.vim/bundle
git clone https://github.com/VundleVim/Vundle.vim.git $HOME/.vim/bundle/Vundle.vim

# install YouCompleteMe
cd $HOME/.vim/bundle
git clone https://github.com/ycm-core/YouCompleteMe
git submodule update --init --recursive
$PYTHON install.py --clang-completer
