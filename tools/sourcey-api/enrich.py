#!/usr/bin/env python3
from __future__ import annotations

import html
import json
import re
import sys
from datetime import datetime, timezone
from pathlib import Path

MODULE = "github.com/go-echarts/go-echarts/v2"
SITE_BASE = "https://go-echarts.github.io/go-echarts/api/"
SOURCE_BASE = "https://github.com/go-echarts/go-echarts"
SOURCEY_COMMIT = "5c699ac3ee2f251c083be8f9deb1b345c52c4a6d"
GENERATOR = "sourcey-godoc 0.2.1"


def utc_timestamp(value: str) -> str:
    return (
        datetime.fromisoformat(value.replace("Z", "+00:00"))
        .astimezone(timezone.utc)
        .strftime("%Y-%m-%dT%H:%M:%SZ")
    )


def add_metadata(page: Path, source_url: str, source_label: str, commit: str) -> None:
    text = page.read_text(encoding="utf-8")
    canonical = SITE_BASE + page.name
    head = (
        f'<meta name="generator" content="{GENERATOR}">\n'
        f'<link rel="canonical" href="{html.escape(canonical, quote=True)}">\n'
    )
    assert "</head>" in text and 'name="generator"' not in text
    text = text.replace("</head>", head + "</head>", 1)

    source = (
        '<p class="source-meta">'
        f'Generated with <a href="https://github.com/sourcey/sourcey/tree/{SOURCEY_COMMIT}/go/sourcey-godoc">'
        f'{GENERATOR}</a> · '
        f'<a href="{html.escape(source_url, quote=True)}">'
        f'Source: {html.escape(source_label)}</a> at '
        f'<code>{commit[:12]}</code></p>'
    )
    pattern = r"(<main><h1>.*?</h1><p class=\"(?:module|import-path)\">.*?</p>)"
    replaced, count = re.subn(pattern, r"\1" + source, text, count=1, flags=re.S)
    assert count == 1, page
    page.write_text(replaced, encoding="utf-8")


def main() -> None:
    output = Path(sys.argv[1]).resolve()
    commit = sys.argv[2]
    timestamp = utc_timestamp(sys.argv[3])
    snapshot_path = output / "sourcey-godoc.json"
    snapshot = json.loads(snapshot_path.read_text())
    assert snapshot["source"] == "sourcey-godoc"
    assert snapshot["module_path"] == MODULE
    snapshot["generated_at"] = timestamp
    snapshot_path.write_text(json.dumps(snapshot, indent=2) + "\n")

    add_metadata(
        output / "index.html",
        f"{SOURCE_BASE}/tree/{commit}",
        "go-echarts v2.7.2",
        commit,
    )
    for package in snapshot["packages"]:
        name = package["name"]
        directory = package["dir"]
        add_metadata(
            output / f"pkg-{name}.html",
            f"{SOURCE_BASE}/tree/{commit}/{directory}",
            directory,
            commit,
        )

    css = output / "sourcey-godoc.css"
    css.write_text(
        css.read_text()
        + "\n.source-meta{margin:.65rem 0 1.2rem;color:#5a6573;font-size:.9rem}"
        + ".source-meta a{color:inherit;text-decoration:underline}\n"
    )
    for name in ("llms.txt", "llms-full.txt"):
        path = output / name
        path.write_text(path.read_text().rstrip() + "\n")


if __name__ == "__main__":
    main()
