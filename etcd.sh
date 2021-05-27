ETCD_VER=v3.4.15
GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
DOWNLOAD_URL=${GITHUB_URL}
ETCD_PATH=~/ectd

mkdir $ETCD_PATH

rm -f ${ETCD_PATH}/etcd-${ETCD_VER}-darwin-amd64.zip
rm -rf ${ETCD_PATH}/etcd-download-test && mkdir -p ${ETCD_PATH}/etcd-download-test

curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-darwin-amd64.zip -o ${ETCD_PATH}/etcd-${ETCD_VER}-darwin-amd64.zip
unzip ${ETCD_PATH}/etcd-${ETCD_VER}-darwin-amd64.zip -d ${ETCD_PATH} && rm -f ${ETCD_PATH}/etcd-${ETCD_VER}-darwin-amd64.zip
mv ${ETCD_PATH}/etcd-${ETCD_VER}-darwin-amd64/* ${ETCD_PATH}/etcd-download-test && rm -rf mv ${ETCD_PATH}/etcd-${ETCD_VER}-darwin-amd64

${ETCD_PATH}/etcd-download-test/etcd --version
${ETCD_PATH}/etcd-download-test/etcdctl version