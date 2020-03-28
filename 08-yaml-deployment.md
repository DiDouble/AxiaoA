# Deployment

k8s创建Deployment资源的创建流程

1. 用户通过kubectl 创建Deployment

2. Deployment创建ReplicaSet.

3. ReplicaSet创建pod

   ```yaml
   apiVersion: extensions/v1beta1
   kind: Deployment
   metadata:
     name: kube-100-site
   spec:
     replicas: 2   #副本个数
     template:    #模版
       metadata:   #下面定义的收拾pod的属性,单独pod怎么写,这里就怎么写,但是不能定义name
         labels:
           app: web
       spec:
         containers:
         - name: front-end
           image: nginx
           ports:
             - containerPort: 80
         - name: flaskapp-demo
           image: flaskapp
           ports:
             - containerPort: 5000
   ```

   创建一个带有volumen的Deployment

   ```yaml
   apiVersion: extenctions/v1beta1
   kind: Deployment
   metadata:
     labels:
       name: deploy-mount
   spec:
     replicas: 2
     template:
       metadata:
         labels:
           app: deployment-volume
       spec:
         containers:
         - name: nginx
           image: nginx
           ports:
           #- name: http
           - containerPort: 80
           volumenMount:
           - volumePath: /var/share/nginx/config/
             name: nginx-volume  #这个名字必须跟下面的volume的名字一样,这里不是声明,是引用.
         volumes:  #跟containers是一级的
           - name: nginx-volume
             emptryDir: {}   #这里正常情况下需要填写一个绝对路径,如果没写,默认在宿主机上就是容器中的目录.
           
   ```

   ##### volume讲解

   volumes:

   属于pod	对象的一部分.需要修改template.spec字段

   上面实例添加的一个volumes字段,定义这个pod声明的所有volume.他的名字叫做nginx-volume,类型是emptyDir

   关于__emptyDir__类型

      等同与docker的隐式Volume参数,即: 不显示声明宿主机目录的volume.所以k8s也会有宿主机上创建一个临时目录,这个目录将来就会被绑定关在到容器所声明的Volume目录上.

   k8s的emptyDir类型,只是把k8s创建的临时目录作为volume的宿主机目录,交给了docker,这么做的原因,是k8s不想以来docker自己创建的那个_data目录

   __volumeMount__

   pod中的容器,使用的是volumenMounts字段来声明自己要挂载哪个Volume,并通过mountPath字段来定义容器内的Volume目录

   __hostPath__

     k8s也提供了显示的volume蒂尼,他叫做hostpath.比如下面的yaml文件

   ```yaml
   ...
   volumes:
     - name: nginx-vol
       hostPath:
         path: /var/data
   ```

   

如果我们想让service访问到pod,是需要通过selector来进行选择.

选择器(selector)上跟上创建pod的deployment的标签.这样service就跟pod产生了联系.

```shell
#通过标签查询 pod
kubectl get pods -l app=nginx11 -o wide

```

#### spec

一个k8s的api对象的定义,大多数时候是metadata和spec两部分

前者存放的是对象的元数据,对所有的api对象来说,这一部分的字段跟格式基本上是一样的,而后者存放的是这个数据独有的定义,用来描绘它所要表达的功能.

##### replicas

定义的pod副本个数

template

定义了一个pod模版,这个模版描述了想要创建的pod的细节,



### 创建一个暴露端口的rc实例

```yaml
apiVersion: v1
kind: ReplicationController
metadata:
  name: nginx-controller
spec:
  replicas: 2
  selector:
    matchLabels:
      app: nginx01  #这里的名字,选的是下面具体容器的名字,就是资源的名字
  template:
    metadata:
      labels:
        app: nginx01
    spec:
      containers:
        - name: nginx #这里这个名字可以随便写
          image: nginx
          ports:
            -containerProt: 80
```

#### service

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx-service-nodeport
spec:
  ports:
    - port: 8000
      targetPort: 80
      protocol: TCP
  type: NodePort
  selector:
    name: nginx01
```

pod讲解

```shell
- port: 8080
  targetPort: 8080
  nodePort: 30062
  
  # port是 service暴露在cluster ip 上的端口
  #targetPort 是container的pod,也就是容器的端口
  #nodeport: 这个是service占用服务器的端口.不给的话会给一个随机端口
  
```

ingress

