# Function examples

One family per file. Preview with:

```bash
synthtraffic sample functions/ids_random.yaml --events 3 --seed 42
```

| Family | File |
| --- | --- |
| `seq`, `uuid`, `oneOf`, `cycle`, `uniform`, `normal`, `bool`, `now`, `faker` | [`expressions.yaml`](expressions.yaml) |
| `array`, nested scopes, aggregates | [`collections.yaml`](collections.yaml) |
| `easing`, `easingChain` | [`easing.yaml`](easing.yaml) |
| `ulid`, `int`, `float`, `string`, `bytes`, `categorical` | [`ids_random.yaml`](ids_random.yaml) |
| `merge`, `selectKeys`, `omitKeys`, `range`, `sum` / `count` / `avg` | [`structure.yaml`](structure.yaml) |
| `eq` / `gte` / `and`, `case`, `format`, `concat`, `env` | [`logic_strings.yaml`](logic_strings.yaml) |
| `date`, `datetime`, `formatDateTime`, `uniformDuration` | [`time.yaml`](time.yaml) |
| `add` / `subtract` / `multiply` / `divide` / `pow` / `round` / `min` / `max` | [`math.yaml`](math.yaml) |
| Locale-aware `faker` | [`faker_en_us.yaml`](faker_en_us.yaml), [`faker_en_in.yaml`](faker_en_in.yaml), [`faker_en_au.yaml`](faker_en_au.yaml), [`faker_en_gb.yaml`](faker_en_gb.yaml) |

`ref` and `previous` live under [`../core/`](../core/). Full expression reference: [Docs](https://synthtraffic.dev/docs/expressions/syntax/).
