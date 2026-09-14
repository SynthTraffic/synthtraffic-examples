# Core examples

Generators, scopes, relationships, and lifecycles. Preview with:

```bash
synthtraffic sample core/basic.yaml --events 3 --seed 42
synthtraffic run core/parent_child.yaml --stdout --seed 42
```

| File | Covers |
| --- | --- |
| [`basic.yaml`](basic.yaml) | One generator, faker, `concat` |
| [`vars_const.yaml`](vars_const.yaml) | `const`, `vars`, same-object `$` refs |
| [`parent_child.yaml`](parent_child.yaml) | `ref` + `schedule` (customers then orders) |
| [`relationships.yaml`](relationships.yaml) | Whole-record `ref`, instances, state emit |
| [`previous.yaml`](previous.yaml) | `previous` with instance isolation |
| [`lifecycles.yaml`](lifecycles.yaml) | State machines and terminal states |
| [`start_gap.yaml`](start_gap.yaml) | Staggered instance activation |

Continue with [Learn: scenario](https://synthtraffic.dev/docs/learn/scenario/).
