# /bin/bash

# Author: Oldyuan

searchFile() {
    fileName=`/usr/bin/find /home/gitlab/backups/ -name "*$(date +%Y_%m_%d)*"`
    if [ -z $fileName ];then
        echo "没有生成"
    else
        copyFile
    fi
}

copyFile() {
    dateTime=`date +"%Y_%m_%d"`
    fileName=`/usr/bin/find /home/gitlab/backups/ -name "*$(date +%Y_%m_%d)*"`
    echo $fileName
    scp $fileName root@192.168.0.187:/data/gitlabback
    if [ $? -ne 0 ]; then
        echo "失败的，用于邮件发送失败吧"
    else
        exit
    fi
}

main() {
    searchFile
}

main

