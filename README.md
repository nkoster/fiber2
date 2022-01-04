Kafka search UI backend.

for a limited/selected set of kafka topics, a data pipeline (spark) updates a postgres database with identifier
keys/values and partitions/offsets.

This program serves a UI with an API that queries the database, and seeks for the raw kafka message.

Third-party resources:
* Statically compiled kafkacat docker image from https://github.com/jcaesar/kafkacat-static
