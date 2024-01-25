#!/bin/bash

mkdir -p $1
cd $1

go mod init $1
go mod tidy

touch $1.go

cd ..
go work use $1

#git add $1 go.work
#git commit -m "initial commit"