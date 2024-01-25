#!/bin/bash

for folder in `ls .`
do
    if [ -d "$folder" ]; then
        echo "$folder"
        cd "$folder"
        go build -o ../app
        cd ..
    fi
done