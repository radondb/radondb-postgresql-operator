# ![LOGO](docs/images/logo_radondb.png)

> [English](README.md) | 中文

----

## 什么是 RadonDB PostgreSQL

[RadonDB PostgreSQL](https://github.com/radondb/radondb-postgresql-operator) 是基于 [PostgreSQL](https://www.postgresql.org/) 和 [PGO](https://github.com/CrunchyData/postgres-operator/) 开发的开源、高可用、云原生集群解决方案。

RadonDB PostgreSQL Operator 支持在 [Kubernetes](https://kubernetes.io) 和 [KubeSphere 3.1.x](https://kubesphere.com.cn) 平台部署。

## 架构图

![架构图](docs/images/operator.png)

## 核心功能

* 集群管理

  * 支持创建、扩容及删除 PostgreSQL 集群。
  * 支持从已有集群或集群备份，快速克隆创建集群。

* 高可用

  * 基于分布式一致的高可用解决方案，支持故障自动转移。
  * 支持跨 Kubernetes 集群部署备用 PostgreSQL 集群。

* 连接池
  
  支持使用开源 [pgBouncer](https://access.crunchydata.com/documentation/postgres-operator/v5/tutorial/connection-pooling/) 组件，最前沿连接池技术。

* 同步/异步复制

  支持工作负载同步/异步复制，确保事务不丢失。

* 灾备
  
  基于开源 [pgBackRest](https://www.pgbackrest.org/) 组件，支持备份与恢复。

* 监控

  基于开源 [pgMonitor](https://github.com/CrunchyData/pgmonitor) 组件，支持监控集群的运行状态。

* 备份
  
  * 支持备份到本地存储或任何支持 S3 协议的对象存储，如 QingStor 对象存储。
  * 支持全量、增量和差异增量备份。
  * 支持自定义备份时间策略。

## 支持的组件

RadonDB PostgreSQL Operator 主要由以下功能组件创建：

* [PostgreSQL](https://www.postgresql.org/)
* [pgBouncer](http://pgbouncer.github.io/)
* [pgMonitor](https://github.com/CrunchyData/pgmonitor)
* [pgBackRest](https://www.pgbackrest.org/)

除此之外，针对地理空间增强的 PostgreSQL + PostGIS 容器，还增加了以下两个组件：

* [PostGIS](http://postgis.net/)
* [pgRouting](https://pgrouting.org/)

PostgreSQL Operator Monitoring 使用以下组件:

* [pgMonitor](https://github.com/CrunchyData/pgmonitor)
* [Prometheus](https://github.com/prometheus/prometheus)
* [Grafana](https://github.com/grafana/grafana)
* [Alertmanager](https://github.com/prometheus/alertmanager)

## 快速入门

### 步骤一: 部署 PostgreSQL Operator

```shell
kubectl apply -f https://raw.githubusercontent.com/radondb/radondb-postgresql-operator/main/installers/kubectl/postgres-operator.yml
```

### 步骤二: 安装 PGO 客户端

```shell
curl https://raw.githubusercontent.com/radondb/radondb-postgresql-operator/main/installers/kubectl/client-setup.sh
```

```shell
chmod +x client-setup.sh
```

```shell
./client-setup.sh
```

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

### 步骤三: 部署集群

```shell
pgo create cluster hippo
```

## 协议

RadonDB PostgreSQL 基于 Apache 2.0 协议，详见 [LICENSE](./LICENSE)。

<p align="center">
<br/><br/>
如有任何关于 RadonDB PostgreSQL 的问题或建议，请在 GitHub 提交 Issue 反馈。
<br/>
</a>
</p>
