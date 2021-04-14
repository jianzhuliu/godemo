#/bin/bash
###
## go 不同版本下载安装 , demo: install_go.sh 1.16.3
###

VERSION=$1
DEFAULT_VERSION=1.16.3
SIGN=go
GOROOT_PATH=/usr/local
GOPATH=/go
GOPROXY='https://goproxy.cn,direct'
PROFILE=/etc/profile
MINSIZE=1048

#默认参数判断
if [ "$VERSION" == "" ];then
	VERSION=${DEFAULT_VERSION}
fi


#####################################

TARGET_FILE="${SIGN}${VERSION}.linux-amd64.tar.gz"
#TARGET_FILE="${SIGN}${VERSION}.src.tar.gz"
DOWNLOAD_URL="https://studygolang.com/dl/golang/${TARGET_FILE}"

TMP_FILE=/tmp/${TARGET_FILE}

#版本号格式判断
PATTERN="^[0-9]{1,}\.[0-9]{1,}\.[0-9]{1,}$"
#if [[ ! ${VERSION} =~ ^[0-9]{1,}\.[0-9]{1,}\.[0-9]{1,}$ ]];then
if [[ ! ${VERSION} =~ ${PATTERN} ]];then
	echo "version: ${VERSION} is not right"
	exit
fi

rm -f ${TMP_FILE}

## 下载
#curl -L ${DOWNLOAD_URL} -o ${TMP_FILE}
wget ${DOWNLOAD_URL} -O ${TMP_FILE}

if [ $? -ne 0 ];then
	echo "download fail"
	exit 1
fi

if [ ! -f "${TMP_FILE}" ]; then
	echo "${TMP_FILE} is not exists"
	exit 1
fi

## 文件大小判断
FILE_SIZE=`wc -c <${TMP_FILE}`
if [ ${FILE_SIZE} -lt $MINSIZ ]; then
	echo "${TMP_FILE} size is too small"
	exit
fi

rm -rf /${GOROOT_PATH}/${SIGN} && mkdir -p ${GOROOT_PATH}
tar -xzf ${TMP_FILE} -C ${GOROOT_PATH}/ 
rm -f ${TMP_FILE}

mkdir -p ${GOPATH}/{bin,pkg,src}

cat <<EOF >>${PROFILE}
#########go  `date +"%F %T"`
export GOROOT=${GOROOT_PATH}/${SIGN}
export GOPATH=${GOPATH}
export GO111MODULE=auto
export GOPROXY='${GOPROXY}'
export PATH=\$PATH:\$GOROOT/bin:\$GOPATH/bin
#########go end
EOF

#. ${PROFILE}
#source ${PROFILE}
