# /bin/bash

# Author: Oldyuan


copyFile() {
    dateTime=`date +"%Y_%m_%d"`
    fileName=`/usr/bin/find /home/gitlab/backups/ -name "*$(date +%Y_%m_%d)*"`
    echo $fileName
    scp $fileName root@192.168.0.187:/data/gitlabback
    if [ $? -ne 0 ]; then
        echo "失败的"
    else
        exit
    fi
}

main() {
    copyFile
}

main

