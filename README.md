# Kafka Search UI Backend

For a limited or selected set of kafka topics, a data pipeline (spark) updates a postgres database with kafka message identifier
keys/values and partitions/offsets.

The kafkasearch-ui backend serves an API for a UI, that queries the database, and consults kafka to fetch the raw kafka message as JSON.

### Build

```
CGO_ENABLED=0 go build -ldflags="-extldflags=-static"
docker build -t fhirstation-kafkasearch-ui .
```

The UI directory holds the search UI frontend React App (internal project, not yet on Github)

### Third-party Resources
                                                                                                                   
* Gofiber: https://gofiber.io/ https://github.com/gofiber/fiber
* Statically compiled kafkacat docker image from https://github.com/jcaesar/kafkacat-static
* Docker base: https://registry.hub.docker.com/_/busybox/

<br />
<div style="white-space:nowrap">
  <img src="https://gofiber.io/assets/images/logo.svg" height="40px" alt="Gofiber"> &nbsp;
  <img src="https://raw.githubusercontent.com/edenhill/kcat/master/resources/kcat_small.png" height="40px" alt="kafkacat"> &nbsp;
  <img src="https://upload.wikimedia.org/wikipedia/commons/4/4e/Docker_%28container_engine%29_logo.svg" height="40px" alt="Docker">
</div>
<br />

### More

* Note: this is a complete project used at Portavita BV where . But not all is published on Github. If you want to know more about searching and/or indexing kafka messages, please contact me via niels/dot/koster/at/portavita/dot/nl
* Nice page about compiling statically in Go: https://www.arp242.net/static-go.html

