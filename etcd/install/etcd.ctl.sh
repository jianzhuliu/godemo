#!/bin/bash

#import common config
. etcd.common.sh

echo etcdctl --endpoints=$ENDPOINTS $@
echo "---------------------------------------"
etcdctl --endpoints=$ENDPOINTS $@
