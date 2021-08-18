Contents
=================

   * [在 Kubernetes 上部署 RadonDB PostgreSQL 集群 (Operator)](#在-kubernetes-上部署-radondb-postgresql-集群-operator)
      * [简介](#简介)
      * [部署准备](#部署准备)
      * [部署步骤](#部署步骤)
        * [步骤 1: 部署 PostgreSQL Operator](#步骤-1-部署-postgresql-operator)
        * [步骤 2: 部署 PGO 客户端](#步骤-2-部署-pgo-客户端)
        * [步骤 3: 部署集群](#步骤-3-部署集群)
      * [连接 RadonDB PostgreSQL](#连接-radondb-postgresql-集群)
        * [通过 psql 连接](#通过-psql-连接)
        * [通过 pgAdmin 连接](#通过-pgadmin-连接)


# 在 Kubernetes 上部署 RadonDB PostgreSQL 集群 (Operator)

## 简介

[RadonDB PostgreSQL Operator](https://github.com/radondb/radondb-postgresql-operator) 是基于 [PostgreSQL](https://www.postgresql.org/) 和 [PGO](https://github.com/CrunchyData/postgres-operator/) 开发的开源、高可用、云原生集群解决方案。

RadonDB PostgreSQL Operator 简化了在 Kubernetes 上部署和管理 PostgreSQL 集群的过程，不仅适用于开发测试环境中运行的简单 PostgreSQL 集群，还适用于业务生产环境中运行的高可用或自动容错 PostgreSQL 集群。

## 部署准备

已创建可用 Kubernetes 集群，目前支持如下平台创建的 Kubernetes 集群。

- Kubernetes 1.17+
- OpenShift 3.1+
- Google Kubernetes Engine
- Amazon EKS
- Microsoft AKS
- VMware Tanzu

## 部署步骤

### 步骤 1: 部署 PostgreSQL Operator

PostgreSQL Operator 支持多命名空间部署，以下示例将 Operator 及相关自定义资源部署到 `pgo` 的命名空间下。

- 通过原生 kubectl 工具部署。

```shell
kubectl create namespace pgo
kubectl apply -f https://raw.githubusercontent.com/radondb/radondb-postgresql-operator/main/installers/kubectl/postgres-operator.yml
```

查看 PostgreSQL Operator 部署后的状态。

```shell
 kubectl get all -n pgo --selector=name=postgres-operator
NAME                                   READY   STATUS    RESTARTS   AGE
pod/postgres-operator-cb9bf568-llljs   4/4     Running   1          13d

NAME                        TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)                      AGE
service/postgres-operator   ClusterIP   10.96.46.93   <none>        8443/TCP,4171/TCP,4150/TCP   13d

NAME                                         DESIRED   CURRENT   READY   AGE
replicaset.apps/postgres-operator-cb9bf568   1         1         1       13d
```

- 通过 Helm 工具部署。

```shell
 git clone https://github.com/radondb/radondb-postgresql-operator
 cd radondb-postgresql-operator/installers/helm
 helm install demo .
```

### 步骤 2: 部署 PGO 客户端

> **说明**
> 
> 您也可以选择直接操作 CRD 资源。

PGO 客户端是已编译的 `postgres-operator` 客户端工具，可在与 Kubernetes 集群交互的任何 Linux 环境上安装此客户端。PGO 客户端可提供与 CRD 交互的命令行 CLI，方便与 CRD 进行交互，目前已支持如下功能：

- 创建 PostgreSQL 集群
- 更新 PostgreSQL 集群资源分配
- 向 PostgreSQL 集群添加其他实用程序
- 集群备份与恢复

1. 安装 PGO 客户端。

   ```shell
   curl https://raw.githubusercontent.com/radondb/radondb-postgresql-operator/main/installers/kubectl/client-setup.sh |bash
   ```

2. 设置环境变量，使其正常工作。

   > **说明**
   > 
   > 以下参数，可根据实际情况修改。

   ```shell
   cat <<EOF >> ~/.bashrc
   export PGOUSER="${HOME?}/.pgo/pgo/pgouser"
   export PGO_CA_CERT="${HOME?}/.pgo/pgo/client.crt"
   export PGO_CLIENT_CERT="${HOME?}/.pgo/pgo/client.crt"
   export PGO_CLIENT_KEY="${HOME?}/.pgo/pgo/client.key"
   export PGO_APISERVER_URL='https://127.0.0.1:8443'
   export PGO_NAMESPACE=pgo
   EOF
   source ~/.bashrc
   ```

3. 转发 `svc/postgres-operator` 到本地。

   ```shell
   nohup kubectl port-forward --namespace pgo svc/postgres-operator 8443:8443  &>/dev/null &
   ```

### 步骤 3: 部署集群

1. 启用 PGO 客户端，创建集群。

   ```shell
   pgo create cluster radondb \
   ```

2. 设置部署参数。
   
   部分参数可参见下表，并可执行 pgo 命令行帮助获取更详细的部署参数帮助。

   ```shell
   pgo --help
   pgo create cluster --help     #获取更多参数项说明
   ```

   |命令行 | 命令说明|
   |:----|:----|
   |   --pgbackrest-storage-type="s3" \   |  备份存储的类型。支持 `posix` 、`s3`、`gcs`、`posix,s3` 和 `posix,gcs`五种类型。  |
   |   --replica-count=3 \  |  PostgeSQL 副本数量。   |
   |   --ccp-image=radondb-postgres-ha \      |   使用的镜像名称。<br>带 `gis` 插件的镜像，例如 `radondb-postgres-gis-ha`。<br> 不带 `gis` 插件的镜像，例如 `radondb-postgres-ha`。  |
   |   --ccp-image-prefix=docker.io/radondb \     |   镜像仓库。  |
   |   --ccp-image-tag=centos8-13.3-4.7.1 \       |   ockerhub 上镜像的标签。目前支持 `centos8-12.7-4.7.1` 和 `centos8-13.3-4.7.1`。  |
   |   --pgbackrest-s3-endpoint=s3.pek3b.qingstor.com \ <br> --pgbackrest-s3-key=xxxxx \ <br> --pgbackrest-s3-key-secret=xxxx \ <br> --pgbackrest-s3-bucket=xxxx \ <br> --pgbackrest-s3-region=xxx \ <br> --pgbackrest-s3-verify-tls=false \  |   支持 s3 协议的对象存储设置，主要用于备份。若备份存储选择了 s3 则需要设置这部分参数。  |
   |   --metrics \      |   启用适用于 [Prometheus](https://prometheus.io/) 的指标收集器。  |
   |   --pgbadger \     |   启用 pgbadger。  |
   |   --debug          |   调试模式。        |

3. 查看集群创建状态。

   执行以下命令观察集群创建过程。待所有 Pod 状态切换为 `Running`，则集群创建完成。

   ```shell
   kubectl get po -n pgo --watch
   ```

## 连接 RadonDB PostgreSQL 集群

执行 `pgo show user -n pgo radondb` 命令，获取集群中用户账号信息。

以下以 `radondb` 集群为示例，获取数据库账号并连接数据库。

```shell
 pgo show user -n pgo radondb --show-system-accounts
 
 CLUSTER   USERNAME       PASSWORD                 EXPIRES STATUS ERROR 
--------- -------------- ------------------------ ------- ------ -----
radondb ccp_monitoring Dwpa|MCg,b4M+rY.>ZC0ONz4 never   ok           
radondb pgbouncer      MsTk4.auti9[0L2yDaHu/_Ni never   ok           
radondb postgres       1a4R-d7Po=,PS@R:-=?[gP(9 never   ok           
radondb primaryuser    =B8x*Haf*dCq+V4hkGSfh/.} never   ok           
radondb testuser       yTFeeH1|^DX<Bx4[?:B_/Q;M never   ok 
```

### 通过 psql 连接

1. 查看服务。

   ```shell
    kubectl -n pgo get svc
   NAME                             TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                                AGE
   postgres-operator                ClusterIP   10.96.64.37     <none>        8443/TCP,4171/TCP,4150/TCP             58m
   radondb                        ClusterIP   10.96.171.227   <none>        10000/TCP,9187/TCP,2022/TCP,5432/TCP   5m42s
   radondb-backrest-shared-repo   ClusterIP   10.96.235.247   <none>        2022/TCP                               5m42s
   radondb-pgbouncer              ClusterIP   10.96.234.49    <none>        5432/TCP                               4m16s
   radondb-replica                ClusterIP   10.96.67.45     <none>        10000/TCP,9187/TCP,2022/TCP,5432/TCP   3m50s
   ```

2. 以 `testuser` 账号为示例，连接到数据库。

   ```shell
   kubectl -n pgo port-forward svc/radondb 5432:5432
   PGPASSWORD='yTFeeH1|^DX<Bx4[?:B_/Q;M' psql -h localhost -p 5432 -U testuser radondb
   ```

### 通过 pgAdmin 连接

pgAdmin 是一个图形工具，可用于从 Web 浏览器连接和管理 PostgreSQL 数据库。

1. 执行 `pgo create pgadmin -n pgo radondb`命令，创建 pgAdmin 的实例。
2. 实例创建完成后，查看可用服务列表。

   ```shell
    kubectl -n pgo get svc radondb-pgadmin
   NAME                TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
   radondb-pgadmin   ClusterIP   10.96.239.152   <none>        5050/TCP   2m41s
   ```

3. 修改初始用户账号密码。

    ```shell
   pgo update user -n pgo radondb --username=testuser --password=RadonDB
   ```

4. 创建端口转发并连接数据库。

   ```shell
   kubectl -n pgo port-forward svc/radondb-pgadmin 5050:5050
   ```

5. 在浏览器打开[http://localhost:5050](http://localhost:5050/)，使用数据库用户账号 `testuser` 和密码即可连接到数据库。
