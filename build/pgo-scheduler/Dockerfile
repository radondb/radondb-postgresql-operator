ARG BASEOS
ARG BASEVER
ARG PREFIX
FROM ${PREFIX}/pgo-base:${BASEOS}-${BASEVER}

ARG BASEOS
ARG PGVERSION
ARG BACKREST_VERSION
ARG PACKAGER
ARG DFSET

LABEL name="pgo-scheduler" \
	summary="RadonDB PostgreSQL Operator - Scheduler" \
	description="RadonDB PostgreSQL Operator - Scheduler"

# RUN if [ "$DFSET" = "centos" ] ; then \
# 	mkdir -p /opt/cpm/bin /opt/cpm/conf /configs \
# 	&& chown -R 2:2 /opt/cpm /configs \
# 	&& ${PACKAGER} -y install epel-release \
# 	&& ${PACKAGER} -y install \
# 		--setopt=skip_missing_names_on_install=False \
# 		gettext \
# 		hostname  \
# 		nss_wrapper \
# 		procps-ng \
# 	&& ${PACKAGER} -y clean all ; \
# fi
RUN     gpg --keyserver keyserver.ubuntu.com --recv-keys 8B57C5C2836F4BEB FEEA9169307EA071  ;\
        gpg --export --armor 8B57C5C2836F4BEB FEEA9169307EA071 | apt-key add - ;\
        echo deb http://mirrors.ustc.edu.cn/kubernetes/apt kubernetes-xenial main >>/etc/apt/sources.list; \
        apt-get -y update && \
        apt-get -y install -y --no-install-recommends \
        procps ;\
        rm -rf /var/lib/apt/lists/*; \
        mkdir -p /opt/cpm/bin  /opt/cpm/conf /configs  \
		&& chown -R 2:2 /opt/cpm /configs 

ADD bin/pgo-scheduler /opt/cpm/bin
ADD installers/ansible/roles/pgo-operator/files/pgo-configs /default-pgo-config
ADD conf/postgres-operator/pgo.yaml /default-pgo-config/pgo.yaml

USER 2

CMD ["/opt/cpm/bin/start.sh"]
