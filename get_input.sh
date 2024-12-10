#!/bin/bash

# Check if the user provided a day number
if [ -z "$1" ]; then
    echo "Usage: $0 <day_number>"
    exit 1
fi

DAY=$1
OUTPUT_FILE="./input"

# Define the cookie string
COOKIE="_ga=GA1.2.1646005350.1733067961; session=53616c7465645f5fe4eae111ce7e230b9521f5323d26f7805c15a6f841162f82139d022c370bf632f2e609f531201d200e1a7b21033ae08719ca00513856beba; _gid=GA1.2.938934616.1733244453; _ga_MHSNPJKWC7=GS1.2.1733798213.21.1.1733798214.0.0.0"

# Execute the curl command and save the output
curl "https://adventofcode.com/2024/day/${DAY}/input" \
  -H "cookie: $COOKIE" \
  -o "$OUTPUT_FILE"

# Check if the curl command succeeded
if [ $? -eq 0 ]; then
    echo "Input for day $DAY saved to $OUTPUT_FILE"
else
    echo "Failed to fetch input for day $DAY"
fi

