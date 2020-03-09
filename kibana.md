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









### nginx日志转换成json格式

在nginx的ngx_http_log_module中可以修改,(nginx的一个模块)

```shell
log_format main '{ "time_local":"$time_local",
    
}'
```

![1583765278409](/tmp/1583765363364.png)



