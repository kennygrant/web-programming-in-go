# Serving files in Go

The standard library package [net/http](https://golang.org/pkg/net/http/) allows you to set up an http server with just a few lines of code. We will use this as a basis to build a full stack web application in this guide. 

### Listen and Serve

The ListenAndServe function takes a port value \(as a string\), along with a Router \(or a ServeMux as  \(we can leave that nil for now to use the default ServeMux\). 

```go
http.Handle("/foo", fooHandler)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))
```

You can find the code for this first lesson in the examples folder of this repo. Open the first file and paste in the code above, then run it with go run serve.go. 

### Listen and Serve with TLS

One oversight in the standard library defaults, which cannot now be changed because of the Go 1 guarantee, is that timeouts are not set by default. You can read more about this [here](https://blog.cloudflare.com/exposing-go-on-the-internet/). So in a slight addition to our code above, we're going to use a new method ListenAndServe HTTP/2 is enabled automatically if you use TLS. 





## Free and Automatic Let's Encrypt certificates

The Autocert library allows you to automatically request TLS certs for your domains and serve your website with TLS without having to deal with a certificate authority. 

