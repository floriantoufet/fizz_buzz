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

### Call endpoints

#### Swaggers

All endpoints are detailed in [swagger.yml](https://github.com/floriantoufet/fizzbuzz/blob/main/swagger.yml)

#### Godoc

Available [here](https://pkg.go.dev/github.com/floriantoufet/fizzbuzz)

#### Conditions

- fizz string and buzz string do not exceed 100 characters
- modulus and limit are strictly positives
- limit does not exceed 1000
