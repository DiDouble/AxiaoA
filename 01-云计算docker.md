```shell
 #查看镜像的所有信息
 docker image ls
 docker images
 #只查看镜像的ID号
 docker image ls -q
 #查看镜像的完整ID
 docker image ls --no-trunc
 
 #docker 镜像导出到本地文件中
 docker image save nginx（镜像名称/镜像ID） > /opt/nginx.tar.gz  
 #将已有镜像导入到image中
 docker image load -i /opt/nginx.tar.gz
 #给镜像打标签
 docker image tag 镜像id  nginx:latest
 #导入镜像的时候添加标签
 #docker image tag oldyuan/nginx:v1 load -i /opt/nginx.tar.gz (需要验证，好像不对命令)
 #删除镜像
 docker image rm 镜像名称
 #简写方式
 docker rmi 镜像ID
 dicker rmi -f 镜像ID  #强制删除镜像
 docker image rm -f 镜像名称   #强制删除镜像，当镜像存在依赖的时候删除
 docker image inspect 镜像名称  #查看镜像的详细信息
```

### 容器的管理

```shell
#容器docker container 支持的命令
  attach      Attach local standard input, output, and error streams to a running container
  commit      Create a new image from a container's changes
  cp          Copy files/folders between a container and the local filesystem
  create      Create a new container
  diff        Inspect changes to files or directories on a container's filesystem
  exec        Run a command in a running container
  export      Export a container's filesystem as a tar archive
  inspect     Display detailed information on one or more containers
  kill        Kill one or more running containers
  logs        Fetch the logs of a container
  ls          List containers
  pause       Pause all processes within one or more containers
  port        List port mappings or a specific mapping for the container
  prune       Remove all stopped containers
  rename      Rename a container
  restart     Restart one or more containers
  rm          Remove one or more containers
  run         Run a command in a new container
  start       Start one or more stopped containers
  stats       Display a live stream of container(s) resource usage statistics
  stop        Stop one or more running containers
  top         Display the running processes of a container
  unpause     Unpause all processes within one or more containers
  update      Update configuration of one or more containers
  wait        Block until one or more containers stop, then print their exit codes

```

```shell
#交互式启动
docker container run -it 镜像名称/镜像id#t指的是tti 我们每次打开一个窗口其实是打开一个tti，i就是交互式的
docker run -it 镜像ID
docker run -it --name="centos7" /bin/bash

#STATUS 容器的运行状态
Exeed  up
#查看容器状态
docker container ls
docker ls


```

### 守护类的容器

```shell
#在后台启动容器
docker run -d --name oldyuan_nginx -p 8080:80 nginx
#查看容器的详细信息
docker inspect 容器名称
#当交互式容器退出的时候自动删除容器
docker run -it --name oldyuan --rm nginx  #当退出的时候自动删除容器
```

###  只查看容器ID

```shell
#只查看容器ID
sudo docker container ls -q
```

### 容器管理

```shell
#暴露端口的

docker run -d --name nginx -p 8080:80 nginx:1.14

#启动容器
docker container start  容器ID    # -i 进行交互式启动
docker start  容器ID

ctrl +p +q  #前台启动容器之后退出

#暂停容器
docker container stop 容器ID
docker stop 容器ID
#容器的链接方式

docker container attach 容器名称   #登录到已存在的容器中,打开多个的话其实就是一个,就跟远程桌面一样,不会创建新的tti

docker container exec -it 容器名称 /bin/bash  #重启启动一个tti进行链接,可以打开多个.
                                  /bin/sh  #有系统底层的bash比较好用,如果是简单环境容器的话sh比较好使,经常好些命令不好用.
 #显示全部信息
 --no-trunk
 docker container ls -a --no-trunk
 
```

### 网络

同一个机器上的docker 容器之间是可以互相访问的.

```shell
#制定端口  任何主机都可以访问
docker run -p hostport:containerPort
#制定IP  制定某一个特定IP可以访问
docker run -p ip:hostport:containerPort
#制定IP但是宿主机端口随机   
docker run -p ip:containerPort
#随机占用宿主机的一个端口,允许任何地址访问
docker run -p containerPort
#默认是映射的是tcp的,可以指定UDP的
docker run -p hostport:containerPort/udp
#可以对多个端口进行映射
docker run -p 80:80  -p 81:81


```

#### 显示容器日志

```shell
docker logs testxx   #显示日志

#显示日志加上时间戳
docker logs -t testxx
#显示时时日志.跟tail -f 命令一样
docker logs -f testxx    #里面显示的时间是容器的时间.
#时时显示带有时间戳的日志
docker logs -tf testxx    #只是在上面的日志格式前面加上一个时间戳
#显示制定条数的日志
docker logs -tf --tail 100 testxx  #显示100行日志


docker logs -tf --tail 0 testxx
```

## docker 的数据卷的持久化存储

#### 1. nginx

```shell
#nginx 默认安装的时候目录
/usr/share/nginx/html
#将容器外的文件拷贝到容器内部
docker container cp index.html  容器名称:容器目录
例子:  docker container cp index.html nginx2:/usr/share/nginx/html/
#讲容器内的文件拷贝到本地的方式
docker container cp nginx2:/usr/share/nginx/html/50.html ./   #这里的./是指的当前目录

```

#### 2. 数据卷的存储和共享

```shell
#使用-v  制定外部目录跟容器内目录的映射关系.
docker run -d --name nginx2 -p 80:80 -v /opt/nginx:/usr/share/nginx/html  nginx                                       #如果容器内部没有这个目录,会在容器内创建

```

#### 3.挂载数据卷容器

```shell
#创建一个数据卷容器
docker run -it --name "nginx_volumes" -v /opt/Volume/a:/opt/a  -v /opt/Volume/b:/opt/b  centos6.7 /bin/bash
           #作用:大批量的容器都需要挂载相同的数据卷.
#使用数据卷容器
docker run -d -p 80:80  --volumes-from nginx_volumes --name http8085  nginx
```

### 制作局域网yum源

```shell
yum install -y vsftpd  #yum源文件的上传下载都是通过vsftp实现的.
#检测vsftp是不是搭建好了,通过lftp命令检测.
yum install -y lftp
lftp vsftp服务器地址.
#配置yum仓库
mkdir -p /var/ftp/centos6.9
mkdir -p /var/ftp/centos7.5    #创建目录用于存放相应的yum源文件
#配置yum源
cat >/yum.repos.d/ftp.repo <<EOF
baseurl=ftp://10.0.0.100/centos6.9
enable=1
gpgcheck=0
EOF
centos7的yum源配置
cat >/yum.repos.d/ftp_7.repo <<EOF
baseurl=ftp://10.0.0.100/centos7.5
enable=1
gpgcheck=0
EOF

#挂载镜像
mount -o loop /mnt/镜像名称  /var/ftp/centos6.9  #这样的话就可以通过ftp访问到镜像了.

```






























