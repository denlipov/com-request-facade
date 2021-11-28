#!/bin/bash

~/go/bin/golangci-lint run --issues-exit-code 0 \
        --out-format code-climate | \
        tee gl-code-quality-report.json | \
        jq -r '.[] | "\(.location.path):\(.location.lines.begin) \(.description)"'
