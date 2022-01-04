# Kafka Search UI Backend

For a limited or selected set of kafka topics, a data pipeline (spark) updates a postgres database with identifier
keys/values and partitions/offsets.

The kafkasearch-ui backend serves an API for a UI, that queries the database, and seeks kafka for the raw kafka message.

### Build

```
CGO_ENABLED=0 go build -ldflags="-extldflags=-static
docker build -t fhirstation-kafkasearch-ui .
```

The UI directory holds the React App (internal project, not yet on Github)

### Third-party Resources

<img src="https://gofiber.io/assets/images/logo.svg" height="40px" alt="Gofiber">

* Gofiber: https://gofiber.io/ https://github.com/gofiber/fiber
* Statically compiled kafkacat docker image from https://github.com/jcaesar/kafkacat-static
