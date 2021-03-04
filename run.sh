echo -e "Waiting for app restart"
go run bin/main.go &
while ! nc -z localhost "8080"; do
  sleep 1
done

cd tests
godog
