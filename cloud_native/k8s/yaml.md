#
##
### 存储
hostPath
在Kubernetes（K8s）中，HostPath是一种卷类型，允许Pod访问节点上的主机文件系统的特定路径。它允许将节点上的文件或目录挂载到Pod中

1.定义HostPath卷：
在Pod的配置中，可以通过以下方式定义HostPath卷：

```yaml
Copy code
volumes:
  - name: my-hostpath-volume
    hostPath:
      path: /path/on/node
```
这里/path/on/node是节点上的实际路径，将被挂载到Pod中。

2.在使用HostPath时，要确保所需的路径在所有节点上都存在，以确保Pod能够正常运行。

使用HostPath可能降低Pod的可移植性，因为不同的集群节点可能有不同的文件系统结构。在构建应用程序时，应考虑使用更通用和可移植的存储选项，如持久卷 (Persistent Volumes)。

示例：
下面是一个使用HostPath卷的简单示例：

```yaml
Copy code
apiVersion: v1
kind: Pod
metadata:
  name: mypod
spec:
  containers:
  - name: mycontainer
    image: nginx
    volumeMounts:
    - name: my-hostpath-volume
      mountPath: /var/www/html
  volumes:
  - name: my-hostpath-volume
    hostPath:
      path: /data/html
```
在这个例子中，Pod中的Nginx容器将节点上的/data/html路径挂载到/var/www/html路径。
注意hostPath是机器上的地址，把机器的这个目录挂载到容器的mountPath去。

```yml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: my-hostpath-pv
spec:
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: /host/path/data
```