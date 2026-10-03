# Manager demo — ShopLane story (15–20 min)

Do **not** lead with the marketing homepage. Walk the lab: YAML in → related,
repeatable storefront traffic out.

## Goal

Show why Synthtraffic is useful: replace brittle generator scripts with a
readable scenario that keeps **identity**, **relationships**, **state over
time**, and optional **real sinks**.

## Prep

```bash
cd /path/to/synthtraffic-examples
export SYNTHTRAFFIC_LICENSE_FILE=./license.env   # production license.env
# or in this agent workspace:
export SYNTHTRAFFIC_LICENSE_DEV=1
export SYNTHTRAFFIC_LICENSE_FILE=/workspace/testdata/license/valid.env
BIN=synthtraffic   # or /workspace/bin/synthtraffic
```

Have a terminal open on `lab/ecommerce/`. Optional: Studio on that folder.

## Agenda

| Time | Beat | File / command |
| --- | --- | --- |
| 0:00–2:00 | Problem | “We need storefront-shaped traffic for pipelines — related, stateful, repeatable.” |
| 2:00–4:00 | Skeleton → realism | `01` then `02` |
| 4:00–7:00 | Catalog + join | `03` then `05` |
| 7:00–11:00 | State over time | `08` payments, `09` shipments |
| 11:00–14:00 | Whole day | `18` capstone |
| 14:00–16:00 | Delivery (optional) | `14` Kafka sample envelope, or mention 15–17 |
| 16:00–18:00 | Ask | Pick one staging pipeline to drive from a scenario |

## Script

### 1. Problem (≈2 min)

**Say:** Ad-hoc scripts and static fixtures fall down when you need the same
customer on many orders, payment status changes, and a re-runable demo. Synthtraffic
is a scenario engine: describe the day in YAML, get deterministic traffic.

### 2. From blank signup to real shoppers (≈2 min)

```bash
$BIN sample lab/ecommerce/01-customers.yaml --events 3 --seed 42
$BIN sample lab/ecommerce/02-faker-vars.yaml --events 3 --seed 42
```

**Point at:** `store` / `source` / `country` coming from `const` (used in the
payload), coherent `faker(person)` fields, stable `CUST-00x` ids, fixed clock.

**Say:** Same seed → same shoppers. That is how demos and CI stay honest.

### 3. Catalog, then orders that join (≈3 min)

```bash
$BIN sample lab/ecommerce/03-products.yaml --events 4 --seed 42
$BIN run lab/ecommerce/05-orders-ref-schedule.yaml --stdout --seed 42
```

**Point at:** products and customers first; each order carries `customerId` and
`productId` from `ref` — not random joins after the fact.

### 4. Pay and ship (≈4 min)

```bash
$BIN run lab/ecommerce/08-payments-lifecycle.yaml --stdout --seed 42
$BIN run lab/ecommerce/09-shipments-lifecycle.yaml --stdout --seed 42
```

**Say:** Same payment id moves authorized → captured/failed. Shipments walk
packed → out for delivery → delivered. Pipelines that key on entity id can be
tested without production data.

### 5. Capstone — the day in one file (≈3 min)

```bash
$BIN run lab/ecommerce/18-marketplace-day.yaml --stdout --seed 42
```

**Say:** One scenario: catalog + shoppers → orders → payments and shipments.
Re-run with `--seed 42` for the same story.

### 6. Optional sink beat (≈2 min)

```bash
$BIN sample lab/ecommerce/14-kafka-orders.yaml --events 2 --seed 42
```

**Say:** Preview never opens Kafka. When you are ready, the same shape produces
for real — also PostgreSQL, HTTP, files.

### 7. Close

**Ask:** Which staging topic/table should we drive next from a ShopLane-style
scenario?

## Cut list if short on time

Keep **02 → 05 → 08 → 18**. Drop sink slides and browse/cart deep dives.
