#!/bin/bash
set -xv

source modifyAtrr.sh
source refeshNode.sh

export IP=$(ifconfig eth0  |grep inet | awk '{print $2}' | cut -d: -f2)
sed -i "s/nodeip/$IP/g" /usr/local/service/hadoop/etc/hadoop/yarn-site.xml
/usr/local/service/hadoop/sbin/yarn-daemon.sh start nodemanager

echo "yarn pod start succeed"

while(true);do sleep 1;done

