ARG BASE_IMAGE_OS
ARG DOCKERBASEREGISTRY
FROM ${DOCKERBASEREGISTRY}${BASE_IMAGE_OS}

ARG BASEOS
ARG RELVER
ARG PGVERSION
ARG PG_FULL
ARG PACKAGER
ARG DFSET


LABEL vendor="RadonDB" \
	url="https://radondb.com" \
	release="${RELVER}" \
	postgresql.version.major="${PGVERSION}" \
	postgresql.version="${PG_FULL}" \
	os.version="7.7" \
	org.opencontainers.image.vendor="RadonDB" \
	io.openshift.tags="postgresql,postgres,sql,nosql,radondb" \
	io.k8s.description="Trusted open source PostgreSQL-as-a-Service"


ENV TZ=Asia/Shanghai

RUN target=$(uname -m);rm /bin/sh && ln -s /bin/bash /bin/sh && ln -s /usr/lib/${target}-linux-gnu /usr/lib64

RUN echo deb http://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye main contrib non-free >/etc/apt/sources.list && \
	echo deb http://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-updates main contrib non-free >>/etc/apt/sources.list && \
	echo deb http://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-backports main contrib non-free >>/etc/apt/sources.list && \
	echo deb http://mirrors.tuna.tsinghua.edu.cn/debian-security bullseye-security main contrib non-free >>/etc/apt/sources.list && \
	apt-get -y update && \
	apt-get -y install -y --no-install-recommends \
	ca-certificates \
	libnss-wrapper \
	wget \
	gettext \
	gnupg \
	dirmngr \
	curl ;\
	groupmod -g 999 tape ;\
	echo deb [ signed-by=/usr/local/share/keyrings/postgres.gpg.asc ] http://mirrors.tuna.tsinghua.edu.cn/postgresql/repos/apt/ bullseye-pgdg main ${PGVERSION} >/etc/apt/sources.list.d/pgdg.list; \
	key='B97B0AFCAA1A47F044F244A07FCC7D46ACCC4CF8'; \
	export GNUPGHOME="$(mktemp -d)"; \
	mkdir -p /usr/local/share/keyrings/; \
	gpg --batch --keyserver keyserver.ubuntu.com --recv-keys "$key"; \
	gpg --batch --export --armor "$key" > /usr/local/share/keyrings/postgres.gpg.asc; \
	command -v gpgconf > /dev/null && gpgconf --kill all; \
	rm -rf "$GNUPGHOME";\
	apt-get -y update ;\
	rm -rf /var/lib/apt/lists/*; 