# Testing in Go

You may be tempted to look for a testing framework, but you probably don't need one. Try it with the std lib first, and if you find yourself frustrated in ways that can't be solved with a few helper functions, *then* look at the options available. 

## Table driven tests 

Table driven tests take a table of test values -> result and walk through them testing each one against the expected result. This can be a really powerful way of checking lots of possibilites without a lot of boilerplate. Here is an example from the fmt package:

```go
var flagtests = []struct {
    in  string
    out string
}{
    {"%a", "[%a]"},
    {"%-a", "[%-a]"},
    {"%+a", "[%+a]"},
    {"%#a", "[%#a]"},
    {"% a", "[% a]"},
    {"%0a", "[%0a]"},
    {"%1.2a", "[%1.2a]"},
    {"%-1.2a", "[%-1.2a]"},
    {"%+1.2a", "[%+1.2a]"},
    {"%-+1.2a", "[%+-1.2a]"},
    {"%-+1.2abc", "[%+-1.2a]bc"},
    {"%-1.2abc", "[%-1.2a]bc"},
}

func TestFlagParser(t *testing.T) {
    var flagprinter flagPrinter
    for _, tt := range flagtests {
        s := Sprintf(tt.in, &flagprinter)
        if s != tt.out {
            t.Errorf("Sprintf(%q, &flagprinter) => %q, want %q", tt.in, s, tt.out)
        }
    }
}
```

## Benchmarking 

You can also use go to benchmark your code and measure how it runs, simply by naming your test with a Benchmark prefix.

```go 

// Name functions with a Benchmark prefix
func BenchmarkBigLen(b *testing.B) {
    // Perform expensive setup 
    big := NewBig()
    
    // reset the timer before benchmarking 
    b.ResetTimer()
    
    // Allow the go tool to benchmark by wrapping the actual work in this for loop
    for i := 0; i < b.N; i++ {
        big.Len()
    }
}
```

## Unit Tests 

Many of your functions and methods should be testable without any external state. They should operate solely on the parameters and, sometimes, on the data stored in the object they are attached to. This means you can easily test them in isolation. Any functions you can't test in isolation, like handlers, should be tested with integration tests (see below). Unit tests do not use the database or other external state, which makes them extremely easy to test in isolation and reason about.  

An example of a unit test (useful for twitter app) would be testing that setting the role on a user disallows certain actions, or that asking whether a booking falls within certain dates returns the right answer. Below is an example of how to test a function which does not rely on external state. 

```go
// TestDatesBetween tests if bookings 
func TestDatesBetween(t *testing.T) {

  booking := &Holiday{
    CreatedAt: time.Now(),
      CreatedAt: time.Now(),
  }

  if holiday.Within(start,end) {
    t.Errorf("booking between returned true for dates:",start,end)
  }
  
}
```

## Integration Tests

Integration tests are those which test the entire flow of data from one end of the app to the other. If you find yourself wondering about injecting dependencies in order to test, you need an integration test. Integration tests will typically require mocking up your app data, ideally using as much of the real infrastructure as possible. The more that you mock, the more fragile your tests are in the face of changing requirements and dependencies. Ideally, your app will have very few mocks, only a test database which it uses to run tests against - this lets you perform integration tests against an environment close to production and means less work writing fake data sources which have to be kept up to date as dependencies change. 

As an example of an integration test we're going to test that Anon users cannot in fact edit someone else's user record, or delete users. You'll need to import the net/httptest package in order to create mock writers and readers. 

```go

// TestDeleteUser tests that only admin users can delete a user
func TestDeleteUser(t *testing.T) {
  
  w := 
  r := 
  
  // mock admin user on request 
  err := HandleDeleteUser(w,r)
  if err != nil {
    
  }
  
  // Now try with an anon user (no permissions)
  err := HandleDeleteUser(w,r)
  if err != nil {
    
  }
  
  
}
```

Do not just test the happy path (though you should test that as well), make sure your tests attempt edge cases, and feed your functions

## Debugging 

"The most important debugging tool is a reproducible test case."

https://twitter.com/jgrahamc/status/830196404765257728

Something has gone wrong. Your first instinct is to dive in and try to find the problem in the code, but instead you should focus on the following steps:

1. Reproduce the error - you must be able to reproduce it to know it is fixed 
2. Understand the error - trace execution until you are completely confident you understand the problem.
3. Fix the error - only when you can reproduce and understand it should you

The second step might involve the most work - adding logging, adjusting data to eliminate edge cases, checking assumptions carefully about the cause. After that it should be simple to fix the problem as you already understand it. When you are testing, try to think of cases where your functions could go wrong, faulty assumptions and edge cases that you might not have thought of. 

If you are keen to try a debugger, the Delve debugger is in active development targetting the go platform and is written in go. At the time of writing it is still experimental.

## References 

* [The testing package](https://golang.org/pkg/testing/) - read the standard library testing package first
* [Examples coverage](https://rakyll.org/examples-coverage/)
* [Parallelize your table-driven tests](https://rakyll.org/parallelize-test-tables/)
* [Writing Table-Driven tests in Go](https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go)
* [Self-documenting tests](https://rakyll.org/naming-tests-to-doc/)
* [Delve](https://github.com/derekparker/delve) is a debugger for the Go programming language.

