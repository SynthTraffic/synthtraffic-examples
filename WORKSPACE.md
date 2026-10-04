# Working on this repo from the Cloud Agent

This checkout is the **examples** repository. In a Cloud Agent it lives at
`/workspace`. The install script downloads the CLI binary; you do not need the
engine repo to preview scenarios here.

## Branch

Feature work uses `cursor/<name>-be76`. Push via GitHub (MCP / your credentials);
the Cloud Agent `cursor[bot]` token cannot push to this public examples repo.

## Preview locally in the agent

After the environment install completes, the CLI is on your PATH. If license
secrets are configured in the environment, the wrapper writes a license file
under `~/.config/` automatically — no manual export needed.

See [`README.md`](README.md) for preview commands, for example:

```bash
cd /workspace
# CLI on PATH after install — see README for full syntax
sample lab/ecommerce/18-marketplace-day.yaml --events 20 --seed 42
```

To point at your own license file, set the license-file environment variable
documented in [`README.md`](README.md) before running `sample` or `run`.

## Lab story

See [`lab/ecommerce/README.md`](lab/ecommerce/README.md) and the manager script
[`lab/ecommerce/DEMO.md`](lab/ecommerce/DEMO.md).
