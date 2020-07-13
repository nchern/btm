#!/bin/sh

set -ue

username=$1

adduser --gecos "" --disabled-password ${username}
# TODO: create password

usermod -aG sudo ${username}

sudo -i -u ${username} sh -c 'cd $HOME && mkdir .ssh && touch .ssh/authorized_keys'
