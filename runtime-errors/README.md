# Expected runtime failures

These files compile, but a complete run must fail.

| Scenario | Expected failure |
| --- | --- |
| [`history_overflow.yaml`](history_overflow.yaml) | Required parent history exceeds `history.maxEntries` |
| [`unique_exhaustion.yaml`](unique_exhaustion.yaml) | `unique=true` exhausts a finite domain |

```bash
synthtraffic run runtime-errors/history_overflow.yaml --stdout
synthtraffic run runtime-errors/unique_exhaustion.yaml --stdout
```

Both should exit unsuccessfully after partial output. A short `sample` may stop
before the failure.
