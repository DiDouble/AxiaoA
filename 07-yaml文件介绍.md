

# yaml文件

创建一个namespace

```yaml
apiVersion: v1 #该pod适合的版本,一般都是固定的
kind: Namespace  #定义资源类型,
metadata: #源数据
    #下面的内容都是用来描述创建的资源的一些基本信息
    name: 这个名字可以随便写(这里指的的这个namespace叫什么名字)名称不可以重复,在不同的命名空间可以重复
    labels:  #标签
        key: value  #这里就是标签里面的key和value,可以随便写,也可以写name: 跟上面name一样,但是这里只是一个标签.只是一个key: value,一般资源调用的时候才会用到labels下面的东西.而metadata下面的name是用来显示和查询的.标签是可以重复的.
 ---
 apiVersion: v1
 kind: Pod
 metadata:
     name: website  #可以随便写
     labels:
         app: website  #不能写纯数字
         role: frontend
         #上面的这些都是描述的你要创建的资源本身的一些东西,下面是你要用的东西.
 spec: #指定
     containers:  #定义容器的属性
         - name: website  #这里定义的是容器的名称,名字也是可以随便写的,没必要跟资源一样.所谓的容器就是你里面要跑的是些什么应用.
             image: daoclud.io/library/nginx
             ports:  #制定端口,这样暴露的是你内部容器(应用)的端口,访问也只能通过你资源的ip+暴露的端口访问到.
                 - containerPort: 80
             
        
```

```shell
#创建资源
kubectl apply -f 资源名称.yaml
或者是
kubectl create -f 资源名称.yaml
删除的话
kubectl delete -f 资源名称.yaml
#查看pods所在的运行节点
kubectl get pods -o wide
#查看pods的详细信息
kubectl get pods -o yaml
或者
kubectl get pods pods-name  --output yaml
#验证语句的
kubectl create -f ./资源名称.yaml  --validate
#这是干嘛的得验证
echo 3 > /proc/sys/vm/drop_caches

```

资源的介绍

READY

* 1/1  #第二个1表示资源中有几个容器.

RESTARTS

* 资源的重启次数

___cut使用详解___



create和apply的区别

create修改yaml文件重新重置资源的时候,需要先删除之后才能create.

apply不用删除,可以直接apply新的yaml文件.

容器可选的设置属性包括:

```shell
name ,image ,command ,args ,workingDir, ports ,env ,resource ,volumeMounts ,livenessProbe ,readinessProbe ,livecycle ,terminationMessagePath, imagePullPolicy, securityContext, stdin, stdinOnce ,tty
```

#### NodeSelector

是提供用户将pod与node进行绑定的字段,用法:

```yaml
apiVersion: v1
kind: pod
...
spec:
  nodeSelector:
    disktype: ssd
```

#### nodeName

指定在某个node上运行

```yaml
...
spec:
  nodeName: node01
  
```

#### hostAliases

定义pod内部hosts文件

```yaml
apiVersion:ｖ１
kind: Pod
metadata:
  ...
spec:
  hostAliases:
  - IP: "10.1.2.3"
    hostnames:
    - "www.baidu.com"
    - "第二个地址"
```

共享进程名称空间

```yaml
....
spec:
  shareProcessNanmespace: true   #共享进程名称空间
```

pod和宿主机共享名称空间

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  hostNetwork: true
  hostIPC: true
  hostPID: true
  containers:
  - name: nginx
    image:nginx
  - name: shell
    image: busybox
    stdin: true
    tty: true
