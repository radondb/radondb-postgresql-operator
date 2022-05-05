ARG BASEOS
ARG BASEVER
ARG PREFIX
ARG PGVERSION
ARG BACKREST_VERSION
FROM ${PREFIX}/pgo-base:${BASEOS}-${BASEVER}

ARG BASEOS
ARG PACKAGER

LABEL name="pgo-apiserver" \
	summary="RadonDB PostgreSQL Operator - Apiserver" \
	description="RadonDB PostgreSQL Operator - Apiserver"

RUN apt-get -y update && \
	apt-get -y install -y --no-install-recommends \
		"postgresql-${PGVERSION}" \
		; \
	ln -s /usr/lib/postgresql/$PG_MAJOR /usr/pgsql-$PG_MAJOR ;\
	\
	rm -rf /var/lib/apt/lists/*; \
	

ADD bin/apiserver /usr/local/bin
ADD installers/ansible/roles/pgo-operator/files/pgo-configs /default-pgo-config
ADD conf/postgres-operator/pgo.yaml /default-pgo-config/pgo.yaml

USER 2

ENTRYPOINT ["/usr/local/bin/apiserver"]
