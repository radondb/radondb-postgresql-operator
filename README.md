# ![LOGO](docs/images/logo_radondb.png)

> English | [中文](README_zh.md)

----

## What is RadonDB PostgreSQL Operator

[RadonDB PostgreSQL Operator](https://github.com/radondb/radondb-postgresql-operator) is an open-source, cloud-native, highly availability cluster solutions based on [PostgreSQL](https://www.postgresql.org/) and [PGO](https://github.com/CrunchyData/postgres-operator/).

RadonDB PostgreSQL Operator supports [Kubernetes](https://kubernetes.io) or [KubeSphere 3.1.x](https://kubesphere.com.cn) platforms.

## Quick Start

- [Deploy RadonDB PostgreSQL on Kubernetes](docs/deploy_radondb_postgresql_operator_on_kubernetes.md)
- [Monitoring](docs/monitor_prometheus.md)

## Architecture

![Architecture](docs/images/operator.png)

## Main Features

* PostgreSQL Cluster Management
  
  * Create, Scale, Delete PostgreSQL clusters with smooth cluster management.
  * Create new clusters from the existing clusters or backups with efficient data cloning.

* High Availability

  * Support for automated failover that backed by the distributed consensus based high-availability solution.
  * Support for standby PostgreSQL clusters that work both within and across multiple Kubernetes clusters.

* Connection Pooling
  
  Advanced connection pooling support using [pgBouncer](https://access.crunchydata.com/documentation/postgres-operator/v5/tutorial/connection-pooling/).

* Advanced Replication
  
  Support for asynchronous or synchronous replication for workloads that are sensitive to losing transactions.

* Disaster Recovery

  Support for backups and restores that leverage the open source [pgBackRest](https://www.pgbackrest.org/) utility.

* Monitoring

  Track the health of the PostgreSQL clusters using the open source [pgMonitor](https://github.com/CrunchyData/pgmonitor) library.

* Backups

  * Backup to local storage. You can also store backups in any object storage system that supports the S3 protocol, such as QingStor.
  * Support for full, incremental, and differential backups as well as efficient delta restores.，
  * Support for user-defined backup time.

## Included Components

RadonDB PostgreSQL Operator include the following components:

* [PostgreSQL](https://www.postgresql.org/)
* [pgBouncer](http://pgbouncer.github.io/)
* [pgMonitor](https://github.com/CrunchyData/pgmonitor)
* [pgBackRest](https://www.pgbackrest.org/)

In addition to the above, the geospatially enhanced PostgreSQL + PostGIS container adds the following components:

* [PostGIS](http://postgis.net/)
* [pgRouting](https://pgrouting.org/)

PostgreSQL Operator Monitoring include the following components:

* [pgMonitor](https://github.com/CrunchyData/pgmonitor)
* [Prometheus](https://github.com/prometheus/prometheus)
* [Grafana](https://github.com/grafana/grafana)
* [Alertmanager](https://github.com/prometheus/alertmanager)

## License

RadonDB PostgreSQL is released under the Apache 2.0, see [LICENSE](./LICENSE).

<p align="center">
<br/><br/>
Please submit any RadonDB PostgreSQL bugs, issues, and feature requests to GitHub Issue.
<br/>
</a>
</p>