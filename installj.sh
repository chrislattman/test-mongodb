#!/bin/bash

curl -LO https://repo.maven.apache.org/maven2/org/mongodb/mongodb-driver-sync/5.1.0/mongodb-driver-sync-5.1.0.jar
curl -LO https://repo.maven.apache.org/maven2/org/mongodb/bson/5.1.0/bson-5.1.0.jar
curl -LO https://repo.maven.apache.org/maven2/org/mongodb/mongodb-driver-core/5.1.0/mongodb-driver-core-5.1.0.jar
curl -LO https://repo.maven.apache.org/maven2/org/mongodb/bson-record-codec/5.1.0/bson-record-codec-5.1.0.jar
mkdir lib
mv ./*.jar lib
