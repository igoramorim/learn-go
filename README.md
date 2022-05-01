# learn-go

## What is it?

Some Go code used to learn the language. Code from personal 'projects', Go's official documentation and tutorials.

## Notes

A list of tips, good practices and knowledge for easy access.

1. `go mod init example.com/greetings` creates a module named example.com/greetings

2. Test files must use **_test.go** as suffix file's name

3. Test functions must start with '**Test**' followed by something that describes the test. Otherwise the test doesn't run

4. `go build` compiles the packages and generate an executable

5. `go install` compiles the packages and install the results. That means the executable will be saved in a specific folder. This folder can be found by running `go list -f '{{.Target}}'`. This folder probably is in the **PATH**, so after the install you should run your app from any directory

6. `go test` runs the **_test.go** files. Use `go test -v` for details

7. `go mod edit -replace example.com/greetings=../greetings` tells Go to look for the greetings package locally. It will modify the go.mod file

8. `go mod tidy` tells Go to to update the the module dependecies. The command updates the go.mod file

9. `go run .` runs the module

10. TODO: `go:generate`