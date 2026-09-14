# Connector examples

Local sinks and cloud destinations. Preview payloads with `sample` or `run --stdout` before you send for real.

Do not commit credentials. Cloud files already use `env(...)`.

## Local

| File | Sink | Notes |
| --- | --- | --- |
| [`kafka.yaml`](kafka.yaml) | Kafka | `localhost:9092` |
| [`postgres.yaml`](postgres.yaml) | PostgreSQL | needs `POSTGRES_PASSWORD`; `drop-and-create` schema policy |
| [`http.yaml`](http.yaml) | HTTP | start [`../http-server`](../http-server/) first |
| [`file.yaml`](file.yaml) | Local files | writes under `./output/` on a real `run` |
| [`s3-minio.yaml`](s3-minio.yaml) | S3-compatible | local MinIO / path-style endpoint |

Schema-policy edge cases: [`../postgres/`](../postgres/).

Preview without opening the sink:

```bash
synthtraffic run connectors/file.yaml --stdout --events 5 --seed 42
synthtraffic run connectors/http.yaml --stdout --events 3 --seed 42
```

## Cloud destinations

Create the topic / bucket / container first. `sample` and `run --stdout` preview payloads without contacting storage sinks. A real Kafka `run` produces for real.

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

Managed clusters usually require the topic (`synthtraffic-customers` in the example) to exist already.

### Amazon S3

File: [`s3.yaml`](s3.yaml) — AWS credential chain (`AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY` / profile / IAM role) plus:

| Variable | Purpose |
| --- | --- |
| `S3_BUCKET` | destination bucket |
| `AWS_REGION` | bucket region |

```bash
export S3_BUCKET=my-bucket
export AWS_REGION=us-east-1
synthtraffic sample connectors/s3.yaml --events 3
synthtraffic run connectors/s3.yaml
```

For a local stand-in, use [`s3-minio.yaml`](s3-minio.yaml).

### Google Cloud Storage

File: [`gcs.yaml`](gcs.yaml) — Application Default Credentials or a service account, plus:

| Variable | Purpose |
| --- | --- |
| `GCS_BUCKET` | destination bucket |

```bash
export GCS_BUCKET=my-gcs-bucket
synthtraffic sample connectors/gcs.yaml --events 3
synthtraffic run connectors/gcs.yaml
```

### Azure Blob Storage

File: [`azure-blob.yaml`](azure-blob.yaml) — Azure SDK default credential chain, plus:

| Variable | Purpose |
| --- | --- |
| `AZURE_STORAGE_ACCOUNT_URL` | e.g. `https://ACCOUNT.blob.core.windows.net` |
| `AZURE_STORAGE_CONTAINER` | container name |

```bash
export AZURE_STORAGE_ACCOUNT_URL='https://ACCOUNT.blob.core.windows.net'
export AZURE_STORAGE_CONTAINER=synthtraffic
synthtraffic sample connectors/azure-blob.yaml --events 3
synthtraffic run connectors/azure-blob.yaml
```

Connector reference: [Docs](https://synthtraffic.dev/docs/learn/sinks/).
