#!/usr/bin/env bash

set -eu

# Change directory to where this script lives
cd "$(dirname "$0")"

STREAM_URL="http://stream.radiox.de:8000/live"
DURATION="03:00:00"
ISODATE="$(date -I)"
FILENAME="radio-x-fade-${ISODATE}.mp3"
TITLE="$(./radio-x-fade-title)"

echo "Recording $TITLE for $FILENAME"
ffmpeg -loglevel warning \
  -reconnect 1 -reconnect_at_eof 1 -reconnect_streamed 1 -reconnect_delay_max 2 \
  -i "$STREAM_URL" \
  -t "$DURATION" \
  -c copy \
  -metadata artist="DJ Mixes" \
  -metadata album="Radio X" \
  -metadata title="$TITLE" \
  -metadata date="$ISODATE" \
  "$FILENAME" 2> record_errors.log