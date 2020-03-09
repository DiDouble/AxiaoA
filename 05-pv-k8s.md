## pv

### 1. 安装一个nfs

```shell
 yum install -y nfs-utils-*
 mkdir /data
 vim /etc/exports
 
 /data  10.0.0.0/24(rw.async,no_root_squash,no_all_squash)
 /code  10.0.0.0/24(rw.async,no_root_squash,no_all_squash)
 #重启服务使配置生效
 systemctl restart rpcbind
 systemctl restart nfs
 systemctl enable rpcbind
 systemctl enable nfs
 
```

### 2.构建pv

```yaml
apiVersion: v1
kind: PersistenVolume
metadata:
  name: pv001
  labels:   #非常关键
    type: nfs001
spec:
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteMany
  persistenVloimeReclaimPolicy: Recycle
  nfs:
    path: "/data"
    server: 10.0.0.11
    readonly: false
```

### 3.创建pvc

```yaml
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: pvc-wp
spec:
  accessModes:
    - ReadWriteMary
  resource:
    requests:
      storage: 1Gi
  selector:
    matchLabels:
      pv: nfs-pv2

```

