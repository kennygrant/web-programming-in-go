# Error Handling

People coming from other languages are often surprised by the simplistic nature of error handling in Go. 



### Errors in Handlers

When requests come in to a handler, they often. For a web server, you'll want to handle these by showing the user an error message which tells them what went wrong, hopefully rendered in a template and in plain language, not obscure error codes, and usually also logging the error (either locally or in a monitoring system).

This is where the explicit handling in Go really shines - it forces you to consider for each error condition what the user should see. 

Unfortunately Go doesn't have an elegant way to handle errors in handlers built-in. There are two options for dealing with this:

1. Handle the error, then return

2. Change the handler signature to return an error (probably a custom error type), and handle it with an error handler which you define. 