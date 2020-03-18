# kibana使用手册

### 创建索引

```json
//添加索引
PUT /lib/
{
    "settings":{
        "index":{
            "number_of_shards": 5,  //分片是5个
            "number_of_replicas":1  //备份是1个,这个备份是指的在主机上存在多少个数据副本.
        }
    }
}

PUT lib2

查看已经创建好的索引的配置信息
GET /lib/_settings   /所有的类似_setting的都是在es中定义过的
//查看所有索引
GET _all/_settings

```

### filebeat日志收集软件

filebeat 如果中断了,下一次重启会从中断位置开始读新的日志文件,保证日志的更新.

中断的时候会记录上一次读到了什么位置.



```shell
读取的时候按照json格式进行读,需要添加一下两行配置
json.keys_under_root: true
json.overwrite_keys: true
```

![1583766546341](/tmp/1583766546341.png)

###### 自定义显示项目名称

![1583767214427](/tmp/1583767372058.png)



#### 多台日志收集





### logstash日志收集

```shell
#配置文件
input {
  tcp {
    iodec => json  #定义输入文件的格式
    port => 4560   #定义使用的端口 基础配置
  }
}
output {
  elasticsearch { 
    action => "index"              #制定要检测的关键子
    hosts => [ "172.17.0.3:9200" ] #es的地址   基础配置.
    index => "%{[appname]}"        #要获取哪个字段作为关键字,这个是从项目日志message中写着的.
  }
}
```



![1583805775072](/home/lovefei/Documents/AxiaoA/images/1583805775072.png)





### nginx日志转换成json格式

在nginx的ngx_http_log_module中可以修改,(nginx的一个模块)

```shell
log_format main '{ "time_local":"$time_local",
    
}'
```

![1583765278409](/tmp/1583765363364.png)



```shell
#vim 编辑器中进行复制粘贴
:2,7t8  #把从2~7行复制到8行
:2,7m8  #吧2~7行移动到8行


```

### 收集tomcat日志

```shell
#yum 安装tomcat
yum install tomcat  tomcat-webapps tomcat-admin-webapps tomcat-docs-webapp tomcat-javadoc -y




```

##### 1. tomcat日志转换成json格式

修改: ~/tomat/server.xml  文件.

删除139行  ` pattern="%h %l %u %t ....."`

替换为如下

![1583806742638](/home/lovefei/Documents/AxiaoA/images/1583806742638.png)

#### 2. 收集java的错误日志

​	根据时间节点来计算日志 [2020-03-10 12:20] 到下一个[2020-03-10 12:21] 作为分界

```shell
＃这里写的是ｉｎｐｕｔ日志
- type:log
  enabled: true
  paths:
    - /var/your/log/file/path/file.log
   tags: ["tag your use name"]
   multiline.pattern: '^\['　　＃用［作为标记
   multiline.negate: true
   multiline.match: after     #从开始到下一个开头
```

#### 3. docker 日志收集

```shell
#写到input模块中
filebeat.input:
- type: docker
  containers.ids:
    - '容器ID,要写全'
output.elasticsearch:   #表明是谁来收集日志
  hosts: ["es地址"]
   indices:
  - index: "waring-%{[beat.version]}-%{+yyy.MM}"
     when.contains:
       stream: "access"
   - index: "error-%{[beat.version]}-%{+yyy.MM}"
      when.containers:
        stream: "error"
setup.template.name: "docker"
setup.template.pattern: "docker-*"
setup.template.enabled: false
setup.template.overwrite: true
    
```

#### docker-compose  容器编排工具



#### 分开收集nginx的access日志和err日志

![1583997387128](/home/lovefei/Documents/AxiaoA/images/1583997387128.png)



#### filebeat 自带的modules收集日志

1. nginx
2. mongo
3. redis
4. mysql



## kibana画图

#### 使用redis作为缓存

#### 使用kafka作为缓存





