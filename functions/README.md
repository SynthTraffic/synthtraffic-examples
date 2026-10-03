# Function examples

Focused YAML demos for expression families. New files omit `version` and
needless top-level `defaults`.

## Map (function → file)

| Family | File |
| --- | --- |
| `seq`, `uuid`, `oneOf`, `cycle`, `float`, `normal`, `bool`, `now`, `faker` | [`expressions.yaml`](expressions.yaml) |
| `array`, nested scopes, aggregates intro | [`collections.yaml`](collections.yaml) |
| `easing`, `easingChain` | [`easing.yaml`](easing.yaml) |
| `ulid`, `int`, `float`, `string`, `bytes`, `weightedOneOf` | [`ids_random.yaml`](ids_random.yaml) |
| `merge`, `selectKeys`, `omitKeys`, `range`, `sum` / `count` / `avg` | [`structure.yaml`](structure.yaml) |
| `eq`/`gte`/`and`, `case`, `formatString`, `concat`, `env` | [`logic_strings.yaml`](logic_strings.yaml) |
| All six comparisons and `and`/`or`/`not` | [`comparisons.yaml`](comparisons.yaml) |
| `nullable`, `optional`, `unique`, `cardinality`, numeric modifiers | [`modifiers.yaml`](modifiers.yaml) |
| `abs`, `floor`, `ceil` | [`numeric_helpers.yaml`](numeric_helpers.yaml) |
| `dateBetween`, `dateTimeBetween`, `formatDateTime`, `randomDuration` | [`time.yaml`](time.yaml) |
| `addDuration`, `subtractDuration` | [`timestamp_arithmetic.yaml`](timestamp_arithmetic.yaml) |
| `timestamp`, `clockStart` | [`timestamp_sequence.yaml`](timestamp_sequence.yaml) |
| `add`/`subtract`/`multiply`/`divide`/`pow`/`round`/`min`/`max` | [`math.yaml`](math.yaml) |
| Locale-aware `faker` | `faker_en_us.yaml`, `faker_en_in.yaml`, `faker_en_au.yaml`, `faker_en_gb.yaml` |
| Custom faker CSV packs | [`faker_pack.yaml`](faker_pack.yaml) |

`ref` and `previous` live under [`../core/`](../core/). Full expression reference:
[Docs](https://www.synthtraffic.io/docs/expressions/syntax/).
