Contents
=================

   * [Prometheus Server 监控管理](#prometheus-server-监控管理)
      * [简介](#简介)
      * [操作准备](#操作准备)
      * [部署 Prometheus Server](#部署-prometheus-server)
      * [查看监控信息](#查看监控信息)


# Prometheus Server 监控管理

## 简介

[Prometheus](https://prometheus.io/) 通过数学算法实现强大的监控需求，并且支持监控容器化服务的动态变化。结合 Grafana 绘制可视化监控图形，并可联动 alertmanager 、Grafana 实现监控告警。

[RadonDB PostgreSQL Operator](https://github.com/radondb/radondb-postgresql-operator) 支持通过部署 Prometheus Server 实现数据库服务和资源监控。

## 操作准备

- 已启用适用于 Prometheus 的指标收集器。

## 部署 Prometheus Server

1. 执行如下命令，部署 Prometheus 服务端。

   ```shell
    cd radondb-postgresql-operator/installers/metrics/helm
    helm install demo-monitor  .
   ```

2. 查看 Prometheus 服务状态。

   ```shell
    kubectl get svc -n pgo qingcloud-grafana
   NAME                TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
   qingcloud-grafana   ClusterIP   10.96.222.20   <none>        3000/TCP   4m4s
   ```

3. 创建端口转发并连接 Prometheus 服务。

   ```shell
   kubectl port-forward --namespace pgo svc/qingcloud-grafana --address 0.0.0.0 3000:3000
   ```

## 查看监控信息

1. 在浏览器打开[http://localhost:3000](http://localhost:3000/) 监控页面。
2. 使用监控初始用户账号和密码，登录监控平台。

    初始用户账号 `admin`，初始密码 `admin`。

3. 您可以根据需要，分别查看数据库集群资源和服务状态。

更多 Prometheus 监控管理使用说明，请参见 [Prometheus Docs](https://prometheus.io/docs/introduction/overview/)。
