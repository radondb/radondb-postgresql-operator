## RadonDB PostgreSQL Operator简介

[RadonDB PostgreSQL Operator](https://github.com/radondb/radondb-postgresql-operator) 是基于 [PostgreSQL](https://www.postgresql.org/) 和 [PGO](https://github.com/CrunchyData/postgres-operator/) 开发的开源、高可用、云原生集群解决方案。它简化了在Kubernetes上部署和管理PostgreSQL集群的过程，无论是在开发测试环境中运行简单的PostgreSQL集群还是在生产中部署高可用，自动容错的集群，RadonDB PostgreSQL Operator均提供了满足您需要的特性。

## 支持的平台

- Kubernetes 1.17+
- OpenShift 3.1+
- Google Kubernetes Engine
- Amazon EKS
- Microsoft AKS
- VMware Tanzu

## 安装

### 步骤一: 部署 PostgreSQL Operator

PostgreSQL Operator支持多命名空间部署，以下示例会将`operator`及相关自定义资源部署到`pgo`的命名空间下:

```shell
kubectl create namespace pgo
kubectl apply -f https://raw.githubusercontent.com/radondb/radondb-postgresql-operator/main/installers/kubectl/postgres-operator.yml
```

查看部署后的状态：

```shell
 kubectl get all -n pgo --selector=name=postgres-operator
NAME                                   READY   STATUS    RESTARTS   AGE
pod/postgres-operator-cb9bf568-llljs   4/4     Running   1          13d

NAME                        TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)                      AGE
service/postgres-operator   ClusterIP   10.96.46.93   <none>        8443/TCP,4171/TCP,4150/TCP   13d

NAME                                         DESIRED   CURRENT   READY   AGE
replicaset.apps/postgres-operator-cb9bf568   1         1         1       13d
```

或者您可以用helm工具部署`operator`：

```shell
 git clone https://github.com/radondb/radondb-postgresql-operator
 cd radondb-postgresql-operator/installers/helm
 helm install demo .
```

### 步骤二: 部署 `pgo`客户端

`pgo`客户端是已经编译好的`postgres-operator`客户端，该工具提供了与 `CRD`交互的命令行`CLI`，方便您和CRD进行交互，如：

- 创建 PostgreSQL 集群

- 更新 PostgreSQL 集群资源分配

- 向 PostgreSQL 集群添加其他实用程序

- 备份/恢复集群

当然您也可以选择直接操作CRD资源。

您可以在能与`Kubenetes`集群交互的任何`Linux`环境上安装此客户端：

```shell
curl https://raw.githubusercontent.com/radondb/radondb-postgresql-operator/main/installers/kubectl/client-setup.sh |bash
```

安装完成之后，需要设置一些环境变量使其能正常工作：

```shell
cat <<EOF >> ~/.bashrc
export PGOUSER="${HOME?}/.pgo/pgo/pgouser"
export PGO_CA_CERT="${HOME?}/.pgo/pgo/client.crt"
export PGO_CLIENT_CERT="${HOME?}/.pgo/pgo/client.crt"
export PGO_CLIENT_KEY="${HOME?}/.pgo/pgo/client.key"
export PGO_APISERVER_URL='https://127.0.0.1:8443'
export PGO_NAMESPACE=pgo
#转发svc/postgres-operator到本地
EOF
source ~/.bashrc
```

> 以上的参数需要根据实际情况更改

转发svc/postgres-operator到本地：

```shell
nohup kubectl port-forward --namespace pgo svc/postgres-operator 8443:8443  &>/dev/null &
```

### 步骤三: 部署集群

详细的部署参数可以参考命令行帮助：

```shell
pgo --help
#更细项的帮助
pgo create cluster --help
```

```shell
pgo create cluster qingcloud \
#备份存储的类型,支持"posix", "s3", "gcs", "posix,s3" or "posix,gcs"
--pgbackrest-storage-type="s3" \
#PostgeSQL副本数量
--replica-count=3 \
#使用的镜像名称，带gis插件的镜像：qingcloud-postgres-gis-ha,不带gis插件的镜像：qingcloud-postgres-ha
--ccp-image=qingcloud-postgres-ha \
#镜像仓库
--ccp-image-prefix=docker.io/radondb \
#ockerhub上镜像的标签，目前支持centos8-12.7-4.7.1 和centos8-13.3-4.7.1
--ccp-image-tag=centos8-13.3-4.7.1 \
#支持s3协议的对象存储设置，主要用于备份，如果备份存储选择了s3则需要填写以下参数
--pgbackrest-s3-endpoint=s3.pek3b.qingstor.com \
--pgbackrest-s3-key=xxxxx \
--pgbackrest-s3-key-secret=xxxx \
--pgbackrest-s3-bucket=xxxx \
--pgbackrest-s3-region=xxx \
--pgbackrest-s3-verify-tls=false \
#启用适用于prometheus的指标收集器
--metrics \
#启用pgbadger
--pgbadger \
#启用pgbouncer
--pgbouncer \
#调试模式
--debug
```

使用以下命令观察集群创建过程：

```shell
kubectl get po -n pgo --watch
```

直到所有`POD`状态翻转为`Running`

#### 连接到PostgreSQL 集群

您可以使用以下`pgo show user -n pgo qingcloud`命令获取有关集群中用户的信息：

```shell
 pgo show user -n pgo qingcloud --show-system-accounts
 
 CLUSTER   USERNAME       PASSWORD                 EXPIRES STATUS ERROR 
--------- -------------- ------------------------ ------- ------ -----
qingcloud ccp_monitoring Dwpa|MCg,b4M+rY.>ZC0ONz4 never   ok           
qingcloud pgbouncer      MsTk4.auti9[0L2yDaHu/_Ni never   ok           
qingcloud postgres       1a4R-d7Po=,PS@R:-=?[gP(9 never   ok           
qingcloud primaryuser    =B8x*Haf*dCq+V4hkGSfh/.} never   ok           
qingcloud testuser       yTFeeH1|^DX<Bx4[?:B_/Q;M never   ok 
```

##### `psql`连接方式

查看服务：

```shell
 kubectl -n pgo get svc
NAME                             TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                                AGE
postgres-operator                ClusterIP   10.96.64.37     <none>        8443/TCP,4171/TCP,4150/TCP             58m
qingcloud                        ClusterIP   10.96.171.227   <none>        10000/TCP,9187/TCP,2022/TCP,5432/TCP   5m42s
qingcloud-backrest-shared-repo   ClusterIP   10.96.235.247   <none>        2022/TCP                               5m42s
qingcloud-pgbouncer              ClusterIP   10.96.234.49    <none>        5432/TCP                               4m16s
qingcloud-replica                ClusterIP   10.96.67.45     <none>        10000/TCP,9187/TCP,2022/TCP,5432/TCP   3m50s
```

使用`qingcloud`这个服务连接到数据库：

```shell
kubectl -n pgo port-forward svc/qingcloud 5432:5432
PGPASSWORD='yTFeeH1|^DX<Bx4[?:B_/Q;M' psql -h localhost -p 5432 -U testuser qingcloud
```

##### `pgAdmin`连接

`pgAdmin`是一个图形工具，可用于从 Web 浏览器管理和 PostgreSQL 数据库

`pgo create pgadmin -n pgo qingcloud`

创建pgAdmin的4实例需要一些时间, 等待创建完成后查看可用服务列表：

```shell
 kubectl -n pgo get svc qingcloud-pgadmin
NAME                TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
qingcloud-pgadmin   ClusterIP   10.96.239.152   <none>        5050/TCP   2m41s
```

修改初始用户密码：

`pgo update user -n pgo qingcloud --username=testuser --password=Qingcloud`

创建端口转发并连接:

```shell
kubectl -n pgo port-forward svc/qingcloud-pgadmin 5050:5050
```

将您的浏览器导航到[http://localhost:5050](http://localhost:5050/)并使用您的数据库用户名 ( `testuser`) 和密码(Qingcloud)连接

### 步骤三: 部署Prometheus服务端

[RadonDB PostgreSQL Operator](https://github.com/radondb/radondb-postgresql-operator) 支持您方便的部署`Prometheus Server`,请参考以下步骤：

```shell
 cd radondb-postgresql-operator/installers/metrics/helm
 helm upgrade demo-monitor  .
```

查看服务：

```shell
 kubectl get svc -n pgo qingcloud-grafana
NAME                TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
qingcloud-grafana   ClusterIP   10.96.222.20   <none>        3000/TCP   4m4s
```

创建端口转发并连接:

```shell
kubectl port-forward --namespace pgo svc/qingcloud-grafana --address 0.0.0.0 3000:3000
```

将您的浏览器导航到[http://localhost:3000](http://localhost:3000/)并使用初始用户admin/admin连接
