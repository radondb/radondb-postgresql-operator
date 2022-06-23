# ![LOGO](docs/images/logo_radondb.png)

> English | [中文](README_zh.md)

----

## What is RadonDB PostgreSQL Operator

[RadonDB PostgreSQL Operator](https://github.com/radondb/radondb-postgresql-operator) is an open-source, cloud-native, and high-availability cluster solution based on [PostgreSQL](https://www.postgresql.org/) and [PGO](https://github.com/CrunchyData/postgres-operator/).

RadonDB PostgreSQL Operator can be deployed on [Kubernetes](https://kubernetes.io) and [KubeSphere 3.1.x](https://kubesphere.com.cn).

## Quick start

- [Deploy RadonDB PostgreSQL on Kubernetes](docs/deploy_radondb_postgresql_operator_on_kubernetes.md)
- [Monitoring](docs/monitor_prometheus.md)

## Architecture

![Architecture](docs/images/operator.png)

## Key features

🏠 PostgreSQL cluster management
  
  * Create, scale, and delete PostgreSQL clusters.
  * Create new clusters by cloning the existing clusters or backups.

👏 High availability

  * Support automatic failover based on the distributed and consistent high-availability solution.
  * Deploy backup PostgreSQL clusters across Kubernetes clusters.

🎈 Connection pooling
  
  The advanced connection pooling supports using the open-source [pgBouncer](https://access.crunchydata.com/documentation/postgres-operator/v5/tutorial/connection-pooling/).

🎂 Replication
  
  Support asynchronous and synchronous replication of workloads sensitive to loss of transactions.

🎯 Disaster recovery

  Support backup and restore based on the open-source [pgBackRest](https://www.pgbackrest.org/).

🔔 Monitoring

  Track the health of the PostgreSQL clusters using the open-source [pgMonitor](https://github.com/CrunchyData/pgmonitor) library.

* Backup

  * Supports backup to local storage or the object storage that supports the S3 protocol like QingStor.
  * Support full, incremental, and differential backups.
  * Support customizing backup time policies.

## Included components

The following components are used by RadonDB PostgreSQL Operator:

* [PostgreSQL](https://www.postgresql.org/)
* [pgBouncer](http://pgbouncer.github.io/)
* [pgMonitor](https://github.com/CrunchyData/pgmonitor)
* [pgBackRest](https://www.pgbackrest.org/)

In addition, the following two components are added to the PostgreSQL container:

* [PostGIS](http://postgis.net/)
* [pgRouting](https://pgrouting.org/)

 The following components are used for PostgreSQL Operator monitoring:

* [pgMonitor](https://github.com/CrunchyData/pgmonitor)
* [Prometheus](https://github.com/prometheus/prometheus)
* [Grafana](https://github.com/grafana/grafana)
* [Alertmanager](https://github.com/prometheus/alertmanager)

## License

📖RadonDB PostgreSQL is released under the Apache 2.0, see [LICENSE](./LICENSE).

<p align="center">
<br/><br/>
😊For any RadonDB PostgreSQL bugs, issues, and feature requests, you can create an issue on GitHub.
<br/>
</a>
</p>