```

定义了共享宿主机的Network,IPC和PID Namespace .这样,此pod里的所有容器,会直接使用宿主机的网络、直接与宿主机进行IPC通信，看到宿主机正在运行的所有进程。

#### IPC

容器交互的时候还是采用了Linux常见的进程间交互方法（）	，包括信号量、消息队列和共享内存、socket、管道等，然而同vm不同的是，容器的进程间交互实际上还是host上具有相同 	pid名字空间中的进程间交互，因此需要IPC资源申请时加入名字可能回家呢信息，每个IPC资源有一个唯一的32位ID

#### 容器属性

pod属性中的Containers

“Containers”和“init Containers” 这两个字段属于pod对容器的定义，内容完全相同，只是init containers的生命周期，会先于所有的containers，兵器严格按照定义的顺序执行。

![1585054498451](/home/lovefei/Documents/AxiaoA/images/1585054498451.png)

#### imagePullPollcy

定义镜像的拉取策略

默认是Always

表示每次创建pod都重新拉取一次镜像

Never和IfNotPresent

表示pod永远不会注定拉取镜像，或者只在宿主机上不存在镜像的时候才拉取

#### Lifecycle

定义container Lifecycle Hooks 作用是在容器状态法伤变化时出发的一系列钩子

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: Lifecycle-demo
spec:
  containers:
  - name: Lifecycle-demo
  image: nginx
  lifecycle:
    postStart: #进入之后
      exec:
        command: ["/bin/sh","-C","echo Hello from the postStart handler > /usr/share/message"]
    preStop:  #退出之前
      exec:
        command: ["/usr/sbin/nginx","-S","quit"]
```

pod的生命周期

* pending
* runing
* succeeded
* failed
* unknown

### projected volument



#### secret

​     secret 用来保存小片敏感数据的k8s资源,例如密码,token,或者秘钥,这类数据当然也可以存放在POD或者镜像中,但是放在Secret中是为了更方便的控制如何使用数据,并减少暴露的风险

用户可以创建自己的secret,系统也会有自己的secret

pod需要先引用才能使用某个secret

__pod有两种模式调用secret__

1. 作为volume的一个域被一个或者多个容器挂载
2. 在拉取镜像的时候被k8s使用
   1. 使用kubectl create secret 命令创建
   2. yaml文件创建secret

内建的secret

​	由SecviceAccount创建的API证书附加的秘钥

​	k8s自动生成的用来访问apiserver的secret,所有的pod会默认使用这个Secret与apiserver通信

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: mysecret
type: Opaque
  username: 转码之后的
  password:转码之后的
 #转码命令  echo "hello" | base64
 
```

创建一个使用实例

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: mypod1
spec:
  containers:
  - name: mypod1
    image: redis
    volumeMounts:  #挂载一个卷
    - name: foo    #名字是下面创建的那个volume的名字
      mountpath: "/etc/foo"  #挂载在pod的这个目录下
      readOnly: true    #拥有只读权限
  volume:   #创建一个卷
  - name: foo  #给卷起一个名字
    secret:   
      secretName: mysecret-01  #选择的secret的名称
      items:     #用于将secret中的用户秘钥分别存放在不同的目录下
        - key: username  
          path: abc/name  #这个路径是相对路径,是/etc/foo下的路径.username重命名为name
        - key: password   #value跟secret中要一致
          path: def/pwd  #同上
```

#### configMap

用于存储文件的k8s资源对象,所有的配置内容都存储在etcd中.

ConfigMap跟Secret很像.

1. 通过命行创建

    ```shell
kubectl create configmap  myconfigmap_one  --from-literal=db.host=10.5.10.116  --from-literal=db.port='3306'
#不支持下划线
    ```

```shell
kubectl create configmap myconfigmapone --from-literal=db.user=zhaomeiyang  --from-literal=db.pass=yangyang
```

2. 通过指定文件创建

创建的文件是

![1585139223816](/home/lovefei/Documents/AxiaoA/images/1585139223816.png)

然后执行命令创建

```shell
kubectl create configmap test-config2 --from-file=./app.properties
```

3. 指定目录创建

   ```shell
   kubectl create configmap test-config3 --from-file=./configs
   # configs下面有两个配置文件,可以是一个或者多个
   ```

