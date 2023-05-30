#!/usr/bin/env bash
WORK_SPACE=$(cd $(dirname $(readlink -f "$0"));cd ..;pwd;)
echo "$WORK_SPACE"
rm -f "$WORK_SPACE"/dist.tar.gz
echo "###################"
echo "# begin go build"
export PATH=$PATH:/usr/local/go/bin
export GOROOT=/usr/local/go
export GOPROXY=https://goproxy.cn
CGO_ENABLED=0 GOOS=linux GOARCH=amd64
go mod tidy
go build -ldflags "-s -w" -tags="sonic avx" -o zy-apilabs "$WORK_SPACE"/cmd/api/
upx zy-apilabs
tar -zcv --exclude='configs/*.go' -f dist.tar.gz zy-apilabs scripts/ configs/
echo "###################"