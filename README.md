# go-testing

https://app.pluralsight.com/library/courses/go-testing-applications/table-of-contents

**Command to execute the tests:**
```
go test go-testing/tests
```

**Command to execute the tests with right code coverage:**
```
go test -coverpkg=./... go-testing/tests
```
Further information [here!](https://lk4d4.darth.io/posts/multicover/)

```
go test go-testing/tests -cover
```
will give you 0% of code coverage because math.go is not in the same package as math_test.go
