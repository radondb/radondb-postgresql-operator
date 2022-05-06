ARG BASEOS
ARG BASEVER
ARG PREFIX
FROM ${PREFIX}/pgo-base:${BASEOS}-${BASEVER}

ARG BASEOS
ARG ANSIBLE_VERSION
ARG PACKAGER
ARG DFSET

LABEL name="pgo-deployer" \
    summary="RadonDB PostgreSQL Operator - developer" \
    description="RadonDB PostgreSQL Operator - developer"



RUN     gpg --keyserver keyserver.ubuntu.com --recv-keys 8B57C5C2836F4BEB FEEA9169307EA071  ;\
        gpg --export --armor 8B57C5C2836F4BEB FEEA9169307EA071 | apt-key add - ;\
        echo deb http://mirrors.ustc.edu.cn/kubernetes/apt kubernetes-xenial main >>/etc/apt/sources.list; \
        apt-get -y update && \
        apt-get -y install -y --no-install-recommends \
        kubectl \  
        python3-jmespath \   
        ansible; \
        rm -rf /var/lib/apt/lists/*; 

RUN mkdir -p /opt/cpm/bin

COPY installers/ansible /ansible/postgres-operator
COPY installers/metrics/ansible /ansible/metrics
ADD tools/pgmonitor /opt/radondb/pgmonitor
COPY installers/image/bin/pgo-deploy.sh /pgo-deploy.sh
ADD bin/common /opt/cpm/bin

ENV ANSIBLE_CONFIG="/ansible/postgres-operator/ansible.cfg"
ENV HOME="/tmp"

# Defines a unique directory name that will be utilized by the nss_wrapper in the UID script
ENV NSS_WRAPPER_SUBDIR="deployer"

ENTRYPOINT ["/opt/cpm/bin/uid_daemon.sh"]

USER 2

CMD ["/pgo-deploy.sh"]
