
## 端口
1. containerPort：
- containerPort 是一个 Pod 中容器暴露的端口号。
- 当一个容器内有多个服务或应用程序运行时，containerPort 用于指定哪个服务的端口号应该被公开。
- 例如，如果在容器中有一个运行在端口 8080 上的 Web 服务，你可以使用 containerPort: 8080 来公开这个端口。
```shell
containers:
- name: example-container
  image: example-image
  ports:
  - containerPort: 8080
```
2. port:
- port 是服务（Service）暴露给集群内其他对象（如其他Pod）的端口号。
- 当你创建一个服务时，可以指定服务的 port。其他 Pod 可以通过这个端口与服务通信。
- 例如，创建一个服务，将外部流量（external traffic）引导到 Pod 的 containerPort 上：
```shell
apiVersion: v1
kind: Service
metadata:
  name: example-service
spec:
  selector:
    app: example-app
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
```
3. targetPort:
targetPort 是服务关联的后端 Pod 中容器的端口号。
当流量通过服务进入集群时，它将被转发到 targetPort 指定的端口上。
例如，在上述 Service 的定义中，targetPort: 8080 将流量转发到后端 Pod 中的 containerPort: 8080。
总结：
containerPort 是容器内服务的实际端口号。
port 是服务对外暴露的端口号，用于与集群外部的流量通信。
targetPort 是服务将流量转发到后端 Pod 中的容器的端口号。


## 持久化存储
1. PV
- 定义： PV 是集群中的一个资源对象，它表示集群中的一块持久化存储。它可以是集群中的物理存储资源，例如网络存储设备、本地存储设备或云存储服务。

- 手动或动态创建： PV 可以由集群管理员手动创建，也可以通过动态存储类（StorageClass）进行动态创建。动态创建是 Kubernetes 的一个特性，它允许管理员预定义存储类，而 PV 会在需要时由系统自动创建。

- 访问模式（Access Modes）： PV 有不同的访问模式，例如 ReadWriteOnce（单个节点读写）、ReadOnlyMany（多个节点只读）和 ReadWriteMany（多个节点读写）。这取决于存储后端的特性和配置。

- 持久卷的状态： PV 有不同的状态，例如 Available（可用）、Bound（已绑定，已被某个 PVC 使用）、Released（已释放，但仍然可以被重新绑定）等。


2. PVC
- 定义： PVC 是应用程序发出的请求，用于获取持久卷（PV）的一部分或全部容量。PVC 提供了抽象层，使得应用程序无需关心底层存储的实际细节。

- 声明存储需求： PVC 通过声明所需的存储类别、访问模式和存储容量来向集群发出请求。Kubernetes 根据这些要求选择合适的 PV 并进行绑定。

- 生命周期与应用绑定： PVC 的生命周期与应用程序的 Pod 绑定在一起。当 Pod 删除时，PVC 也会被释放，但 PV 可能会继续存在，以便重新绑定到其他 PVC。

- 自动或手动绑定： PVC 可以由开发人员手动创建并指定要使用的 PV，也可以通过动态存储类进行动态绑定，由系统自动选择满足条件的 PV。

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi

```
上述 YAML 文件的解释如下：

- apiVersion: 指定 Kubernetes API 版本，v1 表示使用核心 API 版本。
- kind: 指定资源类型，这里是 PersistentVolumeClaim。
- metadata: 包含元数据，其中 name 字段指定 PVC 的名称。
- spec: PVC 的规格，包含两个重要字段：
- accessModes: 定义访问模式，这里使用 ReadWriteOnce，表示卷将被挂载到一个节点，并且只能由一个节点读写。
- resources: 定义 PVC 的资源需求，其中 requests 字段指定存储的容量，这里使用 5Gi 表示5GB的存储容量。  


3. PV和PVC绑定
如果在创建 Persistent Volume (PV) 和 Persistent Volume Claim (PVC) 时没有明确指定 StorageClass，它们将会使用默认的 StorageClass（如果有的话）或者为 PV 不指定 StorageClass，为 PVC 不指定 StorageClass 将视为 standard 类型。

在这种情况下部署之后，pv没有和pvc绑定在一起。为pv和pvc添加字段`{"spec":{"storageClassName":"manual"}}`pv和pvc成功绑定在一起。
以下yaml为部署一个pv和pvc成功绑定在一起的case
```shell
apiVersion: v1
kind: PersistentVolume
metadata:
  name: nginx-pv
  labels:
    type: local
spec:
  storageClassName: manual
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/root/lhq/ltest"
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: nginx-pvc
spec:
  storageClassName: manual
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
```
也可以通过其他方式完成对于pv和pvc的绑定
```yml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi
  volumeName:  <pv-name> #请在这里将pv-name替换成需要绑定的pv名即可实现绑定
```
实际操作后pvc依然处于pending的状态，在保证了pv和pvc的storageClassName字段处于一致之后pvc才成功绑定到pv上去。

## 认证与安全


1. ServiceAccount
ServiceAccount 是一个用于身份验证的对象，它允许 Pod 在集群中运行时关联到一个身份。ServiceAccount 通常用于给 Pod 分配一组权限，以便它可以访问其他资源。
```yml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: my-service-account
```
2. Role
Role 是一种 Kubernetes 资源，用于定义对某个 Namespace 内资源的访问权限。它通过定义一组规则（rules）来实现这一目的。
```yml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: pod-reader
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "list"]
```
3. ClusterRole
ClusterRole 类似于 Role，但是它在整个集群中生效，而不仅仅在一个 Namespace 内。ClusterRole 允许定义对集群级别资源的访问权限。
```yml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: pod-reader
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "list"]
```
4. RoleBinding
RoleBinding 用于将 Role 绑定到特定的 ServiceAccount、User 或 Group。通过 RoleBinding，你可以将某个 ServiceAccount 的权限与某个 Role 关联起来。
```yml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: read-pods
subjects:
- kind: ServiceAccount
  name: my-service-account
  namespace: default
roleRef:
  kind: Role
  name: pod-reader
  apiGroup: rbac.authorization.k8s.io
```
5. ClusterRoleBinding
ClusterRoleBinding 类似于 RoleBinding，但是它将 ClusterRole 绑定到 ServiceAccount、User 或 Group，从而为其提供集群级别的权限。
```yml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: read-pods-cluster
subjects:
- kind: ServiceAccount
  name: my-service-account
  namespace: default
roleRef:
  kind: ClusterRole
  name: pod-reader
  apiGroup: rbac.authorization.k8s.io
```
6. CustomResourceDefinition
CustomResourceDefinition (CRD) 允许用户在 Kubernetes 中定义自己的自定义资源。通过定义 CustomResourceDefinition，用户可以引入新的资源类型和对象。这使得用户可以扩展 Kubernetes API，添加对自定义资源的支持。
```yml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: mycustomresources.example.com
spec:
  group: example.com
  names:
    kind: MyCustomResource
    plural: mycustomresources
  scope: Namespaced
  version: v1
```


