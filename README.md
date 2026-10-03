# Synthtraffic examples

Runnable YAML scenarios for [Synthtraffic](https://www.synthtraffic.io). Clone
this repo, point the CLI or Docker image at a file, and preview events on your
machine.

Docs: [Install](https://www.synthtraffic.io/docs/getting-started/install/) ·
[Quickstart](https://www.synthtraffic.io/docs/getting-started/quickstart/) ·
[Learn](https://www.synthtraffic.io/docs/learn/scenario/)

Binaries: [GitHub Releases](https://github.com/SynthTraffic/synthtraffic-releases) ·
Image: [`synthtraffic/synthtraffic`](https://hub.docker.com/r/synthtraffic/synthtraffic)

## 1. Install and set a license

`sample`, `run`, and `studio` need a `license.env`. Get one from
[Pricing](https://www.synthtraffic.io/pricing/), then:

```bash
export SYNTHTRAFFIC_LICENSE_FILE=./license.env
synthtraffic sample core/basic.yaml --events 3 --seed 42
```

Docker (pass the license on every `docker run`):

```bash
docker run --rm --env-file license.env \
  -v "$(pwd)/core/basic.yaml:/work/scenario.yaml:ro" \
  synthtraffic/synthtraffic:1.0.0 \
  sample /work/scenario.yaml --events 3 --seed 42
```

PowerShell: `$env:SYNTHTRAFFIC_LICENSE_FILE = ".\license.env"` and `.\synthtraffic.exe`.

`--seed 42` repeats random choices. Do not put license files or cloud credentials
in scenario YAML — connectors that need secrets already use `env(...)`.

## 2. Preview vs send

| Command | What it does |
| --- | --- |
| `synthtraffic sample <file> --events 3 --seed 42` | Fast preview. No connectors, no rate pacing. |
| `synthtraffic run <file> --stdout --events 5` | Dry-run payloads to the terminal. |
| `synthtraffic run <file>` | Real delivery to the connection in the file. |
| `synthtraffic studio --folder . --port 8787` | Open the folder in Synthstudio. |

Studio in Docker:

```bash
docker run --rm -p 8787:8787 --env-file license.env \
  -v "$(pwd):/work" -w /work \
  synthtraffic/synthtraffic:1.0.0 \
  studio --port 8787 --folder /work
```

Then open `http://127.0.0.1:8787` and load a file under `/work`.

## Layout

| Folder | What it teaches |
| --- | --- |
| [`lab/ecommerce/`](lab/ecommerce/) | ShopLane story lab (18 steps) + [manager demo script](lab/ecommerce/DEMO.md) |
| [`core/`](core/) | Generators, `const` / `vars`, parent–child `ref`, relationships, `previous`, lifecycles, bounded history |
| [`functions/`](functions/) | Expression families (`uuid`, `faker`, `case`, time, math, collections, modifiers, …) |
| [`pacing/`](pacing/) | Interval, rate windows, duration limits, delay, discard, repeat, schedules |
| [`connectors/`](connectors/) | Kafka (incl. serializers / SCRAM), PostgreSQL, HTTP, files, S3, GCS, Azure Blob |
| [`postgres/`](postgres/) | Schema policies and constraint behavior |
| [`ecommerce/`](ecommerce/) | Polished customers / orders / payments domain pack |
| [`healthcare/`](healthcare/) | Patient encounters, bedside observations, claims |
| [`runtime-errors/`](runtime-errors/) | Intentional history overflow and uniqueness exhaustion |
| [`http-server/`](http-server/) | Local companion for [`connectors/http.yaml`](connectors/http.yaml) |

Start with [`lab/ecommerce/`](lab/ecommerce/) or [`core/basic.yaml`](core/basic.yaml),
then [`core/parent_child.yaml`](core/parent_child.yaml), then a connector under
[`connectors/`](connectors/).

## Core

| File | Covers |
| --- | --- |
| [`basic.yaml`](core/basic.yaml) | One generator, faker, `concat`, field order |
| [`vars_const.yaml`](core/vars_const.yaml) | `const`, `vars`, same-object `$` refs |
| [`parent_child.yaml`](core/parent_child.yaml) | Customers then orders via `ref` + `schedule` |
| [`relationships.yaml`](core/relationships.yaml) | Whole-record `ref`, instances, state emit |
| [`previous.yaml`](core/previous.yaml) | `previous` with instance isolation |
| [`lifecycles.yaml`](core/lifecycles.yaml) | State machines and terminal states |
| [`start_gap.yaml`](core/start_gap.yaml) | Staggered instance activation |
| [`bounded_history.yaml`](core/bounded_history.yaml) | Explicit retained-history limit |

## Style

- YAML-first. Omit `version` (compiler default is `1`).
- Prefer generator `config.maxEvents` and CLI `--seed` over needless top-level defaults.
- Use `seq(pattern=...)` (not `format=`).
- Secrets only through `env(...)`.

See [`COVERAGE.md`](COVERAGE.md) for the capability → file map.

## License

Scenario files in this repository are MIT licensed. Synthtraffic itself is a
licensed product — see [synthtraffic.io](https://www.synthtraffic.io).
