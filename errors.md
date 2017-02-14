# Error Handling

People coming from other languages are often surprised by the simplistic nature of error handling in Go. While option types would may or may not be a nice addition, the error handling imposes simplicity on go programmers, and. 

* Don't use panic for normal errors
* Do handle errors, don't use _ to ignore them
* Do handle errors as close to the point of error as possible 
* Do store information about the problem (using custom types if necessary)
* Do use if error return and then continue, never if happy else error - see [Line of sight](https://medium.com/@matryer/line-of-sight-in-code-186dd7cdea88)

### Panic is not for errors

Panic is not for error handling in the normal flow but for unrecoverable errors - every error in a web application should be gracefully recovered from and displayed to the user, so don't use panic. 

### Errors are Values

Errors can be values, so you can use custom error types. While they should not be overused, if you find yourself checking strings for errors, or wishing you had other information in them (like an error code), you need a custom error type. You can then assert against this type in your error handler and obtain more information. 


### Handling Errors 

In libraries, you should simply create or annotate and error and pass it back to the caller. Avoid logging in libraries if you can, as this avoids imposing your choice of logging on users of your library pkg. You can insert logging for debugging and take it out again when done. 

In handlers, you should attempt to inform the user and server log, and then if it is serious exit the handler, or if it is not serious log and continue. 

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
* [Error handling and Go](https://blog.golang.org/error-handling-and-go) Error handling