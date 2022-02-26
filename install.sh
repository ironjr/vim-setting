#!/bin/bash
PYTHON=python

# copy .vimrc
cp .vimrc $HOME

# Install VimPlug
curl -fLo $HOME/.vim/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

# Install Vundle
mkdir $HOME/.vim/bundle
git clone https://github.com/VundleVim/Vundle.vim.git $HOME/.vim/bundle/Vundle.vim

# Install YouCompleteMe
# cd $HOME/.vim/bundle
# git clone --recursive https://github.com/ycm-core/YouCompleteMe
# cd YouCompleteMe
