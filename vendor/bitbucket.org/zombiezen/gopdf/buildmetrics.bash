#!/bin/bash

DIR="$(mktemp -d -t buildmetricsXXX)"
TARBALL="Core14_AFMs.tar"
URL="https://partners.adobe.com/public/developer/en/pdf/$TARBALL"
OUTPUT="pdf/metrics.go"

echo "$DIR"
curl "$URL" > "$DIR/$TARBALL"
(cd "$DIR" ; tar -xf "$TARBALL")

rm -f "$OUTPUT"
touch "$OUTPUT"
echo "package pdf" >> "$OUTPUT"

for afmFile in "$DIR"/*.afm
do
    fontName="$(basename "$afmFile" .afm)"
    varName="$(echo "$fontName" | sed 's/-//')Widths"
    echo >> "$OUTPUT"
    echo "// $fontName" >> "$OUTPUT"
    echo "var $varName = []uint16{" >> "$OUTPUT"
    sed -n 's/^C \([0-9]\|[0-9][0-9]\|1[012][0-9]\) ; WX \([0-9]\+\) ; N \([a-zA-Z]\+\).*/\1: \2, \/\/ \3/p' < "$afmFile" >> "$OUTPUT"
    echo "}" >> "$OUTPUT"
done

gofmt -w "$OUTPUT"
rm -rf "$DIR"
