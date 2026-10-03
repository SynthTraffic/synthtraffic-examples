# Ecommerce domain pack

Marketplace vocabulary: customers, orders, and payments. Seed `42` and a fixed
UTC clock keep output stable.

| Scenario | Behavior |
| --- | --- |
| [`customers.yaml`](customers.yaml) | Coherent `faker(person)` fields, const metadata |
| [`orders.yaml`](orders.yaml) | Parent customers, child orders with whole-record `ref` |
| [`payments.yaml`](payments.yaml) | Payment instances with weighted capture/fail lifecycle |

```bash
synthtraffic run ecommerce/customers.yaml --stdout --seed 42
synthtraffic run ecommerce/orders.yaml --stdout --seed 42
synthtraffic run ecommerce/payments.yaml --stdout --seed 42
```

For a progressive learning path, start with [`../lab/ecommerce/`](../lab/ecommerce/).
