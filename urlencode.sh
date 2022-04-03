#!/bin/bash
echo $1 | python3 -c "import urllib.parse, sys; print(urllib.parse.quote(sys.stdin.read()))"
