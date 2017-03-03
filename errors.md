# Error Handling

People coming from other languages are often surprised by the simplistic nature of error handling in Go. While option types would may or may not be a nice addition, the error handling imposes simplicity on go programmers, and. 

* Don't use panic for normal errors
* Return early, use if err != nil return and then continue on the happy path
* Do handle errors, don't use _ to ignore them
* Do handle errors as close to the point of error as possible 
* Do store information about the problem (using custom types if necessary)
* In a web application, inform the user of exactly what went wrong - don't tell them 'Internal Server Error'

### Panic is not for errors

Panic is not for error handling in the normal flow or user errors but for unrecoverable programmer errors - every error in a web application should be gracefully recovered from and displayed to the user, so don't use panic in your web app or library. There are very exceptional situations where it is useful, but it should be obvious to you if you encounter them (for example when writing a parser if you encounter invalid markup).

### Return Early 

Keep your code mostly aligned to the left, with error conditions in if blocks, and return when you hit an error, so that the flow of control is clear. For example use:

```go
// Keep your normal flow to the left

// Do something, return on error
thing, err := doSomething()
if err != nil  { 
    return server.InternalError(err)
}

// Do something else 
err := doSomethingElse()
...
```

NOT this:

```go
// Don't do this, which becomes confusing quickly
thing, err := doSomething()
if err == nil  { 
    err := doSomethingElse()
    if err == nil {
      doThirdThing()
    } else {
      return server.InternalError(err)
    }
} else {
  return server.InternalError(err)
}
```

See [Line of sight](https://medium.com/@matryer/line-of-sight-in-code-186dd7cdea88) by Mat Ryer for more explanation. 


### Errors are Values

Errors can be values, so you can use custom error types. While they should not be overused, if you find yourself checking strings for errors, or wishing you had other information in them (like an error code), you need a custom error type. You can then assert against this type in your error handler and obtain more information. 

You can also collect errors by wrapping operations in a type which no-ops after it encounters an error, and stores the first error encountered, or stores a list of errors encountered. In some situations this may be neater than a large number of if err != nil blocks. e.g.:

```go
ew := &errWriter{w: fd}
ew.write(p0[a:b])
ew.write(p1[c:d])
ew.write(p2[e:f])
// and so on
if ew.err != nil {
    return ew.err
}
```

For more on this approach see [Errors are values](https://blog.golang.org/errors-are-values) by Rob Pike.

### Handling Errors 

In libraries, you should simply create or annotate and error and pass it back to the caller. Avoid logging in libraries if you can, as this avoids imposing your choice of logging on users of your library pkg. You can insert logging for debugging and take it out again when done. 

In handlers, you should attempt to inform the user and server log, and then if it is serious exit the handler, or if it is not serious log and continue. In many cases in a web server errors may be transient or due to bad input, so you should attempt to inform the user exactly why the error occurred (without giving away important detail as to the workings of your server). 

Every error message in a web application should have a clear message for the user telling them how they can remedy the situation (usually by changing their input or visiting another page). 

### Errors in Handlers

When requests come in to a handler, they often fail for similar reasons, which can be nicely mapped to http codes, like NotAuthorized, or NotFound etc. 

For a web server, you'll want to handle these by showing the user an error message which tells them what went wrong, hopefully rendered in a template and in plain language, not obscure error codes, and usually also logging the error (either locally or in a monitoring system).

This is where the explicit handling in Go really shines - it forces you to consider for each error condition what the user should see. 

Unfortunately Go doesn't have an elegant way to handle errors in handlers built-in. There are two options for dealing with this:

1. Handle the error by rendering it somehow, then return

```
if err != nil {
  log.Printf("myhandler: %s",err)
  server.RenderError("User Not Found", "Sorry, this user couldn't be found")
  return
}
```

2. Change the handler signature to return an error (probably a custom error type), and handle it with an error handler which you define. 

```
if err != nil {
  return server.NotFound(err,"User Not Found", "Sorry, this user couldn't be found")
}
```

They are not very different, but I find it slightly more elegant to have a central error handler, which can decide whether to log or not, and can use a common template to render the error gracefully, and on one line rather than two. As handlers typically contain several possible error paths each, the extra return line does add up. The advantage of option 1 is it sticks to the standard http.HandlerFunc signature.  


## References 

* [Error Handling and Go](https://blog.golang.org/error-handling-and-go)
* [Error handling vs Exceptions](https://dave.cheney.net/2014/11/04/error-handling-vs-exceptions-redux)
* [Line of sight](https://medium.com/@matryer/line-of-sight-in-code-186dd7cdea88) explains why you should try to avoid else.
* [Errors are values](https://blog.golang.org/errors-are-values) 