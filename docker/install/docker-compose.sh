#/bin/bash

VERSION=1.29.0
SIGN=docker-compose
GITHUB="github.com/docker/compose"
TO_FOLD=/usr/local/bin


TARGET_FILE="${SIGN}-$(uname -s)-$(uname -m)"
CMD_NAME=${SIGN}
TO_FILE="${TO_FOLD}/${CMD_NAME}"

#####################################


# choose either URL
GITHUB_URL="https://${GITHUB}/releases/download"
DOWNLOAD_URL="${GITHUB_URL}/${VERSION}/${TARGET_FILE}"
echo -e "going to download\r\n${DOWNLOAD_URL}\r\nto ${TO_FILE}"
#exit 1

mkdir -p ${TO_FOLD} && rm -f ${TO_FILE}

#curl -L ${DOWNLOAD_URL} -o ${TO_FILE}
wget ${DOWNLOAD_URL} -O ${TO_FILE} 

chmod +x ${TO_FILE}

${CMD_NAME} --version
