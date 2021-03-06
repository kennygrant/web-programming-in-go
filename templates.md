# Go HTML Templates

The Go Standard Library has simple templating for text files and html files built-in which is more than adequate for most uses. The documentation is quite thorough, and worth reading in its entirety, but be sure to read the docs for [text/template](https://golang.org/pkg/text/template/) as well as [html/template](https://golang.org/pkg/html/template/), as the base text template has much more detail on operators, whereas the documentation for the html template focusses mostly on escaping content.


### Helper functions

You can define helper functions in Go for each template, however most applications will have a common set of functions which are applied to all templates. 

While you can put a lot of functions into your templates, you should try to restrict these to formatting or passing data in from global data, rather than containing too much logic. Most of your logic should live in the resources and your handlers, with helper functions simply used to format data in resources. 

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

```html
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

This requires rendering the partial first, then placing it in a context for rendering in the main layout. Here is an example:

 ```go 
 
  // Render content first, then add to context for rendering layout 
  renderContent(writer io.Writer, 
                layout string, 
                template string, 
                context map[string]interface{}) error {
  
    // First, render the template 
    tmpl := Templates[template]
    var rendered bytes.Buffer
    err := tmpl.Execute(rendered, context)
    if err != nil {
      return err
    }
    // Save the template content as html 
    context["content"] = template.HTML(rendered.String())

    // Now render the layout, using template content key .content
    layout := Templates[layout]
    err := layout.Execute(writer, context)
    if err != nil {
      return err
    }
    
  }
 ```


## Escaping 

Go HTML templates treat data values as plain text which should be encoded so they can be safely embedded in an HTML document. The escaping is contextual, so actions can appear within JavaScript, CSS, and URI contexts.

The package also defines some typed strings which you can use to declare content as safe without escaping for certain contexts - html.HTML, html.JS and html.URL. Be very careful that your content can never include user input when using this feature. 

### Loading templates 

Let's load the templates we need from our app. We'll load all the templates into one set, apply our functions to them, and store them ready for rendering later. To get hot reload of templates as they change during development, we can simply reload them on every request. 


```go 
// Load the templates under a given directory - start with a base template 
// all templates under will be named by relative path from root
htmlTemplateSet = template.New("").Funcs(got.FuncMap(helpers))

err = filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {

  if err != nil {
    return err
  }

  // Deal with files
  if !info.IsDir() {
    if strings.HasSufffix(path,".tmpl") {
      // Add to our template set - avoid duplicates 
      contents, err := t.readFile(t.fullpath)
    	if err == nil {
    		return err
    	}
      _, err = htmlTemplateSet.New(path).Parse(contents)
      if err == nil {
    		return err
    	}
    }
  }

  // Return nil on directories to recurse
  return nil
})

```

## Common Pitfalls

Go templates require other templates to be loaded in the same context to be referenced. You may want to define one root set of templates and include all your templates in it. 

Go templates use prefix notation even for operators like *and* (they are defined as functions), so they come before both arguments. Thus you use:

```go
{{if  and .page.Published .user.Admin }}
// NOT {{if .page.Published and .user.Admin }}
```

Also, the operators do not short-circuit evaluation, all arguments are evaluated before the function is called, this means you cannot rely on something like this to work if .user is nil:

```go
{{ if and .user .user.Admin }}
```

you must instead nest the if check:

```go
{{ if .user}} {{ if .user.Admin }} {{ end }} {{ end }}
```

When writing headers, you should always call WriteHeader once (never twice), and always call it after calling SetHeader, or SetHeader will be ignored (CHECK).


## References 

* [text/template](https://golang.org/pkg/text/template/) - basic templating without escaping 
* [html/template](https://golang.org/pkg/html/template/) - templates with contextual escaping for html,js,css etc.
