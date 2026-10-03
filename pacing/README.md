# Pacing examples

Rate, interval, windows, limits, delivery controls, and schedules.

| File | Covers |
| --- | --- |
| [`interval.yaml`](interval.yaml) | Generator `interval` |
| [`rate_windows.yaml`](rate_windows.yaml) | Cron rate windows |
| [`max_duration.yaml`](max_duration.yaml) | Duration-only finite rate |
| [`combined_limits.yaml`](combined_limits.yaml) | `maxEvents` vs `maxDuration` |
| [`loop_duration.yaml`](loop_duration.yaml) | Schedule cycle vs duration |
| [`schedule_cycle.yaml`](schedule_cycle.yaml) | Interleaved stage batches |
| [`schedule_restart.yaml`](schedule_restart.yaml) | Restarting batches with lifetime caps |
| [`delay.yaml`](delay.yaml) | Connector release delay |
| [`discard.yaml`](discard.yaml) | Probabilistic discard |
| [`repeat.yaml`](repeat.yaml) | Duplicate connector writes |
| [`delivery_inheritance.yaml`](delivery_inheritance.yaml) | Defaults delivery overridden per generator |

```bash
synthtraffic sample pacing/interval.yaml --events 5 --seed 42
synthtraffic run pacing/schedule_cycle.yaml --stdout --seed 42 --events 20
```

Time and delivery docs: [Time](https://www.synthtraffic.io/docs/learn/time/) ·
[Delivery](https://www.synthtraffic.io/docs/learn/delivery-behavior/).
