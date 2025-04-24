#!/bin/bash
cd /ki || exit

export CGO_ENABLED=0
export GODOG_TAGS

if [ -n "$RUN_ONCE" ];
then
  echo "run once all tests..."
  go test -timeout 30m ./tests -count 1
  exit $errorCode
fi

count_tags() {
   find ./tests/features -name "*.feature" -exec grep -l '@dev' {} \; | wc -l
}

run_tests() {
    echo "Starting tests..."
    local tags=$(count_tags)
    if [ "$tags" -gt 0 ]; then
        echo "Running tests with @dev tag..."
        GODOG_TAGS="@dev"
    else
        echo "@dev tag not found, running all tests..."
        unset GODOG_TAGS
    fi
    sleep 1
    go test ./tests -v -count 1
}

run_tests

while inotifywait -qq -r -e moved_to -e create -e modify -e delete -e close_write /ki > /dev/null 2>&1; do
    run_tests
done