4. 使用yaml文件创建

   ```yaml
   apiVersion: v1
   kind: ConfigMap
   metadata:
     name: test-config4
     namespace: default
   data:   #data下面是你要传输的数据
     cache_host: memcached-gcxt   #如果说key对应value比较短,可以不用换行,如果说跟下面的my.cnf一样
     cache_port: "11211"
     cache_prefix: gcxt
     my.cnf |  #需要用管道符声明这里的配置是多行的,就是一个key:有多个value
       [mysqld]
       log-bin = mysql-bin
       haha = hehe
   ```

   configmap的使用

   * 通过环境变量的方式,直接传递给pod

   * 通过pod的命令行下运行的方式.

   * 使用volume的方式挂载到pod中

     实例configmap

     ```yaml
     apiVersion: v1
     kind: ConfigMap
     metadata:
       name: special-configmap
       namespace: default
     data:
       sepcial.how: very
       special.type: charm
     ```

     使用valueFrom (从什么地方获取值), configMapKeyRef, name, key 制定要用的key:
```yaml
     #调用configmap的实例
     apiVersion: v1
     kind: Pod
     metadata:
       name: dapi-test-pod
     spec:
       containers:
         - name: test-containers
           image: daocloud.io/library/nginx
           env:   # 设置环境变量的标志
             - name: SPECIAL_LEVEL_KEY   //这里是容器里设置的新的变量的名字  #下面key获取的值需要赋值给这个变量(SPECIAL_LEVEL_KEY)
               valueFrom:
                 configMapKeyRef:  #通过它获取下面的value
                   name: special-config  #configmap的名字
                   key: special.how    #需要获取的value所对应的key
             - name: SPECIAL_TYPE_KEY
               valueFrom:
                 configMapKeyRef:
                   name: special-config
                   key: special.type
restartPolocy: Never
```

     2. 通过envFrom  ,configMapRef , name 使得configmap中所有的key/value对都自动变成环境变量

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: dapi-test-pod
spec:
  containers:
    - name: test-containers
      image: image/url/nginx
      envFrom:
      - configMapRef:
        name: special-config
restartPolicy: Never
```

作为vloume挂载使用

```yaml
apiVersion: extensions/v1bate1
kind: Deployment
metadata:
  name: nginx-configmap
spec:
  replicas:1
  template:
    metadata:
      labels:
        app: nginx-configmap
  spec:
    containers:
    - name: ngnx-configmap
      image: nginx
      port:
      - containerPort: 80
      volumeMount: 
      - name: config-volume4
        mountPath: /tmp/config4
    volumes:
    - name: config-volume4
      configMap:
        name: test-config4
```

___注意: 一旦有pod使用了configmap,那么当configmap被删除的时候,原有的pod是不会受到影响的,但是新创建的会找不到configmap___

___如果用kubectl edit configmap configmapname 修改的话,那么pod里面的也会改变___

![1585143534172](/home/lovefei/Documents/AxiaoA/images/1585143534172.png)

#### Downward 用来获取pod的基本信息

通过Download API来讲POD的IP, 名称以及所对应的namespace注入到容器的环境变量中去,然后在容器中打印全部环境变量来进行验证

常用关联字段

![1585145653465](/home/lovefei/Documents/AxiaoA/images/1585145653465.png)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: test-env-pod
  namespace: default
spec:
  containers:
  - name: test-env-pod
    image: nginx
    command: ["/bin/sh","-C","env"]
    env:
    - name: POD_HOME
      valueFrom:
        fieldRef:
          fieldPath: metadata.name
    - name: POD_NAMESPACE
      valueFrom:
        filedRef:
          filedPath: metadata.namespace
```

vim 编辑小技巧:   set list  #可以显示每行末尾的空格或者换行.





#### ServiceAccount  (投射数据卷第四种)

​     当用户访问集群(例如使用kubectl命令)时, apiservice会将用户认证为一个特定的user account (目前通常是admin, 除非系统管理员自定义集群配置)

​     pod容器中的进程也可以与apiserver联系.当他们在联系apiserver的时候,他们会被认证为一个滕丁的ServiceAccount(;比如default)



不是用于管理k8s账户的,而是给pod里面的进程使用的.



##### 应用场景

service account 它并非不是给k8s集群的用户使用的,而是给pod里面的进程使用的,它为pod提供必要的身份认证.



___查看pod内部的一些目录___

kubectl  exec nginx ls /usr/secrets/kubernetes.io/serviceaccount       #这个目录是serviceaccent的认证存放的地方.

创建一个ServiceAccount

