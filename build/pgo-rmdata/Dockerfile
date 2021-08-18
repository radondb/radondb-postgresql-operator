ARG BASEOS
ARG BASEVER
ARG PREFIX
FROM ${PREFIX}/pgo-base:${BASEOS}-${BASEVER}

ARG PGVERSION
ARG BACKREST_VERSION
ARG PACKAGER
ARG DFSET

LABEL name="pgo-rmdata" \
	summary="RadonDB PostgreSQL Operator - Remove Data" \
	description="RadonDB PostgreSQL Operator - Remove Data"

ADD bin/pgo-rmdata/ /usr/local/bin

USER 2

CMD ["/usr/local/bin/start.sh"]
