# Kafka Search UI Backend

The kafkasearch-ui backend serves an API for a UI, that queries the database, and consults kafka to fetch the raw kafka message as JSON.

Use case at Portavita BV: for a limited or selected set of kafka topics (FHIR resources), a data pipeline (spark) updates a postgres database with kafka message identifier key/value and partition/offset data.

### Build

```
git clone git@github.com:nkoster/kafkasearch-ui
cd kafkasearch-ui/
rsync -a <UI build dir>/ ui/
CGO_ENABLED=0 go build -ldflags="-extldflags=-static"
docker build -t fhirstation-kafkasearch-ui .
```

The `<UI build dir>` holds the search UI frontend React App (internal project, not yet on Github, read below)

### OpenID-Connect

The backend uses OIDC middleware, but that is currently still Work-in-Progress.
You can disable the OIDC middleware by using the environment variable `USE_AUTH=false`

### Deploy

You can deploy the docker image into your kubernetes machinery, we use Helm for that, but you can also run it locally.
In that case you need to provide configuration via "`.env`" and add the following to the `Dockerfile`:

```
COPY .env /.env
```

The following environment variables need to have proper values:

```
```

### Third-party Resources
                                                                                                                   
* Gofiber: https://gofiber.io/ https://github.com/gofiber/fiber
* Statically compiled kafkacat docker image from https://github.com/jcaesar/kafkacat-static
* Docker base: https://registry.hub.docker.com/_/busybox/
* Busybox: https://www.busybox.net

<br />
<div style="white-space:nowrap">
  <img src="https://gofiber.io/assets/images/logo.svg" height="40px" alt="Gofiber"> &nbsp;
  <img src="https://raw.githubusercontent.com/edenhill/kcat/master/resources/kcat_small.png" height="40px" alt="kafkacat"> &nbsp;
  <img src="https://upload.wikimedia.org/wikipedia/commons/4/4e/Docker_%28container_engine%29_logo.svg" height="40px" alt="Docker"> &nbsp;
  <img src="https://www.busybox.net/images/busybox1.png" height="40px" alt="Busybox">
</div>
<br />

### More Info

* This is a complete project used at Portavita BV. But not all is published on Github. If you want to know more about searching and/or indexing kafka context, because you persist stuff (too) long or even forever inside kafka, please contact me via niels/dot/koster/at/portavita/dot/nl
* Nice page about compiling statically in Go: https://www.arp242.net/static-go.html
