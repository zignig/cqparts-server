# cqparts-server

Proxy server and file watcher for cparts

This is an adjunct server to the cqparts project.

https://github.com/fragmuffin/cqparts

## Usage

define a CQPARTS_SERVER=http://localhost:8080 as an enviromental variable
and the latest version of cqparts will post to it and update a threejs
window automagically.

run the server with -d pointing towards a folder of .py files, it will watch for changes

./bin/server -d /opt/cqparts-bucket

running display() in a python script will push to the server and display.

## Building

Needs a golang dev environment and gb (https://github.com/constabulary/gb) for building.

## TODO

I will try to do binary release once it has settled.

