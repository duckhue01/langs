## testing, benchmarking and fuzzing
testing your code is not something that you should wait to do until after you're finished developing your program. With's go testing framework, unit testing and benchmarking can happen during the development process. Just like the `go build` command, there's a `go test` command to execute explicit test code that you write.

### Unit testing
- there are several ways in Go to write unit test: basic test, table test
1. basic unit test
- basic unit test use for almost of case but it hard to extend later.
- t.Log will expose the message when we run test with -v flag and test case fail
- t.Fatal will expose message then stop execution
- t.Error
- try to log all needed information 

### Table testing
- when you're testing code that can accept a set of different parameters with different result, a table test should be used.

### mocking calls
- the unit tests we wrote are great, but they do have a couple of flaws. First, they require access to the internet in order for the tests to run successfully.
- in case you writing upper layer of application and the bellow layer work perfectly then you should use mock the in order to decoupling with bellow layer

### testing end-point