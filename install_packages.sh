#!/bin/bash
# Install Rust.
# curl https://sh.rustup.rs -sSf | sh

# Install Go.
# add-apt-repository ppa:longsleep/golang-backports -y
# apt update
# apt install golang-go -y

# Install latest nodejs
if [ ! -x "$(command -v node)" ]; then
    curl -sfLS install-node.vercel.app/lts | sudo -E bash -
    # curl -sL https://deb.nodesource.com/setup_14.x | sudo -E bash -
fi
