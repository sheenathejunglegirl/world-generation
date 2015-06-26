#!/usr/bin/env bash

apt-get update > /dev/null

echo "Install Go 1.4"
echo "Installing dependency: git"
apt-get install git --assume-yes > /dev/null

echo "Downloading and unpacking go1.4linux-amd64"
curl -s https://storage.googleapis.com/golang/go1.4.linux-amd64.tar.gz | tar xz -C /usr/local

echo "Setting up GOPATH"
cat <<PROFILE >> /home/vagrant/.profile
# Setup Go
export GOPATH=\$HOME/go
export PATH=\$PATH:/usr/local/go/bin:\$GOPATH/bin
PROFILE

# Setup go workbench
echo "Creating go workbench"
su - vagrant -c 'mkdir -p $GOPATH/src $GOPATH/bin $GOPATH/pkg'

echo "Linking vagrant directory to workbench"
# Init project directory
su - vagrant -c 'mkdir -p $GOPATH/src/github.com/sheenathejunglegirl && ln -s /vagrant $GOPATH/src/github.com/sheenathejunglegirl/world-generation'

cat <<PROFILE >> ~vagrant/.profile
# Change directory to project
cd \$GOPATH/src/github.com/sheenathejunglegirl/world-generation
PROFILE
