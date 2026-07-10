#!/bin/sh
# Prepend CHANGELOG.md with sections for every tag newer than the newest
# version already in the file, plus the upcoming tag passed as $1.
# Existing sections are never regenerated or modified.
set -e

TAG="$1"
[ -n "$TAG" ] || { echo "usage: $0 <next-tag>" >&2; exit 1; }

LAST=$(grep -m1 -o '<a name="v[0-9][^"]*"' CHANGELOG.md | cut -d'"' -f2)
[ -n "$LAST" ] || { echo "cannot find newest version in CHANGELOG.md" >&2; exit 1; }

FROM=$(git tag --sort=v:refname | awk -v last="$LAST" '$0 == last { getline; print; exit }')
RANGE="${FROM:-$TAG}.."

TMP=$(mktemp -d)
trap 'rm -rf "$TMP"' EXIT INT TERM

git-chglog --template .chglog/SECTION.tpl.md --next-tag "$TAG" "$RANGE" > "$TMP/generated.md"

grep '^\[v' "$TMP/generated.md" > "$TMP/links" || true
grep -v '^\[v' "$TMP/generated.md" > "$TMP/body"

awk -v body="$TMP/body" -v links="$TMP/links" -v tag="$TAG" '
    !inserted && /^<a name="v/ {
        while ((getline line < body) > 0) print line
        print ""
        inserted = 1
    }
    /^\[Unreleased\]:/ {
        sub(/compare\/.*\.\.\.HEAD/, "compare/" tag "...HEAD")
        print
        while ((getline line < links) > 0) print line
        next
    }
    { print }
' CHANGELOG.md > "$TMP/CHANGELOG.new"

mv "$TMP/CHANGELOG.new" CHANGELOG.md
