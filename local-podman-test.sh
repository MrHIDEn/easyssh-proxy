#!/bin/sh
#
# Script to run tests in an isolated Podman environment.
# Works the same on macOS, Linux, and Windows (via Git Bash/WSL).
#

echo "--- Running tests in Podman container ---"

# Run the container, map the current directory, and then execute the tests
# The container will be automatically removed upon completion thanks to the --rm flag
podman run --rm -v "$(pwd):/app" -w "/app" golang:1.23-alpine sh -c "apk add --update git make && make ssh-server && make test"

echo "--- Tests finished ---"
