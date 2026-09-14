# PostgreSQL schema examples

These files show how Synthtraffic talks to PostgreSQL: create tables, add columns, or leave the schema alone. They need a reachable database and `POSTGRES_PASSWORD` or `POSTGRES_URL` (see each file).

The happy-path sink example is [`../connectors/postgres.yaml`](../connectors/postgres.yaml).

| File | What it demonstrates |
| --- | --- |
| [`create_if_missing_additive.yaml`](create_if_missing_additive.yaml) | `create-if-missing` adds missing columns |
| [`constraints.yaml`](constraints.yaml) | Declared constraints on generated tables |
| [`manual_existing_table.yaml`](manual_existing_table.yaml) | `manual` — Synthtraffic does not change table structure |
| [`foreign_key_violation.yaml`](foreign_key_violation.yaml) | Expected failure when generated rows break a foreign key |
| [`not_null_violation.yaml`](not_null_violation.yaml) | Expected failure on NOT NULL |
| [`wrong_type.yaml`](wrong_type.yaml) | Expected failure when a value does not match the column type |

The last three files are meant to fail at run time so you can see the diagnostic. Preview payloads without touching the database:

```bash
synthtraffic run postgres/constraints.yaml --stdout --events 2 --seed 42
```

Schema policy reference: [PostgreSQL](https://synthtraffic.dev/docs/connectors/postgresql/).
