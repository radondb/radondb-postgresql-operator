


# 快速开始

[TOC]

## 部署准备

已创建可用 Kubernetes 集群，目前支持如下平台创建的 Kubernetes 集群。

- Kubernetes 1.17+

## 部署步骤

### 步骤 1: 部署 `PostgreSQL Operator`

在设置了默认的`storageclass`的Kubernetes环境中运行以下命令用于部署`PostgreSQL Operator`

```shell
 git clone https://github.com/radondb/radondb-postgresql-operator
 cd radondb-postgresql-operator/installers/helm
 helm install demo . -n pgo --set pgo_client_container_install=true 
```

上面的命令将启动`pgo-deployer`容器来运行Ansible剧本将`PostgreSQL Operator`部署至`pgo`的命名空间下，需要1～5分钟时间，取决与您的网络环境。

### 步骤 2: 部署`PostgreSQL`集群

Postgres Operator 安装完成后,查找集群管理客户端所在的容器

```shell
kubectl get po -n pgo -l name=pgo-client
```

进入容器部署一个三节点的流复制集群

```shell
kubectl exec -it pgo-client-8b77cbcbf-pvtl2 -n pgo bash
pgo create cluster demo --replica-count=2 --password="RadonDB@123" --password-superuser="AmSuperUser" --username="demouser"  --database="demo" -n pgo
```

以上命令将花费1～2分钟完成部署，在客户端容器中，您可以通过如下命令查看集群状态：

`pgo show cluster -n pgo --all`

输出内容将显示集群的`primary`以及`replica`节点的状态，如果为`Running` 则代表集群部署完成

## 连接 RadonDB PostgreSQL 数据库

以下示例将运行一个带有`psql`客户端的容器连接到创建好的PostgreSQL数据库：

```shell
kubectl run -it psql -n pgo --image=governmentpaas/psql --restart=Never 
```

```shell
export PGPASSWORD="RadonDB@123"
psql -Udemouser -hdemo demo
```

或者您也可以通过`postgres`超级用户连接数据库:

```shell
export PGPASSWORD="AmSuperUser"
psql -Upostgres -hdemo 
```