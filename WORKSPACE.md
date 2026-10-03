# Working on this repo from the Synthtraffic agent

This checkout lives at `/workspace/synthtraffic-examples` beside the main
`synthtraffic` engine repo. Use it for ongoing lab/example edits.

## Branch

Feature work uses `cursor/<name>-97b7` (current: `cursor/lab-ecommerce-story-97b7`).
Push via GitHub (MCP / your credentials); the Cloud Agent `cursor[bot]` token
cannot push to this public examples repo.

## Preview locally in the agent

```bash
cd /workspace/synthtraffic-examples
export SYNTHTRAFFIC_LICENSE_DEV=1
export SYNTHTRAFFIC_LICENSE_FILE=/workspace/testdata/license/valid.env
/workspace/bin/synthtraffic sample lab/ecommerce/18-marketplace-day.yaml --events 20 --seed 42
```

Rebuild the CLI after engine changes:

```bash
cd /workspace && go build -o bin/synthtraffic ./cmd/synthtraffic
```

## Lab story

See [`lab/ecommerce/README.md`](lab/ecommerce/README.md) and the manager script
[`lab/ecommerce/DEMO.md`](lab/ecommerce/DEMO.md).
