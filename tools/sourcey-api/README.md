# Sourcey API reference build

This directory contains the reproducible build for `docs/api/`.

## Pinned inputs

- go-echarts source commit: `3f3cf9579987d520eef191d3f9b7d8c47d234276` (v2.7.2 plus workflow-only updates)
- sourcey-godoc release: `0.2.1`
- sourcey-godoc source commit: `5c699ac3ee2f251c083be8f9deb1b345c52c4a6d`
- sourcey-godoc Linux amd64 release asset SHA-256 used for the independently verified local build: `cb602e5e35e756c95fc3b8311109ba5a70e9f74afad1f8b0c8efc0022b683bee`
- Go: 1.21 or newer
- Python: 3.10 or newer

## Rebuild

```bash
tools/sourcey-api/build.sh
```

The script creates a detached worktree at the pinned go-echarts commit, runs sourcey-godoc from its exact source commit with `go run`, enriches every page with canonical and exact-commit source links, normalizes the generated timestamp to the source commit time, copies license/notices, and validates the output.

The generated site is committed under `docs/api/` because the existing Pages workflow uploads `docs/` directly.
