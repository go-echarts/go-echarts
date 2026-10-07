#!/usr/bin/env bash
set -euo pipefail

ROOT=$(git rev-parse --show-toplevel)
SOURCEY_GODOC_COMMIT="5c699ac3ee2f251c083be8f9deb1b345c52c4a6d"
GO_ECHARTS_COMMIT="3f3cf9579987d520eef191d3f9b7d8c47d234276"
OUTPUT_DIR="$ROOT/docs/api"
WORK=$(mktemp -d)
MODULE_DIR="$WORK/go-echarts"

cleanup() {
  git -C "$ROOT" worktree remove --force "$MODULE_DIR" >/dev/null 2>&1 || true
  rm -rf "$WORK"
}
trap cleanup EXIT

command -v git >/dev/null
command -v go >/dev/null
command -v python3 >/dev/null

git -C "$ROOT" worktree add --detach "$MODULE_DIR" "$GO_ECHARTS_COMMIT"
SOURCE_TIMESTAMP=$(git -C "$MODULE_DIR" show -s --format=%cI "$GO_ECHARTS_COMMIT")

(
  cd "$WORK"
  go run "github.com/sourcey/sourcey/go/sourcey-godoc/cmd/sourcey-godoc@${SOURCEY_GODOC_COMMIT}" \
    generate \
    --module "$MODULE_DIR" \
    --packages ./... \
    --out "$OUTPUT_DIR" \
    --title "go-echarts v2.7.2 API Reference"
)

python3 "$ROOT/tools/sourcey-api/enrich.py" \
  "$OUTPUT_DIR" "$GO_ECHARTS_COMMIT" "$SOURCE_TIMESTAMP"
cp "$ROOT/tools/sourcey-api/LICENSE.sourcey.txt" "$OUTPUT_DIR/LICENSE.sourcey.txt"
cp "$ROOT/tools/sourcey-api/THIRD_PARTY_NOTICES.md" "$OUTPUT_DIR/THIRD_PARTY_NOTICES.md"
cp "$ROOT/tools/sourcey-api/BUILD.md" "$OUTPUT_DIR/BUILD.md"
python3 "$ROOT/tools/sourcey-api/verify-output.py" "$OUTPUT_DIR"
