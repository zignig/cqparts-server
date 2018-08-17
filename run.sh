#!/bin/sh

CompileDaemon -build 'gb build' -command './bin/server -d /opt/cqparts_bucket/'  -pattern  "(.+\\.go|.+\\.tmpl|.+\\.js|.+\\.vue)$"
