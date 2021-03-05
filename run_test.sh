echo -e "Start tests"

echo -e "Start lint"
golangci-lint run

echo -e "Start units tests"
go test ./...

echo -e "Start API tests"
echo -e "Waiting for app restart"
# Build app
go build -o /tmp/fizz_buzz bin/main.go
/tmp/fizz_buzz &

# Get pid for ensure killing app after tests
PID=$!

# Wait for app response
while ! nc -z localhost "8080"; do
  sleep 1
done

# Launch tests
cd tests
godog
ERR_CODE=$?

# Kill app
kill -9 $PID 2> /dev/null

# Remove binary
rm -f /tmp/fizz_buzz 2> /dev/null
exit $ERR_CODE
