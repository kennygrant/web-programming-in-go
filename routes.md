# Routing in Go

Since Go has a very simplistic router in the standard library, many people have built their own. At the end of this chapter there are links to a few more popular routers, with a short overview of each. 

### A note on performance

While it is tempting to measure something like a router purely on performance, unless it is pathologically slow (e.g. uses regexp in a naive way or allocates a lot on every new request) it is not likely to take up many resources compared to your handlers which have to talk to the database and write responses. So measures of performance on routers *are* useful indicators but should not be your primary concern when choosing one. As the end user of a router factors which will usually be more important are the way it parses parameters, control over evaluation order of routes, support for middleware and the handler signatures. 

### Let's build a ServeMux

Because Go is open source, you can go and have a look at the [DefaultServeMux](https://golang.org/src/net/http/server.go?#L1865) yourself in the standard library. This is a fairly simple router which has a few drawbacks -

* It doesn't let you collect parameters at all, name them or limit their content
* It doesn't guarantee the evaluation order of routes
* It doesn't let you define groups of routes

However for many uses, it's completely acceptable. If you want to have some control over order of evaluation, or named parameters in your routes, you may want to find a router which has slightly more sophisticated route parsing, or write your own.

Let's start with a router, which has a list of routes, which you can add to with .Add(). The code below is the complete code for a simple router which behaves very much as the built-in servemux does, but preserves evaluation order. 

```go 
// Route is a simplistic representation of a route
type Route struct {
  // A path prefix to match 
	Path    string
  // A handler to execute if this route matches 
	Handler http.HandlerFunc
}

// Router stores a list of routes, protected by a mutex
type Router struct {
	mu     sync.RWMutex
	routes []*Route
}

// Add a route to our list of routes 
func (r *Router) Add(p string, fn http.HandlerFunc) {
	r.mu.Lock()
	r.routes = append(r.routes, &Route{Path: p, Handler: fn})
	r.mu.Unlock()
}

// ServeHTTP serves all HTTP requests
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

  // On receiving a request, walk the list of routes in order of addition
  // looking for a match. 
	r.mu.RLock() // Use a read lock to lock writes
	for _, route := range r.routes {
		// Return a match if our route prefix matches path
		if strings.HasPrefix(req.URL.Path, route.Path) {
			route.Handler(w, req)
			break
		}
	}
	r.mu.RUnlock() // Use a read unlock
  
}
```

The mutex is not strictly necessary if you never update the routes after startup on the main goroutine, but is good practice, just in case someone dynamically adds routes during program execution. 

### Context

In Go 1.7 a new package was introduced to the Standard Library which is intended to allow passing request-scoped values (for example deadlines, cancellations, user ids, request ids) across goroutines and backends so that the caller can cancel a request easily for example, and all resources being used for it will be cleaned up.

In the context of a simpler web app this can be useful for passing *request-specific* values between middleware and handlers, and between handlers and goroutines they spawn to perform tasks like sending mail. You should not abuse it to send lots of data though - better to pass in dependencies. 

It should not be used for passing dependencies like loggers or database connections, or convenient globals like a pointer to your app instance which contains these things. There are other better ways to access dependencies rather than attaching them to a request object which really has nothing to do with them. 


### Handler definitions 

Many routers take the approach of passing in their own context which provides handy dependencies and helper functions for handlers. For example the [Echo context](https://godoc.org/github.com/labstack/echo#Context). While superficially attractive this approach can lead to a few problems:

* Dependency injection - this might seem an advantage, but it means your router will be passing in things like loggers or database connections it has no business controlling. 
* Scope creep - your router starts to know all about your handlers, and pass more and more information in to them by parsing parameters in advance, storing bags of values from middleware, munging requests or providing a wrapper around the request object. 
* Lockin - your handlers become tied to a given signature, not just in the function call, but in all the functions of the router they call. 

Another approach with similar downsides is to define your handlers on an object which contains all your dependencies, or wrap them in a function which defines dependencies. Both approaches lead to handlers being tied very tightly to a given app when really they should know only about the information they require to respond to a given request.  

For this reason I recommend sticking with the standard handler definition, which will be instantly recognisable to people visiting your code, and keeps your handlers in charge of responding to requests. One possible addition to it is to return an error, to allow a shared error handler to be elegantly used. 

## Popular routers 

While benchmarks are not the primary criteria for choosing code, it is important that your router doesn't introduce too much latency in your application, as it will delay every request, you can find some simple benchmarks of these routers here: https://github.com/kennygrant/routebench


#### Gorilla Mux
Handler signature: ```go func (w http.ResponseWriter, req *http.Request) ```
Route signature: /hello/{param:\d+}
Advantages: Flexible regexp params, stdlib signature
Disadvantages: Slightly slower, doesn't support middleware
This router is similar to the stdlib mux but uses regexp to evaluate params. It is probably the most commonly used router and was released early on. 

#### Fragmenta Mux 
Handler signature: ```go func (w http.ResponseWriter, req *http.Request) error ```
Route signature: /hello/{param:\d+}
Advantages: Ordered evaluation, regexp params, deferred parsing of params, Supports middleware
Disadvantages: Requires handlers to return error
This router takes a similar approach to the Gorilla mux in using regexp to define paramaters, so route definitions are the same. It supports middleware chains, and deferred parsing of params so that they are not parsed until required. 

#### HttpRouter 
Handler signature: ```go func (w http.ResponseWriter, req *http.Request) ```
Route signature: /hello/:param
Advantages: Fast, stdlib signature
Disadvantages: Can't specify param type, Can't handle all routes
This router is focussed on speed, and uses a data structure which allows this. It may be suitable for larger sites which don't require control over order of evaluation and have a lot of routes. 

#### Echo 
Handler signature: ```go func (c echo.Context) error ```
Route signature: /hello/:param
Advantages: Supports middleware, Grouped APIs
Disadvantages: Custom handlers, expanded scope
This has ambitions to be more than a router, expanding into serving, logging, view rendering, so it is perhaps better thought of as a framework. 

#### Bufallo
Handler signature: ```go func(c buffalo.Context) error ```
Limitations: Ordered evaluation of routes
This router takes a similar approach to the gorilla one, but has does do significantly more work for every evaluation of a route. It is relatively new and will probably improve with time.  

There is a set of benchmarks of Golang routers here, to give a rough idea of relative performance with large sets of routes. If your app has less than 50 routes this is unlikely to matter, and even then other factors are probably more important in most cases.   

