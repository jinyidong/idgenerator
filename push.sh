#!/bin/bash
appName='IdGenerator'

version=v1.0.0

echo ---------------remove contanter...------------------
echo
     docker  rm -f ${appName} || true
echo
echo ---------------remove image...------------------
echo
   docker rmi  harbor.suiyi.com.cn/monitor/${appName}:${version}
echo
echo ---------------Build image...------------------
echo
   docker build -t harbor.suiyi.com.cn/monitor/${appName}:${version} .
echo
echo ---------------login harbor------------------
   docker login -u admin -p Sy123456 harbor.suiyi.com.cn
echo
echo ---------------push image------------------
echo
   docker push harbor.suiyi.com.cn/monitor/${appName}:${version}
echo

echo ---------------logout harbor------------------
echo
   docker logout harbor.suiyi.com.cn
echo