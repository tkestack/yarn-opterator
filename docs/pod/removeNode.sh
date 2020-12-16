#!/bin/bash
set -xv

export IP=$(ifconfig eth0  |grep inet | awk '{print $2}' | cut -d: -f2)

mkdir /root/.ssh/

echo """""" > /root/.ssh/id_rsa

echo """[106.52.184.104]:36000 ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAPcEf90x7E2A7HXIMLd0pPY2IAj6pqjmzVmsVMw86caRlMejCsho+cdAG0RDUqc1FoZDOB6Jb4RSphFFKTYaXg=""" >/root/.ssh/known_hosts
chmod 400 /root/.ssh/id_rsa

ssh root@106.52.184.104 -p36000  "echo $IP >> /usr/local/service/hadoop/etc/hadoop/yarnexcludedhosts"
ssh root@106.52.184.104 -p36000  "/usr/local/service/hadoop/bin/yarn rmadmin -refreshNodes"
ssh root@106.52.184.104 -p36000  "sed -i '/'"$IP"'/d' /usr/local/service/hadoop/etc/hadoop/yarnexcludedhosts"
ssh root@106.52.184.104 -p36000  "sed -i '/'"$IP"'/d' /usr/local/service/hadoop/etc/hadoop/yarnhosts"
ssh root@106.52.184.104 -p36000  "/usr/local/service/hadoop/bin/yarn rmadmin -refreshNodes"




