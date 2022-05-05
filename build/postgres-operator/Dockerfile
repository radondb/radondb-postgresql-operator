ARG BASEOS
ARG BASEVER
ARG PREFIX
FROM ${PREFIX}/pgo-base:${BASEOS}-${BASEVER}

ARG PGVERSION
ARG BACKREST_VERSION
ARG PACKAGER
ARG DFSET
ARG BASEOS

LABEL name="postgres-operator" \
        summary="RadonDB PostgreSQL Operator" \
        description="RadonDB PostgreSQL Operator"

RUN     apt-get -y update && \
        apt-get -y install -y --no-install-recommends \
        postgresql-${PGVERSION};\
        rm -rf /var/lib/apt/lists/*; 


ADD bin/postgres-operator /usr/local/bin
ADD installers/ansible/roles/pgo-operator/files/pgo-backrest-repo /default-pgo-backrest-repo
ADD installers/ansible/roles/pgo-operator/files/pgo-configs /default-pgo-config
ADD conf/postgres-operator/pgo.yaml /default-pgo-config/pgo.yaml

USER 2

ENTRYPOINT ["postgres-operator"]
