


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