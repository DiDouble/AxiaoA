#!/bin/bash

my_ip_message=`curl -s https://www.ipip.net | grep "IP地址" | awk -F '>' '{print $3$5}' | awk -F '</' '{print $1$2}'`

ip_addr=`echo "$my_ip_message" | sed 's/span/:/g'`

echo "$ip_addr"
