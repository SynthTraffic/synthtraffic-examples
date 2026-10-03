# Ecommerce DSL lab

Progressive shop scenarios that walk the V1 DSL end to end. Work one file at a
time. Early exercises are value-only (preview with `sample`). Later ones add
Kafka, PostgreSQL, HTTP, and file sinks — still safe to `sample` without opening
destinations.

## Story

| Party | Role |
| --- | --- |
| **customers** | Shoppers who register and place orders |
| **products** | Catalog items (SKU, price, category) |
| **orders** | Purchases that reference a customer |
| **payments** | Payment attempts tied to an order |
| **shipments** | Fulfillment events after payment |
| **notifications** | Delivery-controlled side effects |

## Exercises

| # | File | Focus |
| --- | --- | --- |
| 01 | [`01-customers.yaml`](01-customers.yaml) | First generator, `uuid`, `now` |
| 02 | [`02-faker-vars.yaml`](02-faker-vars.yaml) | Faker, `vars`, `const`, clock |
| 03 | [`03-products.yaml`](03-products.yaml) | `seq(pattern=)`, `oneOf`, `weightedOneOf`, `float` |
| 04 | [`04-rate-interval.yaml`](04-rate-interval.yaml) | `rate`, `interval`, `maxDuration` |
| 05 | [`05-orders-ref-schedule.yaml`](05-orders-ref-schedule.yaml) | `ref`, schedule stages, history |
| 06 | [`06-instances-start-gap.yaml`](06-instances-start-gap.yaml) | Instances, stable fields, `startGap` |
| 07 | [`07-previous-cart.yaml`](07-previous-cart.yaml) | `previous(field)` per instance |
| 08 | [`08-payments-lifecycle.yaml`](08-payments-lifecycle.yaml) | Weighted `stateMachine` |
| 09 | [`09-shipments-lifecycle.yaml`](09-shipments-lifecycle.yaml) | Multi-step fixed lifecycle |
| 10 | [`10-line-items-collections.yaml`](10-line-items-collections.yaml) | Arrays, aggregates, `selectKeys` |
| 11 | [`11-logic-comparisons.yaml`](11-logic-comparisons.yaml) | `case`, comparisons, modifiers |
| 12 | [`12-time-and-schedules.yaml`](12-time-and-schedules.yaml) | `timestamp`, `addDuration`, cycle schedule |
| 13 | [`13-delivery-controls.yaml`](13-delivery-controls.yaml) | Delay, discard, repeat |
| 14 | [`14-kafka-orders.yaml`](14-kafka-orders.yaml) | Kafka connection shape |
| 15 | [`15-postgres-customers.yaml`](15-postgres-customers.yaml) | PostgreSQL schema policy + rows |
| 16 | [`16-http-webhooks.yaml`](16-http-webhooks.yaml) | HTTP path/query/headers/body |
| 17 | [`17-file-export.yaml`](17-file-export.yaml) | Local file sink + rolling |
| 18 | [`18-marketplace-day.yaml`](18-marketplace-day.yaml) | Capstone: customers → orders → payments |

## Commands

```bash
export SYNTHTRAFFIC_LICENSE_FILE=./license.env

# Fast preview (no sinks, no real-time pacing)
synthtraffic sample lab/ecommerce/01-customers.yaml --events 5 --seed 42

# Dry-run with runtime scheduling
synthtraffic run lab/ecommerce/05-orders-ref-schedule.yaml --stdout --seed 42

# Capstone
synthtraffic run lab/ecommerce/18-marketplace-day.yaml --stdout --seed 42

# Studio on the lab folder
synthtraffic studio --folder lab/ecommerce --port 8787
```

Connector exercises (14–17) are preview-safe with `sample` / `run --stdout`.
Real delivery needs the matching local service (Kafka, PostgreSQL, the
[`http-server`](../../http-server/) companion, or a writable `./output` directory).

When you finish the lab, explore the polished domain packs
[`ecommerce/`](../../ecommerce/) and [`healthcare/`](../../healthcare/), plus the
focused catalogs under [`core/`](../../core/), [`functions/`](../../functions/),
[`pacing/`](../../pacing/), and [`connectors/`](../../connectors/).
