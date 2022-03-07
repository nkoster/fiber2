FROM liftm/kafkacat:1.4.0 as kafkacat
FROM busybox

COPY --from=kafkacat / /
COPY kafkasearch-ui /kafkasearch-ui
COPY client_postgres.key /client_postgres.key
COPY client_postgres.crt /client_postgres.crt
COPY root.crt /root.crt
COPY ui/ /ui/

CMD ["./kafkasearch-ui"]
