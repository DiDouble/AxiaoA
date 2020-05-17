# /bin/bash

#Author: Oldyuan


funcShellLink() {
    while true
    do
        Free=`free -m | awk 'NR==2{print $NF}'`
        if [ $Free -lt 100 ]
        then
            echo $Free |mail -s "当前内存" 邮箱地址
        fi
        sleep 60
    done
}