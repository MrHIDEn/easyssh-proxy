#!/bin/sh
#
# Script to run tests in an isolated Docker environment.
# Works the same on macOS, Linux, and Windows (via Git Bash/WSL).
#

echo "--- Running tests in Docker container ---"

# Run the container, map the current directory, and then execute the tests
# The container will be automatically removed upon completion thanks to the --rm flag
docker run --rm -v "$(pwd):/app" -w "/app" golang:1.23-alpine sh -c "apk add --update git make && make ssh-server && make test"

echo "--- Tests finished ---"
