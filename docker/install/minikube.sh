#/bin/bash

VERSION=v1.19.0
SIGN=minikube
GITHUB="github.com/kubernetes/minikube"
TO_FOLD=/usr/local/bin


TARGET_FILE="${SIGN}-linux-amd64"
CMD_NAME=${SIGN}
TO_FILE="${TO_FOLD}/${CMD_NAME}"

#####################################

TMP_FILE=/tmp/${TARGET_FILE}
rm -f ${TMP_FILE}

# choose either URL
GITHUB_URL="https://${GITHUB}/releases/download"
DOWNLOAD_URL="${GITHUB_URL}/${VERSION}/${TARGET_FILE}"
echo -e "going to download\r\n${DOWNLOAD_URL}\r\nto ${TMP_FILE}"
#exit 1


#curl -L ${DOWNLOAD_URL} -o ${TMP_FILE}
wget ${DOWNLOAD_URL} -O ${TMP_FILE} 

mkdir -p ${TO_FOLD} && rm -f ${TO_FILE}

install ${TMP_FILE} ${TO_FILE}

rm -f ${TMP_FILE}*

${CMD_NAME} --help
