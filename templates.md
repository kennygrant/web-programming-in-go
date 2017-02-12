# Go HTML Templates

The Go Standard Library has simple templating for text files and html files built-in which is more than adequate for most uses. 


### Helper functions

You can define helper functions. 



### Data

Data can be passed in to templates and referenced using the dot syntax. So for example if you pass in a user as the template context, you'll be able to access it within the template like so:

 
```go
  {{ .user }}
```

As your templates get more complex, and require more and more objects, you can either create a struct which contains all the objects you'd like to reference, or use a structure like a map store the keys available to the template. I prefer the map approach as it allows more flexibility, at the cost of some type safety.

### Nested templates 

The templates allow you to refer to other templates as long as they are registered on the parent. One way to use this is to register all your templates on the same set, so that every template can refer to every other one. To render a template within another template (assuming templates are stored with a relative path as their key), use:

```go
 {{ template "path/to/template" . }}
```

### Layouts and Partials 

Go templates don't have explicit support for the concept of layouts, though there as an addition of blocks in Go this doesn't map well to use for partials. You can emulate layouts and partials though by parsing all your templates as one set, and then rendering the layout, which then renders the included partials in turn.  


Layout file:
```go
  <html>
  <body>
  <header></header>
  <section>{{.content}}</section>
  <footer></footer>
  </html>
```


Partial file to place into content key when rendering:
```go
  {{ .user.Name }}
```

This requires rendering the partial first, then placing it in a context for rendering in the main layout. 


## Escaping 

Go HTML templates treat data values as plain text which should be encoded so they can be safely embedded in an HTML document. The escaping is contextual, so actions can appear within JavaScript, CSS, and URI contexts.

The package also defines some typed strings which you can use to declare content as safe without escaping for certain contexts - html.HTML, html.JS and html.URL. Be very careful that your content can never include user input when using this feature. 

### Loading templates 

Let's load the templates we need from our app. We'll load all the templates into one set, apply our functions to them, and store them ready for rendering later. To get hot reload of templates as they change during development, we can simply reload them on every request. 


## References 

Have a look at the documentation for the [text/template](https://golang.org/pkg/text/template/) package as well as the [html/template](https://golang.org/pkg/html/template/) package, as the base text template has much more detail on operators, whereas the documentation for the html template focusses mostly on escaping content.