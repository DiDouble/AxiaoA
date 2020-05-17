# nfs安装详解

### 1.  环境说明

| Role      | Hostname |
| --------- | -------- |
| NFS服务端 | centos7  |
| NFS客户端 | centos7  |

### 2. NFS服务端

​	2.1 安装nfs-utils

```shell
yum info nfs-utils

Available Packages
Name        : nfs-utils
Arch        : x86_64
Epoch       : 1
Version     : 1.3.0
Release     : 0.65.el7
Size        : 412 k
Repo        : base/7/x86_64
Summary     : NFS utilities and supporting clients and daemons for the kernel NFS server
URL         : http://sourceforge.net/projects/nfs
License     : MIT and GPLv2 and GPLv2+ and BSD
Description : The nfs-utils package provides a daemon for the kernel NFS server and
            : related tools, which provides a much higher level of performance than the
            : traditional Linux NFS server used by most users.
            : 
            : This package also contains the showmount program.  Showmount queries the
            : mount daemon on a remote host for information about the NFS (Network File
            : System) server on the remote host.  For example, showmount can display the
            : clients which are mounted on that host.
            : 
            : This package also contains the mount.nfs and umount.nfs program.
#直接yum安装
yum install nfs-utils
# rpcbind 作为依赖会自动安装
```

 2.2 配置并开启服务

设置开机自启

```shell
systemctl enable rpcbind
systemctl enable nfs
```

 启动相关服务

```shell
systemctl start rpcbind
systemctl start nfs
```

 防火墙允许通过,不过一般情况下就已经关了

```shell
firewall-cmd --zone=public --permanent --add-service={rpc-bind,mountd,nfs}
firewall-cmd --reload
```

 2.3 配置共享目录

```shell
mkdir -p /data/mnt/pvtmp
chmod 755 /data/mnt/pvtmp
#修改配置文件
# 1. 只允许 abelsu7-ubuntu 访问
/data/mnt/pvtmp abelsu7-ubuntu(rw,sync,no_root_squash,no_all_squash)

# 2. 根据 IP 地址范围限制访问
/data/mnt/pvtmp 192.168.0.0/24(rw,sync,no_root_squash,no_all_squash)

# 3. 使用 * 表示访问不加限制
/data/mnt/pvtmp *(rw,sync,no_root_squash,no_all_squash)
```

关于`/etc/exports`中的参数:

* `/data/mnt/pvtmp` 用于给定当网络存储的目录
* `192.168.0.0/24`：**客户端 IP 范围**，`*`表示无限制
* `rw`：**权限设置**，可读可写
* `sync`：**同步共享目录**
* `no_root_squash`：可以使用`root`**授权**
* `no_all_squash`：可以使用**普通用户授权**

保存之后重启nfs服务

```shell
systemctl restart nfs
```

测试本地目录是否生效

```shell
showmount -e localhost   //这里建议修改hosts文件,直接写ip或者主机名.
#例如###########################
[root@DOCKER04 pvtmp]# cat /etc/hosts
127.0.0.1   localhost localhost.localdomain localhost4 localhost4.localdomain4
::1         localhost localhost.localdomain localhost6 localhost6.localdomain6
192.168.0.187    riswein.nfs
############################
[root@DOCKER04 pvtmp]# showmount -e riswein.nfs
Export list for riswein.nfs:
/data/mnt/pvtmp *
[root@DOCKER04 pvtmp]#
########################
```

3. ### 客户端安装

```shell
systemctl enable rpcbind
systemctl start rpcbind
showmount -e riswein.ntf
mkdir -p /mnt/kvm
mount -t nfs riswein.nfs:/data/mnt/pvtmp /mnt/kvm
#查看是否成功
[root@sqlback ~]# df -hT /mnt/pvtmp
Filesystem                  Type  Size  Used Avail Use% Mounted on
riswein.nfs:/data/mnt/pvtmp nfs4  2.2T   23G  2.2T   2% /mnt/pvtmp
[root@sqlback ~]#
######################
#如果需要写死的话，需要修改/etc/fstab文件
```

![image-20200513173936554](E:\AllProject\src\AxiaoA\images\image-20200513173936554.png)

重新加载一下

```shell
systemctl daemon-reload
mount | grep /mnt/kvm
```

测试NFS读写速度

```shell
> time dd if=/dev/zero of=/mnt/kvm-lun/test-nfs-speed bs=8k count=1024
> time dd if=/mnt/kvm-lun/test-nfs-speed of=/dev/null bs=8k count=1024
```

[文章来源](https://abelsu7.top/2019/10/17/centos7-install-nfs/)

