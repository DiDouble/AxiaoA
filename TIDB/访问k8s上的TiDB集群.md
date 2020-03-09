# 访问Kubernetes上的TiDB集群

通过Kubernetes访问TiDB集群，整体来说有两种方式

1. 集群外部访问
2. 集群内部访问

### 1.集群外部访问

​       集群之外的访问方式有两种，一种是通过`NodePort`暴露端口,还有就是通过LoadBalancer访问

##### 1.1 NodePort 



​      

### 2.集群内部访问

​       使用群内部访问的时候，通过访问TiDB server域名  `<release-name>-tidb.<namespace>` 即可。因为本身就是集群内部，通过Knbernetes内部的网络就可以直接使用。



```shell
#!/bin/bash
# Usage:
sudo loopmount file size mount-point touch $1 truncate -s $2 $1 
mke2fs -t ext4 -F $1 1> /dev/null 2> /dev/null 
if [[ ! -d $3 ]];
    then 
    echo $3 " not exist, creating..." 
    mkdir $3 
fi 
mount $1 $3 df -h |grep $3
```





