# Web Programming in Go

This book will guide you through the process of building a complex web app with Go, similar to one you might build in other programming languages with Rails, Django or Flask. Readers coming from other languages or frameworks should not find it difficult to translate this knowledge to Go, but there is a dearth of examples of the details required to handle all the mundane activities of a web app - authentication, authorisation, templates, error handling, async operations etc.

### Code examples

All the code used in this guide is available here. You can fetch it with the command :

```
go get https://github.com/kennygrant/web-programming-with-go/examples
```

which will download the files to $GOPATH/src/github.com/kennygrant/web-programming-with-go/examples

### Caveats

This is not an introductory guide to the Go language, so some prior knowledge of Go is required. You should first try the [Tour](https://tour.golang.org/welcome/1) at golang.org and read the [Documentation](https://golang.org/doc/) before starting. You should also have a working install of Go, and be familiar with concepts like the GOPATH before starting. If you wish to read a great introduction to the Go language, the definitive guide to the language is [The Go Programming Language](http://www.gopl.io/) by Alan A. A. Donovan and Brian W. Kernighan.

This book is not a guide to writing command line apps, microservices, or APIs in Go. While some of the principles may be applicable to those domains, the focus is on writing serve-side web apps which serve html to browser clients. Obviously not all lessons applicable to this domain transfer to others, particularly in areas like authentication or project structure. 

There are often arbitrary choices to be made \(for example about project layout, or which library to use\), and this book is not intended to be a prescriptive guide to the only way to produce Go applications. Suggestions for alternative approaches are welcome and may be included in future editions. 

The exercises in this book have been tested on Mac OS X and Linux and some advice may be OS specific, but the vast majority of code should run on Windows as well as linux. If you spot an error or have a suggestion for improvements, please let me know by filing an issue against the github repo for the book.

### Let's build something useful

As you follow along with the exercises in this book, you're going to build a service very like twitter, complete with handles, avatars, and inane wittering. You can see the finished product over at gophr.club. 

[Screenshot]

