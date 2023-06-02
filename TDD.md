# Test Driven Development

Writing test first and then writing the code to fix failing test

1. Write the test case
2. Make the compiler pass
3. Run the test, see for the failing test cases
3. Write code to fix the failing test cases
5. Refactor


### Writing Tests in Go Lang
1. Create a file name suffixed with `xxxx_test.go`
2. The test function should start with `Test`
3. The test function should take one argument only `t *testing.T`
4. import the `testing` package, the golang testing framework


#### Subtests
Sometimes it is useful to group tests around a "thing" and then have subtests describing
different scenarios. A benefit of this approach is you can set up shared code that can
be used in the other tests.


## Appendix
#### Installing go doc locally
Run `go install golang.org/x/tools/cmd/godoc@latest`

You can launch it by running `godoc -http :8000`

#### Running test in golang
Run `go lang`