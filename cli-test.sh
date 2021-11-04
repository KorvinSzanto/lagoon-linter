#!/usr/bin/env sh
set -eu

LINTER=./dist/lagoon-linter_linux_amd64/lagoon-linter
chmod +x "$LINTER"

# profile: required
for lagoonyml in ./internal/lagoonyml/required/testdata/valid.*.yml; do
	echo "$lagoonyml"
	ln -fs "$lagoonyml" .lagoon.yml
	$LINTER
done
for lagoonyml in ./internal/lagoonyml/required/testdata/invalid.*.yml; do
	echo "$lagoonyml"
	$LINTER validate --lagoon-yaml="$lagoonyml" && exit 1
done

# profile: deprecated
for lagoonyml in ./internal/lagoonyml/deprecated/testdata/valid.*.yml; do
	echo "$lagoonyml"
	$LINTER validate --profile=deprecated --lagoon-yaml="$lagoonyml"
done
for lagoonyml in ./internal/lagoonyml/deprecated/testdata/invalid.*.yml; do
	echo "$lagoonyml"
	$LINTER validate --profile=deprecated --lagoon-yaml="$lagoonyml" && exit 1
done

exit 0
