# Handlers 

Handlers control the flow of requests coming in to the server, and typically should perform a sequence of actions something like:

1. Parse Request 
2. Authorise 
2. Act on Resources
3. Respond 

## Parsing Requests 

HTTP Requests consist of a URL (which may contain parameters) and (optionally) a body which contains more data usually as a form. 

```go
// Add query string params from request
queryParams := r.URL.Query()
for k, v := range queryParams {
  params.Add(k, v)
}

// If the body is empty, return now without error
if r.Body == nil {
  return params, nil
}

// Parse based on content type - different types must be parsed differently. 
contentType := r.Header.Get("Content-Type")
if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
  err := r.ParseForm()
  if err != nil {
    return nil, err
  }
  // Add the form values
  for k, v := range r.Form {
    params.Add(k, v)
  }

} else if strings.HasPrefix(contentType, "multipart/form-data") {
  err := r.ParseMultipartForm(20 << 20) // 20MB
  if err != nil {
    return nil, err
  }

  // Add the form values
  for k, v := range r.MultipartForm.Value {
    params.Add(k, v)
  }

  // Add the form files
  for k, v := range r.MultipartForm.File {
    params.Files[k] = v
  }
}

```

Since you will probably be doing this a lot, you may want to pull it out into a function, or use a library that does this already. Most of the popular routing packages will offer functions for parsing the request parameters. 


## Authorise 

There are a few ways of doing this (with Authorise headers, with encrypted cookies, with server-side sessions), but here we're going to consider encrypted cookies which are the most common approach taken by web frameworks in other languages like Rails or Django. 

A simple approach is to check for this cookie in middleware, and then query for it as necessary in the cookie of the request. Here is an example of middleware to do this. 

```go 
// Middleware sets a token on every GET request so that it can be
// inserted into the view. It currently ignores requests for files and assets.
func Middleware(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// If a get method, we need to set the token for use in views
		if shouldSetToken(r) {

			// This sets the token on the encrypted session cookie
			token, err := auth.AuthenticityToken(w, r)
			if err != nil {
				log.Error(log.Values{"msg": "session: problem setting token", "error": err})
			} else {
				// Save the token to the request context for use in views
				ctx := r.Context()
				ctx = context.WithValue(ctx, view.AuthenticityContext, token)
				r = r.WithContext(ctx)
			}

		}

		h(w, r)
	}

}

```


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
// HandleUpdate handles the POST of the form to update a user
func HandleUpdate(w http.ResponseWriter, r *http.Request) error {

	// Fetch the  params
	params, err := mux.Params(r)
	if err != nil {
		return server.InternalError(err)
	}

	// Find the user
	user, err := users.Find(params.GetInt(users.KeyName))
	if err != nil {
		return server.NotFoundError(err)
	}

	// Check the authenticity token
	err = session.CheckAuthenticity(w, r)
	if err != nil {
		return err
	}

	// Authorise update user
	err = can.Update(user, session.CurrentUser(w, r))
	if err != nil {
		return server.NotAuthorizedError(err)
	}

	// Set the password hash from the password
	hash, err := auth.HashPassword(params.Get("password"))
	if err != nil {
		return server.InternalError(err)
	}
	params.SetString("password_hash", hash)

	// Validate the params, removing any we don't accept
	userParams := user.ValidateParams(params.Map(), users.AllowedParams())

	err = user.Update(userParams)
	if err != nil {
		return server.InternalError(err)
	}

	// Redirect to user
	return server.Redirect(w, r, user.ShowURL())
}

```

### Performing a request 

You may find your handlers need to call out to another service, either synchronously or asyncrhonously. In Go you can easily make a syncrhonous function call async simply be prefixing it with the go keyword. 


```go 
// Get data from  another service synchronously (normally this would be async)
  res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
  // use body
	fmt.Printf("%s", body)
  
```


```go 
// Post data to another service asynchronously with the go keyword 
go postData()

func postData() {
  data := url.Values{}
  data.Set("name", "foo")
  data.Add("surname", "bar")

  client := &http.Client{
    Timeout: 10*time.Second
  }
  r, _ := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
  r.Header.Add("Authorization", APIAuthToken)
  r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

  resp, _ := client.Do(r)
  fmt.Println(resp.Status)
}

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









