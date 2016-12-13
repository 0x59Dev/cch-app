#!/usr/bin/env bash

HOST_IP=$(ip route | grep default | awk '{print $3}')
APP_PORT="${APP_PORT:-8080}"

m4 "-DHOST_IP=$HOST_IP" "-DAPP_PORT=$APP_PORT" \
  /etc/nginx/conf.d/default.conf.m4 \
  >/etc/nginx/conf.d/default.conf

exec nginx -g "daemon off;"
