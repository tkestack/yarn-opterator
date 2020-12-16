#!/bin/bash
set -xv

export IP=$(ifconfig eth0  |grep inet | awk '{print $2}' | cut -d: -f2)

mkdir /root/.ssh/

echo """-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA6k3cFPKyqydV9tc37oEA9sNs86Ng/NyavYiI+Zd9YaJGaLBp
psxZm4M0JNQCqYW3futnTP3w94ngdT+rxMezNiY+34wXXjve9O+4AYR+U371SHmo
cb1fd0fFC6xrRWpKqwUP70bvfnQSFKQ+2qYojmYsCX0H0bv/ZMqEAJsUbvr3Vm2r
De/WgPURwVOWaJP6JEc7wmpXKIRcBnNtao+4PMQDbGFA02XjsKMY4nURhNP2dSbO
CkOAXrMi5q+Fx3EvkosUhDY8isanH2cJBn6MiaiDJTm/QkZHVhhxY2EBlhtF9FgA
0F+wdmhCM6HnmvONSj5xukspaPeuWztDFUN7hwIDAQABAoIBAB0w6BuABTyHoREo
zSIc1mboABn2n+3A+lJkwVP/SLKySf1fBTqvuPZECWoRM+e07iCU6YDRHoVomxtg
fGD+1FgJucmWJY8q/GMdvpoJzMdQSPTnm1HYWx18RpNmvtKeJIFcFxkjiFED2wDR
WbdZ/jGHjzL7bc72kiIXjQyaRZhx1PkonMGc+BTIE1SBkkXM1ySkMJCmXpBmvP0l
NUU3THUyhGNraGwiLoFR/8x2DLqLe87rYTJMbYYYaVn8g6Sg1qVHvJ4K/71Iwudr
zam0VWLsaAZL5A3BvbMqEeR+AtOYaaoK5vYiTGFauBOwWroW5NMiNiEOQoMxpxqc
3Sj1FLECgYEA97s1zE6gnfcJeX2fDKKIyU/4HKOH+BubXHXlRSk9dBM376AL9DPw
37wNMXumWv0YQbLmn9xRfeJ44WjJoHF8vpXByhLNt4W8+HORl1ctkeD8GCPu3owk
VXC6gIOlSDWJU4ErzSeYGtK4QBntyLdXwdHJ8TDoEI8nHuDRd4S6xckCgYEA8h/r
JrI2XgOI9CSTD0u7FbGPI3vZVbbd+yXCeRId4pM9dMaqJXoYqOvCbAMDub2vh6j5
WX0PxQQteyNakr8hVKvyBaEb1g8mNTvpWc8b1RbrGE3kL0M6oxOezybIcY8fT6Yq
sltkNa1/fba9KzJxWFcYEHCrIRUhSTT/e66QHs8CgYEA1mBaQM/17FzgBRNxdJSe
fF9InTfirRDu1Adt/PigJnePGz1LuurL0kFAxYZ0Qh7tQ8VWEBavKpm185IjkUVE
JwUfawf0n5ELI5GW9vlJBQlF/nnx9wIdWxavPhEuEZvKl8mbJvDRjry1FzuY6u3F
8oLiF2c0G0hWGUGB4sSogAkCgYEA2vEd9InmO7E6oHqKOKSkcgNzigSuDKIlrLsC
VfSZ7Y5zXitfJDB6KBW4Y29+aPErzqJviApcviz+64CWoGgQvb4WRhzfTPu58x1P
75QmeNQWlo2or0w9s8VEL9HEI1vmVmHN7iZSQW+3/3fFK1CbyeRHsGYReQLgbJBu
DetWC2UCgYA1hR7QlbXCryGlXVBHDqwAoTD8smDBo9bdCQNVbCYTTRAdBKft3eI8
RSPswnfGm2e6nrbP+DHtQ2lc+VgihM00h9hNlIcCkmKBHGmFUEOoEW0GcHCwLpiq
rEUa9UReX66HBwmdr631mFleALBf6cIEjOKpBZil/FjwNc8cUuIvMQ==
-----END RSA PRIVATE KEY-----""" > /root/.ssh/id_rsa

echo """[106.52.184.104]:36000 ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAPcEf90x7E2A7HXIMLd0pPY2IAj6pqjmzVmsVMw86caRlMejCsho+cdAG0RDUqc1FoZDOB6Jb4RSphFFKTYaXg=""" >/root/.ssh/known_hosts
chmod 400 /root/.ssh/id_rsa

ssh root@106.52.184.104 -p36000  "echo $IP >> /usr/local/service/hadoop/etc/hadoop/yarnhosts"
ssh root@106.52.184.104 -p36000  "/usr/local/service/hadoop/bin/yarn rmadmin -refreshNodes"
