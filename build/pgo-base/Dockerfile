ARG BASE_IMAGE_OS
ARG DOCKERBASEREGISTRY
FROM ${DOCKERBASEREGISTRY}${BASE_IMAGE_OS}

ARG BASEOS
ARG RELVER
ARG PGVERSION
ARG PG_FULL
ARG PACKAGER
ARG DFSET

MAINTAINER info@radondb.com

LABEL vendor="RadonDB" \
	url="https://radondb.com" \
	release="${RELVER}" \
	postgresql.version.major="${PGVERSION}" \
	postgresql.version="${PG_FULL}" \
	os.version="7.7" \
	org.opencontainers.image.vendor="RadonDB" \
	io.openshift.tags="postgresql,postgres,sql,nosql,radondb" \
	io.k8s.description="Trusted open source PostgreSQL-as-a-Service"

COPY redhat/atomic/help.1 /help.1
COPY redhat/atomic/help.md /help.md
COPY licenses /licenses
ENV TZ=Asia/Shanghai
COPY conf/pip.conf pip.conf
ENV PIP_CONFIG_FILE pip.conf
ADD conf/Centos-8.repo /etc/yum.repos.d/
RUN rpm --import https://mirrors.aliyun.com/centos/RPM-GPG-KEY-CentOS-Official

RUN if [ "$BASEOS" = "centos7" ]; then \
    ${PACKAGER} -y update \
    && ${PACKAGER} -y clean all ; \
fi

RUN if [ "$BASEOS" = "centos8" ]; then \
    ${PACKAGER} -y update \
    && ${PACKAGER} -y install \
        --setopt=skip_missing_names_on_install=False \
        glibc-langpack-en \
    && ${PACKAGER} -y clean all \
    && ${PACKAGER} -qy module disable postgresql ; \
fi

RUN if [ "$BASEOS" = "rhel7" ] ; then \
	${PACKAGER} -y --enablerepo=rhel-7-server-ose-3.11-rpms update \
	&& ${PACKAGER} -y --enablerepo=rhel-7-server-ose-3.11-rpms clean all ; \
fi

RUN if [ "$BASEOS" = "ubi7" ] ; then \
        ${PACKAGER} -y --enablerepo=rhel-7-server-ose-3.11-rpms update \
        && ${PACKAGER} -y --enablerepo=rhel-7-server-ose-3.11-rpms clean all ; \
fi

RUN if  [ "$BASEOS" = "ubi8" ]; then \
    ${PACKAGER} -y update \
    && ${PACKAGER} -y install \
        glibc-langpack-en \
        procps-ng \
        less \
        vim-minimal \
    && ${PACKAGER} reinstall tzdata -y \
    && ${PACKAGER} -y clean all ; \
fi

# Create module file to disable postgres module, microdnf cannot do this with the current version
RUN if [ "$BASEOS" = "ubi8" ] ; then \
        echo "[postgresql]"  >> /etc/dnf/modules.d/postgresql.module \
        && echo "name=postgresql"  >> /etc/dnf/modules.d/postgresql.module \
        && echo "stream=10"  >> /etc/dnf/modules.d/postgresql.module \
        && echo "profiles="  >> /etc/dnf/modules.d/postgresql.module \
        && echo "state=disabled"  >> /etc/dnf/modules.d/postgresql.module ; \
fi

# RadonDB PostgreSQL repository
ADD conf/radondbpg${PGVERSION}.repo /etc/yum.repos.d/