#/bin/bash

VERSION=v3.4.15
SIGN=etcd
GITHUB="github.com/etcd-io/etcd"
TO_FOLD=/usr/local/bin


TARGET_FILE=${SIGN}-${VERSION}-linux-amd64.tar.gz
CMD_NAME=${SIGN}

#####################################

TMP_FOLD=${SIGN}-download-test

# choose either URL
GITHUB_URL="https://${GITHUB}/releases/download"
DOWNLOAD_URL=${GITHUB_URL}

rm -f /tmp/${TARGET_FILE}

#curl -L ${DOWNLOAD_URL}/${VERSION}/${TARGET_FILE} -o /tmp/${TARGET_FILE}
wget ${DOWNLOAD_URL}/${VERSION}/${TARGET_FILE} -O /tmp/${TARGET_FILE}

rm -rf /tmp/${TMP_FOLD} && mkdir -p /tmp/${TMP_FOLD}
tar xzvf /tmp/${TARGET_FILE} -C /tmp/${TMP_FOLD} --strip-components=1

rm -f /tmp/${TARGET_FILE}

#/tmp/${TMP_FOLD}/${CMD_NAME} --help

mkdir -p ${TO_FOLD}
rm -rf ${TO_FOLD}/${CMD_NAME}* && mv /tmp/${TMP_FOLD}/${CMD_NAME}* ${TO_FOLD}/

rm -rf /tmp/${TMP_FOLD}

${CMD_NAME} --help
