# Serving files in Go

The standard library package [net/http](https://golang.org/pkg/net/http/) allows you to set up an http server with just a few lines of code. We will use this as a basis to build a full stack web application in this guide.

### Listen and Serve

Let's define a handler function - this is what responds to http requests after they come in, and writes a response on the wire. The function receives two parameters - http.ResponseWriter which is an interface we can write our response to and a pointer to an http.Request which contains the information about this request.

We start by writing a one-line comment above the function, to let the reader \(perhaps yourself in 6 months\) know what it does. This comment can be used by the Go tools to generate documentation for us automatically later. Inside the function, we print a formatted string  containing Hello and whatever the request path was, escaping it first. Typically we'd use templates to write data, but this is a toy example to demonstrate the server.

```go
// handler says hello and echoes the request path
func handler(writer http.ResponseWriter, request *http.Request) {

    fmt.Fprintf(writer, "Hello, %q", html.EscapeString(request.URL.Path))

}
```

Next we set up our server to use this handler function to respond to requests, and ask it to start listening. ListenAndServe function takes a port value \(as a string\), along with a Router \(or a ServeMux as  \(we can leave that nil for now to use the default ServeMux\).

```go
// Attach a function to the default ServeMux/Router
// for the path / and any path under it
http.HandleFunc("/", handler)

// Ask the http package to listen on port 3000
err := http.ListenAndServe(":3000", nil)
if err != nil {
   log.Fatal(err)
}
```

You can find the code for this first lesson in the examples folder of this repo at 1-Serve/serve.go. Open the first file and paste in the code above, then run it with go run serve.go.

You can then open your web browser at port 3000, and find your web server talking to you. Go ahead and visit [localhost](http://localhost:3000/world), I'll wait here. You'll notice that the server calls your handler for any path since it was attached to / - this is a feature of the default router, which you can change later if you wish.

![](/assets/hello-world.png)

### Listen and Serve with TLS

One oversight in the standard library defaults, which cannot now be changed because of the Go 1 guarantee, is that timeouts are not set by default. You can read more about this [here](https://blog.cloudflare.com/exposing-go-on-the-internet/). So in a slight addition to our code above, we're going to use a new method ListenAndServeTLS to serve using a certificate. We'll need to generate a self-signed cert first. If you have openssl available you can use that:

```bash
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
```

Now that you have the key.pem and cert.pem files in the same directory as serve.go, replace the line creating the server with this:

```go
// Set up a new http server with some default timeouts and our desired 
server := &http.Server{
        // Set the port in the preferred string format
    Addr: ":3000",
}

// Ask the http package to listen with TLS
err := http.ListenAndServeTLS("cert.pem", "key.pem")
if err != nil {
   log.Fatal(err)
}
```

As a nice bonus, Go supports HTTP/2 - HTTP/2 is enabled automatically if you use TLS and the client supports it.

## Free and Automatic Let's Encrypt certificates

The Autocert library allows you to automatically request TLS certs for your domains and serve your website with TLS without having to deal with a certificate authority.

```
// Ask the http package to listen on port 3000
err := http.ListenAndServe(":3000", nil)
if err != nil {
log.Fatal(err)
}
```



