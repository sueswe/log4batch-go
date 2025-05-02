#!/bin/bash

source ~/.profile

cd "$HOME"/compile/log4batch-go || {
    echo "Status: $?"
    exit 4
}

echo "------------------------------------"
env | grep PATH
env | grep LOADED
pwd
echo "------------------------------------"

APPRELEASEVERSION=$(git rev-list -1 HEAD)
export APPRELEASEVERSION
echo "REV: $APPRELEASEVERSION"

echo ""
echo "compiling: go build log4batch.go -ldflags -X main.REV=$APPRELEASEVERSION"
go build -ldflags "-X main.REV=$APPRELEASEVERSION" -v -o /tmp/log4batch || {
    echo "Status: $?"
    exit 4
}

echo ""
echo "compiling: GOOS=aix GOARCH=ppc64 go build log4batch.go -ldflags -X main.REV=$APPRELEASEVERSION"
GOOS=aix GOARCH=ppc64 go build -ldflags "-X main.REV=$APPRELEASEVERSION" -v -o /tmp/log4batch.aix || {
    echo "Status: $?"
    exit 4
}

echo ""
echo "compiling: GOOS=windows GOARCH=amd64 go build log4batch.go -ldflags -X main.REV=$APPRELEASEVERSION"
GOOS=windows GOARCH=amd64 go build -ldflags "-X main.REV=$APPRELEASEVERSION" -v -o /tmp/log4batch.win64 || {
    echo "Status: $?"
    exit 4
}
