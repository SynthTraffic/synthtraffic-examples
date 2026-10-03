# ShopLane lab — one storefront day

Progressive scenarios that tell **one ecommerce story**. Each file adds a
capability you need for the next beat. Early files are value-only (`sample`).
Later files add Kafka, PostgreSQL, HTTP, and files — still safe to preview.

## The story

ShopLane opens for the day: catalog lists, shoppers sign up, traffic arrives,
orders join real customers and SKUs, carts update, payments capture or fail,
shipments move, then the same shapes leave through real sinks.

| Party | Role in the day |
| --- | --- |
| **products** | SKUs and prices listed before checkout |
| **customers** | Shoppers with coherent faker identity |
| **page_views / carts** | Browse and cart activity |
| **orders** | Purchases that `ref` customers (and products) |
| **payments** | Capture / fail lifecycle on an order |
| **shipments** | Packed → out for delivery → delivered |
| **notifications** | Imperfect delivery (delay / discard / repeat) |

Shared fixtures across the lab: `seed: 42`, clock `2026-01-15T10:00:00Z`,
`store: ShopLane`, id patterns `CUST-` / `SKU-` / `ORD-` / `PAY-` / `SHP-`.

## Exercises

| # | File | Story beat | DSL focus |
| --- | --- | --- | --- |
| 01 | [`01-customers.yaml`](01-customers.yaml) | First signups | generator, `uuid`, `now` |
| 02 | [`02-faker-vars.yaml`](02-faker-vars.yaml) | Realistic shoppers | faker, `vars`, **used** `const`, clock |
| 03 | [`03-products.yaml`](03-products.yaml) | Catalog live | `seq`, `oneOf`, `weightedOneOf` |
| 04 | [`04-browse-traffic.yaml`](04-browse-traffic.yaml) | Morning traffic | `rate`, `interval` |
| 05 | [`05-orders-ref-schedule.yaml`](05-orders-ref-schedule.yaml) | Orders join people + SKUs | `ref`, schedule, history |
| 06 | [`06-instances-start-gap.yaml`](06-instances-start-gap.yaml) | Carts open | instances, `startGap` |
| 07 | [`07-previous-cart.yaml`](07-previous-cart.yaml) | Cart deltas | `previous` |
| 08 | [`08-payments-lifecycle.yaml`](08-payments-lifecycle.yaml) | Pay for orders | `stateMachine` + `ref` |
| 09 | [`09-shipments-lifecycle.yaml`](09-shipments-lifecycle.yaml) | Fulfill orders | multi-step lifecycle + `ref` |
| 10 | [`10-line-items-collections.yaml`](10-line-items-collections.yaml) | Multi-line baskets | arrays, aggregates |
| 11 | [`11-logic-comparisons.yaml`](11-logic-comparisons.yaml) | Risk review | `case`, modifiers |
| 12 | [`12-time-and-schedules.yaml`](12-time-and-schedules.yaml) | Flash promos | `timestamp`, cycle schedule |
| 13 | [`13-delivery-controls.yaml`](13-delivery-controls.yaml) | Notify shoppers | delay / discard / repeat |
| 14 | [`14-kafka-orders.yaml`](14-kafka-orders.yaml) | Stream orders | Kafka |
| 15 | [`15-postgres-customers.yaml`](15-postgres-customers.yaml) | Persist shoppers | PostgreSQL |
| 16 | [`16-http-webhooks.yaml`](16-http-webhooks.yaml) | Partner webhook | HTTP |
| 17 | [`17-file-export.yaml`](17-file-export.yaml) | Batch export | file sink |
| 18 | [`18-marketplace-day.yaml`](18-marketplace-day.yaml) | Full day | capstone |

Manager-facing walkthrough: [`DEMO.md`](DEMO.md).

## Commands

```bash
export SYNTHTRAFFIC_LICENSE_FILE=./license.env

synthtraffic sample lab/ecommerce/01-customers.yaml --events 5 --seed 42
synthtraffic run lab/ecommerce/05-orders-ref-schedule.yaml --stdout --seed 42
synthtraffic run lab/ecommerce/18-marketplace-day.yaml --stdout --seed 42
synthtraffic studio --folder lab/ecommerce --port 8787
```

From this Cloud Agent workspace (dev license + built CLI):

```bash
export SYNTHTRAFFIC_LICENSE_DEV=1
export SYNTHTRAFFIC_LICENSE_FILE=/workspace/testdata/license/valid.env
/workspace/bin/synthtraffic sample lab/ecommerce/02-faker-vars.yaml --events 3 --seed 42
```
