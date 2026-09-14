# Pacing examples

Control how fast events appear, how long a generator runs, and how connectors release or drop writes.

```bash
synthtraffic sample pacing/interval.yaml --events 5 --seed 42
synthtraffic run pacing/delay.yaml --stdout --seed 42
```

| File | Covers |
| --- | --- |
| [`interval.yaml`](interval.yaml) | Generator `interval` |
| [`rate_windows.yaml`](rate_windows.yaml) | Cron rate windows |
| [`max_duration.yaml`](max_duration.yaml) | Duration-only finite rate |
| [`combined_limits.yaml`](combined_limits.yaml) | `maxEvents` vs `maxDuration` |
| [`loop_duration.yaml`](loop_duration.yaml) | Schedule cycle vs duration |
| [`schedule_cycle.yaml`](schedule_cycle.yaml) | Interleaved stage batches |
| [`delay.yaml`](delay.yaml) | Connector release delay |
| [`discard.yaml`](discard.yaml) | Probabilistic discard |
| [`repeat.yaml`](repeat.yaml) | Duplicate connector writes |

Time and delivery docs: [Time](https://synthtraffic.dev/docs/learn/time/) · [Delivery](https://synthtraffic.dev/docs/learn/delivery-behavior/).
