# notes
go test coverage
```

go test -cover

# full version
go test -v -coverprofile cover.out .
go tool cover -html cover.out -o cover.html
open cover.html
```
host into a html
red is covered, white not covered

## Discipline
Write a test

Make the compiler pass

Run the test, see that it fails and check the error message is meaningful

Write enough code to make the test pass

Refactor

## Questions

- how to godoc my own package? run from package dir? or add it to goroot/gopath?