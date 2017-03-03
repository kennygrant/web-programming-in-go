# Handlers 

Handlers control the flow of requests coming in to the server, and typically should perform a sequence of actions something like:

1. Parse Request 
2. Authorise 
2. Act on Resources
3. Respond 

## Parsing Requests 

HTTP Requests consist of a URL (which may contain parameters) and (optionally) a body which contains more data usually as a form. 

CODE for parsing params (from frag)


CODE for parsing form (from frag)

Afterward a note that you'll want to abstract this since it is shared by all handlers, see Fragmenta for an example of that. 


## Authorise 

There are a few ways of doing this (with Authorise headers, with encrypted cookies, with server-side sessions), but here we're going to consider encrypted cookies which are the most common approach taken by web frameworks in other languages like Rails or Django. 

A simple approach is to check for this cookie in middleware, and then query for it as necessary in the cookie of the request. (link to example, also show snippet here):


### Paths and Security 

Be aware that when you're choosing which routes you , you don't inadvertantly open your app to a directory traversal attack - this attack lets the attacker inject an unexpected path from the file system (for example by using /../ in urls), and. So a naive file handler might serve up unexpected files if given a path which triggered a directory traversal. 

For this reason you should silo only public documents inside your web root at public (never keep private documents there), and never allow handlers to access files outside of this by paths which are tainted by information from the outside world. 


## Acting on Resources

After you have parsed the params and authorised your request, you need to perform some action, usually on a resource associated with this URL. So for example consider the url /users/1/update. You will expect to receive form data associated with a user, identified by the url. 

Show CODE for updating user from the form data...



## Responding to Requests 

Responses to http verbs are typically limited to three options:

1. Respond with an error code (optionally rendering the error)
2. Render a view (for example html or json)
3. Redirect to another endpoint (for example after delete)

### Responding with an error code

Show code for simple example with WriteHeader, then more complex code to render an error template. 



### Rendering a view 

Show code to render a template manually with struct, then compare and contrast with frag approach.

view.AddKey("user",user)
view.AddKey("page",page)


### Redirect to another endpoint 

After deleting a resource, it makes little sense to render a page simply saying deleted, so it is preferable to redirect to the list of resources, to show the updated list without the one just deleted. So for example if you visit:

``` 
/users/1/delete 
```

you might afterward expect to be sent to : /users, you can achieve this quite simply with net/http 

```
http.Redirect("/users",http.Found)
```









