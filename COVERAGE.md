# Scenario coverage

This repository is an executable curriculum for the current Synthtraffic V1 DSL.
Docs: [DSL concepts](https://www.synthtraffic.io/docs/learn/scenario/) ·
[Expressions](https://www.synthtraffic.io/docs/expressions/syntax/) ·
[CLI](https://www.synthtraffic.io/docs/cli/sample/).

| V1 capability | Examples |
| --- | --- |
| Scopes, dependencies, stable instance fields | `core/vars_const.yaml`, `core/relationships.yaml`, `core/start_gap.yaml` |
| Whole-record and field references; retained-history bound | `core/relationships.yaml`, `core/bounded_history.yaml`, `healthcare/patient_encounters.yaml`, `lab/ecommerce/05-orders-ref-schedule.yaml` |
| Previous-event isolation | `core/previous.yaml`, `healthcare/bedside_observations.yaml`, `lab/ecommerce/07-previous-cart.yaml` |
| Fixed and weighted state transitions, terminal states | `core/lifecycles.yaml`, `healthcare/claims.yaml`, `lab/ecommerce/08-payments-lifecycle.yaml` |
| All V1 function families | [function map](functions/README.md), including `functions/comparisons.yaml` |
| Null, omission, uniqueness, cardinality, numeric modifiers | `functions/modifiers.yaml`, `lab/ecommerce/11-logic-comparisons.yaml` |
| Coherent person objects | `healthcare/patient_encounters.yaml`, `lab/ecommerce/02-faker-vars.yaml` |
| Nested arrays, projected aggregates | `functions/collections.yaml`, `healthcare/claims.yaml`, `lab/ecommerce/10-line-items-collections.yaml` |
| Custom faker packs | `functions/faker_pack.yaml` (+ `faker_pack_vehicle.csv`) |
| Rate, interval, windows, event/duration limits | [pacing index](pacing/README.md), `lab/ecommerce/04-rate-interval.yaml` |
| Schedule once, cycle, restart; stage batches | `core/parent_child.yaml`, `pacing/schedule_cycle.yaml`, `pacing/schedule_restart.yaml` |
| Delay, discard, repeat, inherited delivery | `pacing/delay.yaml`, `pacing/discard.yaml`, `pacing/repeat.yaml`, `pacing/delivery_inheritance.yaml` |
| Kafka key, headers, JSON/string/bytes serializers | `connectors/kafka.yaml`, `connectors/kafka_serializers.yaml` |
| Kafka plaintext, TLS + PLAIN, TLS + SCRAM | `connectors/kafka.yaml`, `connectors/kafka_cloud.yaml`, `connectors/kafka_scram.yaml` |
| PostgreSQL policies, constraints | `connectors/postgres.yaml`, `postgres/` |
| HTTP paths, query, headers, requests | `connectors/http.yaml`, [local receiver](http-server/README.md) |
| File and cloud storage | `connectors/file.yaml`, `connectors/s3.yaml`, `connectors/s3-minio.yaml`, `connectors/gcs.yaml`, `connectors/azure-blob.yaml` |
| Runtime overflow and exhausted uniqueness | [runtime-errors/](runtime-errors/README.md) |
| End-to-end healthcare domain | [healthcare/](healthcare/README.md) |
| End-to-end ecommerce domain | [ecommerce/](ecommerce/README.md) |
| Progressive learning path | [lab/ecommerce/](lab/ecommerce/) |

`sample` and `run --stdout` preview connector scenarios without contacting
destinations. Real connector auth, serializers on the wire, SQL constraints, and
cloud permissions need a live sink.
