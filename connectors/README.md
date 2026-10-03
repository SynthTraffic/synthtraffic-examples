# Connector examples

Local sinks and cloud destinations. Preview with `sample` / `run --stdout`.
Omit needless top-level `defaults` — use CLI `--seed` for repeatable output.

## Local

| File | Sink | Notes |
| --- | --- | --- |
| [`kafka.yaml`](kafka.yaml) | Kafka | `localhost:9092` |
| [`kafka_serializers.yaml`](kafka_serializers.yaml) | Kafka | JSON, string, and bytes key/value serializers |
| [`postgres.yaml`](postgres.yaml) | PostgreSQL | needs `POSTGRES_PASSWORD`; schema policy demo |
| [`http.yaml`](http.yaml) | HTTP | companion [`../http-server`](../http-server/) |
| [`file.yaml`](file.yaml) | Local files | writes under `./output/` on a real `run` |
| [`s3-minio.yaml`](s3-minio.yaml) | S3-compatible | local MinIO / path-style endpoint |

Schema-policy edge cases: [`../postgres/`](../postgres/).

```bash
synthtraffic sample connectors/kafka_serializers.yaml --events 6 --seed 42
synthtraffic run connectors/file.yaml --stdout --events 5 --seed 42
```

## Cloud destinations

These scenarios already use `env(...)`. Create the topic / bucket / container
first. Do not commit credentials. `sample` and `run --stdout` preview without
contacting storage sinks; a real Kafka `run` produces for real.

### Kafka (TLS + SASL/PLAIN)

File: [`kafka_cloud.yaml`](kafka_cloud.yaml)

| Variable | Purpose |
| --- | --- |
| `KAFKA_BOOTSTRAP_SERVERS` | bootstrap host:port |
| `KAFKA_SASL_USERNAME` | SASL username |
| `KAFKA_SASL_PASSWORD` | SASL password |

```bash
export KAFKA_BOOTSTRAP_SERVERS='pkc-xxxxx.region.aws.confluent.cloud:9092'
export KAFKA_SASL_USERNAME='...'
export KAFKA_SASL_PASSWORD='...'
synthtraffic sample connectors/kafka_cloud.yaml --events 2 --seed 42
synthtraffic run connectors/kafka_cloud.yaml
```

### Kafka (TLS + SASL/SCRAM)

[`kafka_scram.yaml`](kafka_scram.yaml) uses the same three environment variables
with `scram-sha-256` (V1 also accepts `scram-sha-512`). Create the
`synthtraffic-scram` topic before a real run.

```bash
synthtraffic sample connectors/kafka_scram.yaml --events 2 --seed 42
```

### Amazon S3

[`s3.yaml`](s3.yaml) — AWS credential chain plus `S3_BUCKET` and `AWS_REGION`.
Local stand-in: [`s3-minio.yaml`](s3-minio.yaml).

### Google Cloud Storage

[`gcs.yaml`](gcs.yaml) — ADC / service account plus `GCS_BUCKET`.

### Azure Blob

[`azure-blob.yaml`](azure-blob.yaml) — connection string / account URL plus container env vars as documented in the file.

Connector reference: [Sinks](https://www.synthtraffic.io/docs/learn/sinks/).
