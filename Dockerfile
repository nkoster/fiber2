FROM --platform=${TARGETPLATFORM:-linux/amd64} bitnami/minideb:latest

COPY kafkasearch-ui kafkasearch-ui
COPY client_postgres.key client_postgres.key
COPY client_postgres.crt client_postgres.crt
COPY root.crt root.crt
COPY ui/ ui/

RUN \
    apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get --no-install-recommends install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    git \
    gnupg \
    inetutils-ping \
    less \
    locales \
    net-tools \
    nginx \
    openssh-client \
    parallel \
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
    chown app:app /client_postgres.key && \
    touch .env && \
    chmod 777 /tmp

USER app

CMD ["/kafkasearch-ui"]
