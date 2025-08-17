#!/bin/sh
#
# Script to run tests in an isolated Docker environment.
# Works the same on macOS, Linux, and Windows (via Git Bash/WSL).
#

echo "--- Running tests in Docker container ---"

# Detect operating system for correct path mapping
if [ -n "$WINDIR" ]; then
  # On Windows (in Git Bash/WSL), get the path using cmd.exe
  CURRENT_DIR=$(pwsh -NoProfile -Command '(Get-Location).Path')
else
  # On macOS or Linux, use the standard `pwd`
  CURRENT_DIR=$(pwd)
fi

echo "--- Using directory: $CURRENT_DIR ---"

# Run the container, map the correctly resolved directory, and then execute the tests
# The container will be automatically removed upon completion thanks to the --rm flag
docker run --rm -v "$CURRENT_DIR:/app" -w //app golang:1.23-alpine sh -c "apk add --update git make && make ssh-server && make test"
# docker run --rm -v "$(pwd):/app" -w "/app" golang:1.23-alpine sh -c "apk add --update git make && make ssh-server && make test"

echo "--- Tests finished ---"
