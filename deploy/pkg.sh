#!/usr/bin/env bash

if [[ ! $2 ]];then
    echo "do=tidy  or do=vendor"
    exit 2
fi

root=$1
cd ${root}

export GO111MODULE=on

export GOPROXY=https://goproxy.io

go mod $2
