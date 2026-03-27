#!/usr/bin/env bash

set -e

GO_MOD_PACKAGE="github.com/thesixnetwork/six-protocol"

echo "Generating gogo proto code"
cd proto
proto_dirs=$(find . -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep -q "option go_package" "$file" && grep -H -o -c "option go_package.*$GO_MOD_PACKAGE/api" "$file" | grep -q ':0$'; then
      buf generate --template buf.gen.gogo-nonignite.yaml $file
    fi
  done
done

echo "Generating pulsar proto code"
buf generate --template buf.gen.pulsar-nonignite.yaml

cd ..

cp -r $GO_MOD_PACKAGE/* ./
rm -rf github.com

go mod tidy