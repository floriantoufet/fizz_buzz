# Fizz Buzz ![](https://github.com/floriantoufet/fizzbuzz/workflows/Tests/badge.svg) ![](https://github.com/floriantoufet/fizzbuzz/workflows/Releases/badge.svg) [![Go Reference](https://pkg.go.dev/badge/github.com/floriantoufet/fizzbuzz.svg)](https://pkg.go.dev/github.com/floriantoufet/fizzbuzz)

Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are
replaced by str2, all multiples of int1 and int2 are replaced by str1str2. :

### Download binaries

Go to [latest release](https://github.com/floriantoufet/fizzbuzz/releases/latest)

### Run

Linux:

 ```shell script
 $ chmod a+x fizz_buzz_linux-amd64
 $ ./fizz_buzz_linux-amd64
 ```

MacOS:

 ```shell script
 $ chmod a+x fizz_buzz_darwin-amd64
 $ ./fizz_buzz_darwin-amd64
 ```

Windows:

 ```shell script
 C:\....\fizz_buzz_windows-amd64.exe
 ```

#### Config

Create `env.yml` in `conf` directory

```yaml
transports:
  http:
    host: 127.0.0.1
    port: 8080

logger:
  env: "dev" # Supports dev or prod value

```

#### Launch with environment variable

 ```shell script
 $ CONF_ENV=env ./fizz_buzz_darwin-amd64
 ```

`CONF_ENV` take value of wanted conf file

### Call endpoints

#### Swaggers

All endpoints are detailed in [swagger.yml](https://app.swaggerhub.com/apis/floriantoufet/FizzBuzz/1.0.0#/StatisticResponse)

#### Godoc

Available [here](https://pkg.go.dev/github.com/floriantoufet/fizzbuzz)

#### Conditions

- fizz string and buzz string do not exceed 100 characters
- modulus and limit are strictly positives
- limit does not exceed 1000

### Tests

#### Requirements

- install [golangci](https://github.com/golangci/golangci-lint)
- install [godog](https://github.com/cucumber/godog)

#### Run tests

 ```shell script
 $ ./run_test.sh
 ```


