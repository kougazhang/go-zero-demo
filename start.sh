# start etcd
~/ectd/etcd-download-test/etcd &

. ./env.sh

# start user
pushd $project/rpc/user
go run user.go -f etc/user.yaml &
popd

# start blog
pushd $project/api
go run blog.go -f etc/blog-api.yaml &
popd
