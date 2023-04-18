#!/bin/bash

REQUIRED=$1
COVERAGE=$(go test -cover ./... | grep 'coverage: ' | sed -e 's/.*coverage:\(.*\)% of.*/\1/')

if [[ "${COVERAGE%.*}" -lt "$REQUIRED" ]]; then
    echo "Coverage is $COVERAGE%, but $REQUIRED% is required"
    exit 1
fi