# Interesting Go projects

One of the most effective ways to improve as a developer is to read code other people have written. This can be intimidating at first, but Go is a simple language to learn and to read. Below are listed some interesting Go projects which you may want to use, read about or refer to when developing in Go. This is by no means a definitive or exhaustive list, but hopefully will give you an idea of the breadth of software  being written in Go.

## [The Go Standard Library](https://golang.org/pkg/)

This is a great place to start for examples of idiomatic Go and has a lot of packages useful for web development. Start with he simpler packages and work up to net/http, and read the source by clicking on the title of any function. The extensive documentation of the Go stdlib is worth reading in detail before you start using a package, there are also links to the source code which is also worth reading for examples of idiomatic Go.

[sort](https://golang.org/pkg/sort) - sorting functions for arbitrary data
[strings](https://golang.org/pkg/sort) - string manipulation
[strings](https://golang.org/pkg/bytes) - string manipulation
[strings](https://golang.org/pkg/io) - input/output - look at the simple interfaces here like io.Reader
[context](https://golang.org/pkg/context) - for cancellations, deadlines and request scoped values 
[sql](https://golang.org/pkg/database/sql/) - The common interface for db drivers
[net/http](https://golang.org/pkg/net/http/) - A web server and client, and lots of helpers - read this package carefully.

Also have a look at the unofficial packages at [golang.org/x](https://godoc.org/-/subrepo), which are not held to the Go 1.0 stability guarantee and can evolve a little faster (though they are very stable). There are many interesting packages in there, including an html parser. 

## Developing  

* [Go Imports](https://github.com/golang/tools/tree/master/imports) - automatically set the imports used
* [SourceGraph](https://sourcegraph.com/) - code as a global graph
* [Go Doc](https://godoc.org) - Search for Go Packages

## Servers 

* [Minio](https://github.com/minio/minio) - sore photos, videos, VMs, containers, log files, or any blob of data as objects, compatible with S3
* [Terraform](https://github.com/hashicorp/terraform) - managing infrastructure
* [Vault](https://github.com/hashicorp/vault) - store secrets
* [Kubernetes](https://github.com/kubernetes/kubernetes) - container scheduling
* [Camlistore](https://camlistore.org/code) - cloud storage system
* [upspin](https://github.com/upspin/upspin) -  a global name/sharing system

## Databases 

* [Influxdb](https://github.com/influxdata/influxdb) - A time-series database written in Go
* [Boltdb](https://github.com/boltdb/bolt) - A Key value store written in Go
* [etcd](https://github.com/coreos/etcd) - Distributed reliable key-value store for the most critical data of a distributed system

## Version Control 

* [Go git](https://github.com/src-d/go-git) A highly extensible Git implementation in pure Go.
* [gitgo](https://github.com/ChimeraCoder/gitgo) - A Go implementation of Git functions
* [Gogs](https://github.com/gogits/gogs) - a port of github written in Go

## Graphics 

Michael Fogleman has written many great graphics projects in Go, the first three packages below are his work.

* [Primitive](https://github.com/fogleman/primitive) - Reproducing images with geometric primitives.
* [ln](https://github.com/fogleman/ln) - 3D line art engine.
* [pt](https://github.com/fogleman/pt) - This is a CPU-only, unidirectional path tracing engine written in Go. It has lots of features and a simple API.
* [Gopherize.me](https://gopherize.me/) - create and download your own personalised Gopher avatar
* [glhf](https://github.com/faiface/glhf) - A Go package that makes life with OpenGL enjoyable.
* [identicon](https://github.com/barthr/Identicon) - 


## About Frameworks 

A brief word on frameworks - there is some hostility in some quarters to the idea of frameworks, simply because they are a huge dependency pulling in lots of code you won't need, often hide many dependencies, and can make it very easy to get started and very hard to finish your project if you pick the wrong one.

You can easily create a go service without a framework, and for a very simple service, this should probably be your preferred route. For larger services, if you're working in a large team and need a common consensus on design decisions, or if you're creating a large number of services, you will either end up creating your own ad-hoc, poorly specified version of half of these frameworks, you owe it to yourself to look at the alternatives and consider whether they can teach you anything, or at least show you the ground that needs to be covered. 

## Microservice frameworks 

* [Gizmo](https://github.com/nytimes/gizmo) - A Microservice Toolkit from The New York Times
* [Go Kit](https://github.com/go-kit/kit) - A standard library for microservices
* [Go Micro](https://github.com/micro/go-micro) - A pluggable RPC framework for microservices
* [Goa](https://goa.design/) - API first microservice design

## Web frameworks 

* [Gorilla](https://github.com/gorilla/) - a web toolkit for creating Go websites
* [Fragmenta](https://github.com/fragmenta) - code generation and libraries for Go websites
* [Gin](https://github.com/gin-gonic/gin) - an HTTP web framework written in Go
* [Echo](https://github.com/labstack/echo) - High performance, minimalist Go web framework 

### Community

Some more useful links to continue learning about programming in Go:

* [The Go Forum](https://forum.golangbridge.org/) - a friendly place to ask questions about programming in Go
* [Golang News](https://golangnews.com/) - fresh links about Go programming every day
* [Golang Weekly](http://golangweekly.com/) - a weekly email newsletter with links about Go
* [Go on Github](https://github.com/golang/go) - file issues, and consult the wiki for dev tips
* [Go Documentation](https://golang.org/doc/) - a treasure trove of blog posts and articles about Go
* [Idiomatic Go](https://dmitri.shuralyov.com/idiomatic-go) - a list of do's and dont's from Dmitri Shuralyov
