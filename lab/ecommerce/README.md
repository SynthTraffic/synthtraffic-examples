# Ecommerce lab

Short progressive scenarios for a shop: customers first, then richer faker fields. Work one file at a time. These files have no connectors yet — preview with `sample`.

## Story

| Party | Role |
| --- | --- |
| **customers** | Shoppers who register and place orders |
| **products** | Catalog items (later) |
| **orders** | Purchases that reference a customer (later) |

## Exercises

| # | File | Focus |
| --- | --- | --- |
| 01 | [`01-customers.yaml`](01-customers.yaml) | First generator, `uuid`, `now` |
| 02 | [`02-faker-vars.yaml`](02-faker-vars.yaml) | Faker, `vars`, `const` |

```bash
synthtraffic sample lab/ecommerce/01-customers.yaml --events 5 --seed 42
synthtraffic sample lab/ecommerce/02-faker-vars.yaml --events 5 --seed 42
```

When you are ready to send traffic, copy a sink from [`../../connectors/`](../../connectors/) and add a `connection` on the generator. Next lessons (rate, products, Kafka, PostgreSQL, `ref`, schedule, lifecycles) follow the same shop vocabulary in the [Learn](https://synthtraffic.dev/docs/learn/scenario/) docs.
