#!/usr/bin/env python3
from __future__ import annotations

import hashlib
import json
import re
import sys
from html.parser import HTMLParser
from pathlib import Path
from urllib.parse import urlparse

COMMIT = "3f3cf9579987d520eef191d3f9b7d8c47d234276"
GENERATED_AT = "2026-07-12T05:50:03Z"
MODULE = "github.com/go-echarts/go-echarts/v2"
CANONICAL_BASE = "https://go-echarts.github.io/go-echarts/api/"


class Links(HTMLParser):
    def __init__(self) -> None:
        super().__init__()
        self.links: list[str] = []

    def handle_starttag(self, tag: str, attrs: list[tuple[str, str | None]]) -> None:
        values = dict(attrs)
        key = "href" if tag in {"a", "link"} else "src" if tag in {"img", "script"} else None
        if key and values.get(key):
            self.links.append(values[key] or "")


def main() -> None:
    root = Path(sys.argv[1]).resolve()
    snapshot = json.loads((root / "sourcey-godoc.json").read_text())
    assert snapshot["source"] == "sourcey-godoc"
    assert snapshot["module_path"] == MODULE
    assert snapshot["generated_at"] == GENERATED_AT
    assert len(snapshot["packages"]) == 9

    counts = {key: 0 for key in ("consts", "vars", "funcs", "types", "fields", "methods")}
    for package in snapshot["packages"]:
        for key in ("consts", "vars", "funcs", "types"):
            counts[key] += len(package.get(key) or [])
        for item in package.get("types") or []:
            counts["fields"] += len(item.get("fields") or [])
            counts["methods"] += len(item.get("methods") or [])
    assert counts == {"consts": 50, "vars": 9, "funcs": 125, "types": 200, "fields": 1177, "methods": 175}, counts

    html = sorted(root.glob("*.html"))
    assert len(html) == 10, len(html)
    broken: list[tuple[str, str]] = []
    source_pages = canonical_pages = generator_pages = 0
    for page in html:
        text = page.read_text()
        source_pages += int(f"https://github.com/go-echarts/go-echarts/tree/{COMMIT}" in text)
        canonical_pages += int(f'href="{CANONICAL_BASE}{page.name}"' in text)
        generator_pages += int('content="sourcey-godoc 0.2.1"' in text)
        parser = Links()
        parser.feed(text)
        for link in parser.links:
            parsed = urlparse(link)
            if parsed.scheme or link.startswith(("#", "mailto:", "javascript:")):
                continue
            target = (page.parent / parsed.path).resolve()
            if parsed.path.endswith("/"):
                target /= "index.html"
            if not target.exists():
                broken.append((page.name, link))
    assert source_pages == len(html)
    assert canonical_pages == len(html)
    assert generator_pages == len(html)
    assert not broken, broken[:20]

    required = ["llms.txt", "llms-full.txt", "sourcey-godoc.css", "LICENSE.sourcey.txt", "THIRD_PARTY_NOTICES.md", "BUILD.md"]
    assert all((root / name).is_file() for name in required)
    private = re.compile(r"/root/|/tmp/|localhost|127\.0\.0\.1|github_pat_|ghp_|BEGIN .*PRIVATE")
    assert not any(private.search(p.read_text(errors="ignore")) for p in root.rglob("*") if p.is_file())

    top_level = counts["consts"] + counts["vars"] + counts["funcs"] + counts["types"]
    total = top_level + counts["fields"] + counts["methods"]
    summary = {
        "status": "valid",
        "packages": len(snapshot["packages"]),
        "top_level_symbols": top_level,
        "total_documented_symbols": total,
        "html_pages": len(html),
        "source_pages": source_pages,
        "canonical_pages": canonical_pages,
        "broken_local_links": len(broken),
        "index_sha256": hashlib.sha256((root / "index.html").read_bytes()).hexdigest(),
        "snapshot_sha256": hashlib.sha256((root / "sourcey-godoc.json").read_bytes()).hexdigest(),
    }
    print(json.dumps(summary, sort_keys=True))


if __name__ == "__main__":
    main()
