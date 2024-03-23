


#### 强制删除pod
```shell
kubectl delete pod <pod-name> --grace-period=0 --force
```
#### 强制删除pv
1. 
```shell
kubectl delete pv <pv-name> --force --grace-period=0
```
2. 如果强制删除仍然无法成功，可能需要手动清理 PV。首先，尝试将 PV 的 claimRef 字段设置为 null：
```shell
kubectl patch pv <pv-name> -p '{"spec":{"claimRef": null}}'
```
再次执行命令1


### 进入pod
kubectl exec -it <pod name> -n <namespace> -- sh
将file复制到pod里的directory
kubectl cp <file> <pod name>:<directory path> -n <namespace
退出pod
ctrl D


### 修改部分字段
```shell
kubectl patch <resource-type> <resource-name> -p '{"spec":{"claimRef":{"name":"my-pvc"}}}'
```
```yml
spec:
  claimRef:
    name: my-pvc
```
#### 展示命名空间下的所有资源
运行以下检验命令可以列出特定namespace下的所有资源
```shell
kubectl api-resources --verbs=list --namespaced -o name | xargs -n 1 kubectl get -o name -n <namespace>
```

#### 纳管其他节点
```shell
kubeadm token create --print-join-command
```
运行之后会输出在其他节点上运行，可以实现纳管功能的命令