```shell
kubectl create serviceaccount mysa  (你要创建的serviceaccount的名字)
#查看创建的sa
kubectl describe sa saname
kubectl get secret
```

调用自己创建的sa

```yaml
apiVersion: v1
kind: pod
metadata:
  name: nginx-pod
  labels:
    app: my-pod
spec:
  containers:
  - name: my-pod
    image: nginx
    ports:
    - name: http
      containerPort: 80
  serviceAccountName: mysa   #写清自己要挂载的sa的名字
```

#### rbac  权限控制

在k8s中,授权有ABAC(基于属性的访问控制), RBAC(基于角色的访问控制), webhock , Node , AlwaysDeny(一直拒绝)  , 和AlwaysAllow(一直允许)六种.

在RBAC API中,通过如下授权

* 定义角色: 在定义角色时会指定次角色对于资源的访问控制规则
* 绑定角色: 讲主体与角色进行绑定,对用户进行访问授权

#### 创建k8s账号与rbac使用

默认使用的账号是service account  ,role(定义权限),   cluster-role  (集群角色,对整个k8s集群的命名空间生效

rolebinding (绑定角色)

clusterRoleBinding (集群访美的权限授予通过ClusterRoleBinding对象完成)  集群角色的绑定

![1585291247198](/home/lovefei/Documents/AxiaoA/images/1585291247198.png)

创建账号:

1. 创建私钥

   (umask 077; openssl genrsa -out wing.key 2048)

   用此私钥创建一个csr(证书签名请求)文件

   openssl req -new -key wing.key  -out wing.csr  -subj  "/CN=wing"

   openssl x509 -req -in wing.csr -CA  /etc/kubernetes/pki/ca.crt -CAkey /etc/kubernetes/pki/ca.key  -CAcreateserial -out wing.crt  -days 365    #需要调用k8s的证书,然后生成wing.crt证书.

2.  查看证书内容

   openssl  x509 -in wing.crt  -text -noout

   kubectl  config set-credentials wing  --client-certificate=./wing.crt  --client-key=./wing.key --embed-certs=true

3. 设置上线文

   kubectl config  set-context wing@kubernetes  --cluster=kubernetes  --user=wing

4. 查看当前工作的上下文

   kubectl  config view

5. 切换用户

   kubectl config use-context  wing@kubernetes

创建一个角色(role)

切回管理账号先

kubectl  config  use-context  kubernetes-admin@kubernetes

创建一个角色

kubectl create role  role-name(随便起)  --verb=get ,list ,watch (设置该role拥有的权限)  --resource=pod,svc(对pod和svc的权限)

绑定一个角色

kubectl create role myrole(name)  myrole-binding  --role=myrole  --user=wing



__切换用户__

kubectl config use-context  wing@kubernetes





#### 探针

```yaml
#命令行探针
...
spec:
  containers:
  - name: liveness
    image: nginx
    args:
    - /bin/sh
    - -C
    - touch /tmp/healthy; sleep 30; rm -rf /tmp/healthy; sleep 600
    livenessProbe:  #下面是探针的设置方法
      exec:   #健康检查的类型是exec,(执行命令)
        command:
        - cat
        - /tmp/healthy
        initialDelaySeconds: 5 #启动五秒后开始执行
        periodSeconds: 5   #每5秒执行一次
```

http状态探针

```yaml
...
spec:
  containers:
  - name: liveness-exec-container
    image: nginx
    imagePullPolicy: IfNotPresent
    ports:
      - name: http
        containerPort: 80
    livenessProbe:
      httpGet:
        port: http
        path: /index.html
      initialDelaySeconds: 1
      periodSeconds: 3
```

#### Pod Preset

![1585296899784](/home/lovefei/Documents/AxiaoA/images/1585296899784.png)

编写一个podPreset

```yaml
apiVersion: v1
kind: PodPreset
metadata:
  name: allow-database
spec:
  selector:
    matchLabels:
  env:
    - name: DB_PORT
      value: "6379"
  volumeMounts:
    - mountPath: /cache
      name: cache-volume
  volume: 
    - name: cache-colume
    emptyDir: {}
```

![1585297671609](/home/lovefei/Documents/AxiaoA/images/1585297671609.png)