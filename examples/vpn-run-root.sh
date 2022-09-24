#!/bin/sh
set -e
set -x

if [ "$EUID" -ne 0 ]; then
  echo "Please run as root"
  exit
fi

[ -z $PORT ] && echo "Please set env PORT" && exit
[ -z $TABLE_FWMARK ] && echo "Please set env TABLE_FWMARK" && exit
command -v pkill >/dev/null 2>&1 || exit

BASEDIR=$(dirname "$0")
cd $BASEDIR

start() {
  # set bypass: fwmark
  ip rule add fwmark $TABLE_FWMARK table main || return
  ip -6 rule add fwmark $TABLE_FWMARK table main || return

  # for Tun2Socket
  iptables -I INPUT -s 172.19.0.2 -d 172.19.0.1 -p tcp -j ACCEPT
  ip6tables -I INPUT -s fdfe:dcba:9876::2 -d fdfe:dcba:9876::1 -p tcp -j ACCEPT

  "./nekobox_core" run -c "$CONFIG_PATH" --protect-listen-path "$PROTECT_LISTEN_PATH" --protect-fwmark $TABLE_FWMARK
}

stop() {
  for local in $BYPASS_IPS; do
    ip rule del to $local table main
  done
  iptables -D INPUT -s 172.19.0.2 -d 172.19.0.1 -p tcp -j ACCEPT
  ip6tables -D INPUT -s fdfe:dcba:9876::2 -d fdfe:dcba:9876::1 -p tcp -j ACCEPT
  ip rule del fwmark $TABLE_FWMARK
  ip -6 rule del fwmark $TABLE_FWMARK
}

if [ "$1" != "stop" ]; then
  start || true
fi

stop || true
