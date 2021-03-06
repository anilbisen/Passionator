#!/bin/sh
if [ x = x$1 ]; then
    echo "Usage: ./patchBuild.sh <app> <mem-limit opt.>"
    exit
fi
APP_NAME=$1

if [ x = x$2 ]; then
    echo "Default limit of 256Mi"
    MEM_LIMIT=256Mi
else
    echo "Custom limit of $2"
    MEM_LIMIT=$2
fi

# Patch missing elements in build
oc patch bc/$APP_NAME -p '{
    "spec":{
        "source":{
            "sourceSecret":{
                "name": "scmsecret"}}}}'
if [ $? -ne 0 ]; then
    echo "ERROR in ./patchBuild.sh"
    exit 1
fi

oc patch bc/$APP_NAME -p '{
    "spec":{
        "resources":{
            "limits":{
                "memory": "'$MEM_LIMIT'"}}}}'
if [ $? -ne 0 ]; then
    echo "ERROR in ./patchBuild.sh"
    exit 1
fi

exit 0
