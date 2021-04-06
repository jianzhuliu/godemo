#!/bin/bash

#import common config
. etcd.common.sh

THIS_IP=`ifconfig | grep "inet addr:" | grep -v "127.0.0.1" | awk '{print $2}' | tr -d "addr:"`
THIS_NAME=${THIS_IP}

echo "this ip is ${THIS_IP}"
echo "----------------------------------------------"

etcd --data-dir=data.etcd --name ${THIS_NAME} \
	--initial-advertise-peer-urls http://${THIS_IP}:2380 \
	--listen-peer-urls http://${THIS_IP}:2380 \
	--advertise-client-urls http://${THIS_IP}:2379 \
	--listen-client-urls http://${THIS_IP}:2379 \
	--initial-cluster ${CLUSTER} \
	--initial-cluster-state ${CLUSTER_STATE} \
	--initial-cluster-token ${TOKEN}	
