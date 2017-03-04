# Testing in Go

You may be tempted to look for a testing framework, but you probably don't need one. Try it with the std lib first, and if you find yourself frustrated in ways that can't be solved with a few helper functions, *then* look at the options available. 

## Table driven tests 

Look at std lib. 


## Unit Tests 

Many of your functions and methods should be testable without any external state. They should operate solely on the parameters and, sometimes, on the data stored in the object they are attached to. This means you can easily test them in isolation. Any functions you can't test in isolation, like handlers, should be tested with integration tests (see below).

An example of a unit test (useful for twitter app) would be testing that setting the role on a user disallows certain actions. 

## Integration Tests

Integration tests are those which test the entire flow of data from one end of the app to the other. If you find yourself wondering about injecting dependencies, you need an integration test. Integration tests will typically require mocking up your app data, ideally using as much fo the real infrastructure as possible. The more that you mock, the more fragile your tests are in the face of changing requirements and dependencies. Ideally, your app will have very few mocks, only a test database which it uses to run tests against - this lets you perform integration tests against an environment close to production and means less work writing fake data sources which have to be kept up to date as dependencies change. 

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

The second step might involve the most work - adding logging, adjusting data to eliminate edge cases, checking assumptions carefully about the cause. After that it should be simple to fix the problem as you already understand it. 

When you are testing, try to think of cases where your functions could go wrong, faulty assumptions and edge cases that you might not have thought of. 
