#!/bin/sh

CompileDaemon -build 'gb build' -command './bin/server -d /opt/cqparts-bucket/'  -pattern  "(.+\\.go|.+\\.tmpl|.+\\.js|.+\\.vue)$"
