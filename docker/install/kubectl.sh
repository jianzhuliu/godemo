#/bin/bash

VERSION=v1.21.0
SIGN=kubectl
TO_FOLD=/usr/local/bin


TARGET_FILE=kubernetes-client-linux-amd64.tar.gz
CMD_NAME=${SIGN}

#####################################

TO_FILE=${TO_FOLD}/${CMD_NAME}
TMP_FOLD=${SIGN}-download-test

# choose either URL
DOWNLOAD_URL="https://dl.k8s.io"

rm -f /tmp/${TARGET_FILE}

#curl -L ${DOWNLOAD_URL}/${VERSION}/${TARGET_FILE} -o /tmp/${TARGET_FILE}
wget ${DOWNLOAD_URL}/${VERSION}/${TARGET_FILE} -O /tmp/${TARGET_FILE}

rm -rf /tmp/${TMP_FOLD} && mkdir -p /tmp/${TMP_FOLD}
tar xzvf /tmp/${TARGET_FILE} -C /tmp/${TMP_FOLD} --strip-components=1

rm -f /tmp/${TARGET_FILE}


mkdir -p ${TO_FOLD}
rm -rf ${TO_FILE} && cp /tmp/${TMP_FOLD}/client/bin/${CMD_NAME} ${TO_FILE}

chmod +x ${TO_FILE}

rm -rf /tmp/${TMP_FOLD}

${CMD_NAME} version
