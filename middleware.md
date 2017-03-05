# Go Middleware

Middleware is an example of the (Pipeline Design Pattern)[http://www.cise.ufl.edu/research/ParallelPatterns/PatternLanguage/AlgorithmStructure/Pipeline.htm] - requests are fed through a chain of handlers until they eventually reach execution. 

Beware though, unless your app is fundamentally suited to this pattern, it might be best to restrict middleware to 

Here is a demonstration of from the article by Mat Ryer linked below:

```go
func log(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    
    log.Printf("Request recevied")
    h.ServeHTTP(w, r) 
    log.Println("After")
  
  })
}
```

## Router support 

It's not necessary for your router to support this technique, though it may make things a little easier. If it does, you can call something similar to this to register your middleware:

```go 
  
```

If your router doesn't support middleware, you can still just wrap your handler functions before passing them to the router like this:

```go 
  
```

### Authentication 


```go 
  
```


### Logging 

Use fragmenta server/log as an example here of middleware to log and add a request identifier. 

```go 
  
```

### Overuse

There is a performance overhead to using middleware on every request, and it should not be overused for that reason, but there is also a complexity overhead. It adds implicit hidden behaviour to your application, which is difficult to reason about as it is not easy to trace execution through lots of different handlers. If you can do something explicitly in your handler, and it's not something you want to do on every handler, it's probably better to handle it explicitly rather than building middleware which is sometimes applied, or has conditional logic in it. 


### References 

* [The Pipeline Pattern](http://www.cise.ufl.edu/research/ParallelPatterns/PatternLanguage/AlgorithmStructure/Pipeline.htm)
* [The http.Handler wrapper technique by Mat Ryer](https://medium.com/@matryer/the-http-handler-wrapper-technique-in-golang-updated-bc7fbcffa702)
* [Writing HTTP Middleware in Go](https://justinas.org/writing-http-middleware-in-go)