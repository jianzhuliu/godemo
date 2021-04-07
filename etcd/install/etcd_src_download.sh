#/bin/bash

VERSION=v3.4.15
GITHUB="github.com/etcd-io/etcd"
SIGN=etcd


#####################################

SRC_FOLD=${SIGN}_${VERSION}
TMP_FILE=${SIGN}_src_${VERSION}.tar.gz

# choose either URL
GITHUB_URL="https://${GITHUB}/archive/refs/tags/${VERSION}.tar.gz"

DOWNLOAD_URL=${GITHUB_URL}

rm -f /tmp/${TMP_FILE}

#curl -L ${DOWNLOAD_URL} -o /tmp/${TMP_FILE}
wget ${DOWNLOAD_URL} -O /tmp/${TMP_FILE}

rm -rf /usr/local/src/${SRC_FOLD} && mkdir -p /usr/local/src/${SRC_FOLD}
tar xzvf /tmp/${TMP_FILE} -C /usr/local/src/${SRC_FOLD} --strip-components=1
rm -f /tmp/${TMP_FILE}



ls -l /usr/local/src/${SRC_FOLD}
