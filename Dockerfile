FROM --platform=${TARGETPLATFORM:-linux/amd64} bitnami/minideb:latest

COPY kafkasearch-ui /app/kafkasearch-ui
COPY client_postgres.key /app/client_postgres.key
COPY client_postgres.crt /app/client_postgres.crt
COPY root.crt /app/root.crt
COPY ui/ /app/ui/

WORKDIR /app

RUN \
    apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get --no-install-recommends install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    inetutils-ping \
    less \
    locales \
    net-tools \
    postgresql-client \
    telnet \
    time \
    tzdata \
    vim-tiny \
    kafkacat \
    netcat \
    wget && \
    rm -rf /var/lib/apt/lists/ && \
    mkdir -p /app && \
    groupadd app && \
    useradd -g app -m -s /bin/bash app && \
    chown app:app -R /app && \
    chown app:app /app/client_postgres.key && \
    touch .env && \
    chmod 777 /tmp

USER app

CMD ["/app/kafkasearch-ui"]
