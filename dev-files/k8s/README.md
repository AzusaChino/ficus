# Kubernetes Deploy

## StatefulSet

- 匹配 Pod name(网络标识)的模式为：$(statefulset名称)-$(序号)，比如上面的示例：web-0，web-1，web-2。
- StatefulSet 为每个 Pod 副本创建了一个 DNS 域名，这个域名的格式为： $(podname).(headless server name)，也就意味着服务间是通过 Pod 域名来通信而非 Pod
  IP，因为当 Pod 所在 Node 发生故障时，Pod 会被飘移到其它 Node 上，Pod IP 会发生变化，但是 Pod 域名不会有变化。
- StatefulSet 使用 Headless 服务来控制 Pod 的域名，这个域名的 FQDN 为：$(service name).$(namespace).svc.cluster.local，其中，“cluster.local”指的是集群的域名。
- 根据 volumeClaimTemplates，为每个 Pod 创建一个 pvc，pvc 的命名规则匹配模式：(volumeClaimTemplates.name)-(pod_name)，比如上面的 volumeMounts.name=www，
  Pod name=web-[0-2]，因此创建出来的 PVC 是 www-web-0、www-web-1、www-web-2。
- 删除 Pod 不会删除其 pvc，手动删除 pvc 将自动释放 pv。
