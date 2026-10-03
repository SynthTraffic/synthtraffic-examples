# Core examples

Generators, scopes, relationships, history, and lifecycles.

| File | Covers |
| --- | --- |
| [`basic.yaml`](basic.yaml) | Minimal generator, faker, `concat` |
| [`vars_const.yaml`](vars_const.yaml) | `const`, `vars`, same-object `$` refs |
| [`parent_child.yaml`](parent_child.yaml) | Simple `ref` + schedule |
| [`relationships.yaml`](relationships.yaml) | Whole-record `ref`, instances, state emit |
| [`previous.yaml`](previous.yaml) | `previous` with instance isolation |
| [`lifecycles.yaml`](lifecycles.yaml) | State machines / terminal states |
| [`start_gap.yaml`](start_gap.yaml) | Staggered instance activation |
| [`bounded_history.yaml`](bounded_history.yaml) | Explicit retained-history limit |

```bash
synthtraffic sample core/basic.yaml --events 3 --seed 42
synthtraffic run core/parent_child.yaml --stdout --seed 42
```

Continue with [Learn: scenario](https://www.synthtraffic.io/docs/learn/scenario/).
