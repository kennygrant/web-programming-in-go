# Routing in Go

Since Go has a rather limited router in the standard library, many people have built their own. There

### Let's build a ServeMux

Because Go is open source, you can go and have a look at the [DefaultServeMux](https://golang.org/src/net/http/server.go?#L1865) yourself in the standard library. This is a fairly simple router which has a few drawbacks -

* It doesn't let you collect parameters at all, name them or limit their content

* It doesn't guarantee the evaluation order of routes

* It doesn't let you define groups of routes

### Performance

While it is tempting to measure something like a router purely on performance, unless it is pathologically slow \(e.g. uses regexp in a naive way or allocates a lot for every request\) it is not likely to take up many resources compared to your handlers which have to talk to the database and write responses. So measures of performance on routers are useful  indicators but should not be your primary concern when choosing one. Factors which will usually be more important are the way it parses parameters, control over evaluation order of routes, support for middleware and the handler signatures.

