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

   

   