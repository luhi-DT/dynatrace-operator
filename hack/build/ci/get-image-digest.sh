#!/bin/bash

digest=$(skopeo inspect docker://"${IMAGE}" --format "{{.Digest}}")
digest_value=$(echo ${digest} | tr : -)
echo "digest=${digest}">> "$GITHUB_OUTPUT"
echo "digest_value=${digest_value}">> "$GITHUB_OUTPUT"

