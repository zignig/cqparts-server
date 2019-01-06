#!/bin/sh

CompileDaemon -build 'gb build' -command './bin/server -m cqparts_bucket -d /opt/cqparts_bucket/'  -pattern  "(.+\\.go|.+\\.tmpl|.+\\.js|.+\\.vue)$"
