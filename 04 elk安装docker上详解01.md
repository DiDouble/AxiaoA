## elk安装使用手册

### 1. 安装环境

- 安装环境使用的是centos服务器,使用的是centos7环境

- 内核: Linux DOCKER04 3.10.0-514.el7.x86_64 #1 SMP Tue Nov 22 16:42:41 UTC 2016 x86_64 x86_64 x86_64 GNU/Linux

- docker 版本  19.03.7

- ```shell
  #阿里云加速配置
  {
          "registry-mirrors": ["https://e86ugcyk.mirror.aliyuncs.com"],
          "live-restore":true
  }
  ```

#### 	

#### 1 .1 syslog调试

```shell
#目的是提交到syslog中
systemctl status rsyslog
#或者是
netstat -antp | grep 514  #514是用于收集日志的.

#修改配置文件
vim /etc/rsyslog.conf
$ModLoad imuxsock # provides support for local system logging (e.g. via logger command)
$ModLoad imjournal # provides access to the systemd journal

#查看版本
rsyslogd -v

#重启服务
systemctl restart rsyslog

```

#### 1.2 测试syslog

```shell
docker run -d -p 80:80 --log-driver syslog --log-opt syslog-address=tcp://localhost:514 --log-opt tag="nginx_tag"  --name nginx_tag nginx
```

	##### 1.2.1 . 讲解
	
	--log-driver 
	
		> 收集日志的方式
		>
		> > none: 容器不输出任何内容
		> >
		> > json-file: 容器输出的格式以json的格式写道文件中
		> >
		> > syslog: 写道宿主机的syslog中
		> >
		> > gelf: 日志用GELF格式写入到Graylog中
		> >
		> > fluentd: 容器输出写入到宿主机的Fluentd中
		> >
		> > awslogs: 容器日志输出到亚马逊的存储中
		> >
		> > splunk: 日志写入到splunk中
		> >
		> > etwlogs: 写入到ETW中
		> >
		> > gcplogs: 日志写到GCP中
		> >
		> > nats: 写入到NATS中
	
	 --log-opt
	
	> 用于指定一些参数
	
	--syslog-address  
	
	> 地址
	
	tag  
	
	> 标签

##### 1.2.2 日志样式

![image-20200306181754521](C:\Users\P7XXTM1-G\AppData\Roaming\Typora\typora-user-images\image-20200306181754521.png)

执行下来就可以收集到日志,



### Elasticsearch 构建

```shell
docker run -d -p 9200:9200 -v /data/es/data:/usr/share/elasticsearch/data  --name elasticsearch elasticsearch
```

#### logstash



```shell
input (
  syslog (
    type => "rsyslog"
    port => 4560
  )
)
output {
  elasticsearch {
    hosts => [ "elasticssearch:9200" ]
  }
}
```

```shell
docker run -d  -p 4560:4560 -v /data/logstash/logstash.conf:/etc/logstash.conf  --link elasticsearch:elasticsearch  --name logstash logstash  logstash -f /etc/logstash.conf

```



### kibana

```shell
docker run -d -p 5601:5601 --link elasticsearch:elasticsearch -e ELASTICSEARCH_URL=http://172.17.0.3:9200 --name kibana kibana
```



```shell
CREATE DEFINER=`root`@`%` FUNCTION `getVideoCategoryChilds`(parentId VARCHAR(100)) RETURNS varchar(1000) CHARSET utf8
```

