#!/usr/bin/env bash

if [[ -z "$1" ]] || [[ -z "$2" ]]; then
  echo "Usage: $0 <mp3> <time>"
  echo
  echo "Remove the given amount of time (HH:MM:SS format) from the"
  echo "beginning of an MP3 and save it with a '_trimmed' suffix"
  exit 1
fi

set -eu

INFILE="$1"

# Add `_trimmed` before extension in resulting filename
BASENAME="${INFILE%.*}"
EXT="${INFILE##*.}"
TRIM_AMOUNT="$2"
OUTFILE="${BASENAME}_trimmed_${TRIM_AMOUNT//:/-}.${EXT}"

ffmpeg -loglevel error \
  -ss "$TRIM_AMOUNT" \
  -i "$INFILE" \
  -vn \
  -acodec copy \
  "$OUTFILE"