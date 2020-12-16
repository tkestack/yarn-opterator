#!/bin/bash

set -xv

export IP=$(ifconfig eth0  |grep inet | awk '{print $2}' | cut -d: -f2)
hostname $IP
echo $IP > /etc/hostname
rm -rf /etc/hosts.backup
cp -i /etc/hosts /etc/hosts.backup
sed -i '/'"$IP"'/d' /etc/hosts.backup
cat  /etc/hosts.backup > /etc/hosts

echo "modifyAttr succeed"
