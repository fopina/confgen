#!/bin/sh

set -e

confgen -o /etc/nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf.tpl

exec "$@"