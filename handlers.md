# Handlers 

Handlers control the flow of requests coming in to the server, and typically should perform a sequence of actions something like:

1. Parse Request 
2. Authorise 
2. Act on Resources
3. Respond 

## Parsing Requests 

HTTP Requests consist of a URL (which may contain parameters) and (optionally) a body which contains more data usually as a form. 

```go
  // Parse URL params from request

```


```go
  // Parse form body params from request

```

Since you will probably be doing this a lot, you may want to pull it out into a function, or use a library that does this already. 


## Authorise 

There are a few ways of doing this (with Authorise headers, with encrypted cookies, with server-side sessions), but here we're going to consider encrypted cookies which are the most common approach taken by web frameworks in other languages like Rails or Django. 

A simple approach is to check for this cookie in middleware, and then query for it as necessary in the cookie of the request. (link to example, also show snippet here):


### Paths and Security 

Be aware that when you're choosing which routes you , you don't inadvertantly open your app to a directory traversal attack - this attack lets the attacker inject an unexpected path from the file system (for example by using /../ in urls), and. So a naive file handler might serve up unexpected files if given a path which triggered a directory traversal. You should never directly use request.URL.Path as a file path without first cleaning it with filepath.Clean(), and if using it for other you should consider restricting it to known good paths. There is a function in the http package to restrict a FileServer to just the files under a given path - http.Dir, this can be used if you want to serve files from a given directory only.

```go
http.FileServer(http.Dir("/my/path/"))
```

For this reason you should silo only public documents inside your web root at public (never keep private documents there), and never allow handlers to access files outside of this by paths which are tainted by information from the outside world. 


## Acting on Resources

After you have parsed the params and authorised your request, you need to perform some action, usually on a resource associated with this URL. So for example consider the url /users/1/update. You will expect to receive form data associated with a user, identified by the url. 


### Updating a record

```go 

// Update a record using form params


```

### Performing a request 

You may find your handlers need to call out to another service, either synchronously or asyncrhonously. In Go you can easily make a syncrhonous function call async simply be prefixing it with the go keyword. 


```go 

// Get data from  another service syncrhonously


```


```go 
// Post data to another service asynchronously with the go keyword 
go postData()

```


## Responding to Requests 

Responses to http verbs are typically limited to three options:

1. Respond with an error code (optionally rendering the error)
2. Render a view (for example html or json)
3. Redirect to another endpoint (for example after delete)

### Responding with an error code

To respond with an error code and some simple text, you can just use the built-in http function Error:

```go
  http.Error(w, "Page not found", http.StatusNotFound)
```

This may suffice, particularly if you're writing an API, but typically users of a larger web app will expect errors to be a little friendlier than a status code and text string.

### Rendering a view 

Instead of sending just an error code, you can also render an error page based on your app layout, so that readers see an informative error message, which lets them know how to proceed. Ideally every error in your web application should inform the user of exactly what went wrong. For example don't say 'Internal Error' or 'Date format error' but 'Please enter dates in the format '2021-01-01'. You may also want to display and/or log extra information like the underlying error string or even a stack trace if in development locally. 

Show code to render a template manually with struct, then compare and contrast with frag approach.

view.AddKey("user",user)
view.AddKey("page",page)


### Redirect to another endpoint 

After deleting a resource, it makes little sense to render a page simply saying deleted, so it is preferable to redirect to the list of resources, to show the updated list without the one just deleted. So for example if you visit:

``` 
/users/1/delete 
```

you might afterward expect to be sent to : /users, you can achieve this quite simply with net/http:

```
http.Redirect("/users",http.Found)
```









