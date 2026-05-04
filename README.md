# NQS (Neurader QuickSplit)

A simple developer process runner.

## Features

- Run multiple services easily
- Named processes
- Simple dashboard
- Cross-platform (Linux, macOS, Windows)

## Usage

Start a process:
nqs -n server node server.js

List processes:
nqs

Delete process:
nqs -d server

## Install (Linux example)

curl -L https://github.com/YOUR_USERNAME/nqs/releases/latest/download/nqs-linux-amd64 -o nqs
chmod +x nqs
sudo mv nqs /usr/local/bin/
