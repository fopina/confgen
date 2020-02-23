#!/bin/sh

set -e

envsubst < /etc/nginx/conf.d/default.conf.1.tpl > /etc/nginx/conf.d/default.conf
[ -n "${SSL_CERT}" ] && envsubst < /etc/nginx/conf.d/default.conf.2.tpl >> /etc/nginx/conf.d/default.conf
envsubst < /etc/nginx/conf.d/default.conf.3.tpl >> /etc/nginx/conf.d/default.conf

exec "$@"