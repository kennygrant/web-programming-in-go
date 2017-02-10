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

Serving a website with TLS over https is a simple matter of calling a different function - ListenAndServeTLS to serve using a certificate and key which we provide the path for. Since this is just an example, an insecure test key and cert are provided - open them to discover they are simply text files with encrypted data. On a real server you would use certificates provided by your Certificate Authority. Don't use self-signed keys on a public server, and certainly never these keys, these examples are provided in order to run the example locally. Since you have the test key.pem and cert.pem files in the same directory as serve.go, replace the line creating the server with this and try launching it again:

```go
// Ask the http package to listen with TLS
err := http.ListenAndServeTLS(":3000", "cert.pem", "key.pem", nil)
if err != nil {
   log.Fatal(err)
}
```

You will then be able to connect to [localhost](https://localhost:3000/tls) using https, though you will have to ignore a warning about your self-signed certificate, which browsers treat as less secure than no certificate at all, this will let you check that your server is using tls though. As a nice bonus, Go supports HTTP/2 - HTTP/2 is enabled automatically if you use TLS and the client supports it.

One oversight in the standard library defaults, which cannot now be changed because of the Go 1 guarantee, is that timeouts are not set by default on the server. You can read more about this [here](https://blog.cloudflare.com/exposing-go-on-the-internet/). So in a slight addition to our code above, we're going to set timeouts on the server. To do this we configure a server instance first, and then call a method on the server to start it, rather than using the default server from the package level function above.

```go
// Set up a new http server with some default timeouts and port
server := &http.Server{
  // Set the port in the preferred string format
  Addr: ":3000",
  // The default server from net/http has no timeouts
  // Set some reasonable timeouts here
  ReadTimeout:  30 * time.Second,
  WriteTimeout: 60 * time.Second,
}

// Start the server with our self-signed cert and key
err := server.ListenAndServeTLS("cert.pem", "key.pem")
if err != nil {
  log.Fatal(err)
}
```

## Free and Automatic Let's Encrypt certificates

The [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert) library allows you to automatically request TLS certs for your domains and serve your website with TLS without having to deal directly with a certificate authority, your server will request certificates from [Let's Encrypt ](https://letsencrypt.org/)Authority \(or any other supporting the ACME protocol\). As you can see from the code below, this is just a few more lines of code.

```go
domains := "example.com www.example.com"
email := "me@example.com"

autocertDomains := strings.Split(domains, " ")
certManager := &autocert.Manager{
  Prompt:     autocert.AcceptTOS,
  Email:      email,                                      // Email for problems with certs
  HostPolicy: autocert.HostWhitelist(autocertDomains...), // Domains to request certs for
  Cache:      autocert.DirCache("secrets"),               // Cache certs in secrets folder
}

// Set up the server with timeouts and the autocert manager
server := &http.Server{
    Addr: ":443", // Set the port 
    ReadTimeout:  30 * time.Second,
    WriteTimeout: 60 * time.Second,

    // Pass in the autocert manager 
    TLSConfig: &tls.Config{
     GetCertificate: certManager.GetCertificate,
    },
}

// Start the TLS server, giving it no cert and key (it uses autocert for this)
err := server.ListenAndServeTLS("", "")
if err != nil {
  log.Fatal(err)
}
```

If you try to run it locally on a linux/unix, you'll get the following error:

```
listen tcp :443: bind: permission denied
```

This is because ports below 1024 are restricted, and you'll need to use setcap to give your app permission. If you do that you'd also receive an error:

```
acme/autocert: host not configured
```

The autocert library requires your server to run on the IP which maps to the domains you want to secure, so that it can use the ACME protocol to confirm the server's identity and request a cert. So unfortunately you can't try it out locally, you need to run it on a server which maps to the domain you want a certificate for. Typically in order to test locally you'd run without TLS on a high port, and then in production you run on a lower port with TLS enabled. We'll cover how to set this up in later chapters as you build your service.

### Congratulations

You have now set up a server which can serve web pages using the latest protocols like http2, and even request its own certificates periodically before they expire. You're on the way to creating your own sophisticated web app with Go.

