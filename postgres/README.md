# PostgreSQL schema examples

Focused scenarios for schema policies and constraint behavior. Use with a local
PostgreSQL instance and `POSTGRES_PASSWORD`.

| File | Covers |
| --- | --- |
| [`constraints.yaml`](constraints.yaml) | Declared column constraints |
| [`create_if_missing_additive.yaml`](create_if_missing_additive.yaml) | Additive `create-if-missing` |
| [`manual_existing_table.yaml`](manual_existing_table.yaml) | `manual` policy (no DDL) |
| [`foreign_key_violation.yaml`](foreign_key_violation.yaml) | Server-side FK rejection |
| [`not_null_violation.yaml`](not_null_violation.yaml) | NOT NULL rejection |
| [`wrong_type.yaml`](wrong_type.yaml) | Type mismatch rejection |

```bash
export POSTGRES_PASSWORD=postgres
synthtraffic sample postgres/constraints.yaml --events 2 --seed 42
# Real insert (needs a running database):
# synthtraffic run postgres/constraints.yaml
```

Schema policy reference: [PostgreSQL](https://www.synthtraffic.io/docs/connectors/postgresql/).
