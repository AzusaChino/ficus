# Kubernetes Deploy

## StatefulSet

- 匹配Pod name(网络标识)的模式为：$(statefulset名称)-$(序号)，比如上面的示例：web-0，web-1，web-2。 
- StatefulSet为每个Pod副本创建了一个DNS域名，这个域名的格式为： $(podname).(headless server name)，也就意味着服务间是通过Pod域名来通信而非Pod IP，因为当Pod所在Node发生故障时，Pod会被飘移到其它Node上，Pod IP会发生变化，但是Pod域名不会有变化。 
- StatefulSet使用Headless服务来控制Pod的域名，这个域名的FQDN为：$(service name).$(namespace).svc.cluster.local，其中，“cluster.local”指的是集群的域名。 
- 根据volumeClaimTemplates，为每个Pod创建一个pvc，pvc的命名规则匹配模式：(volumeClaimTemplates.name)-(pod_name)，比如上面的volumeMounts.name=www， Pod name=web-[0-2]，因此创建出来的PVC是www-web-0、www-web-1、www-web-2。 
- 删除Pod不会删除其pvc，手动删除pvc将自动释放pv。
