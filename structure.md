# App Structure

The structure of your app depends on the purpose; command line tools, libraries, and web applications all have a different focus and are suited to a different structure. The discussion below is of package layout for server-side web applications, so it may not be applicable to other types of application, and even within web application, there is disagreement as to the best layout. You can take comfort in the fact that this is largely subjective, and as long as you follow some simple rules should not matter too much. 

The structure of your application should:

* Be consistent and straightforward 
* Be simple to navigate 
* Make the important self-evident 

Go has several unique features which constrain app layout:

* Every app must have one package main, which contains the entrypoint. 
* Packages are based on folders, so every folder within your app is a new package. 
* Packages may contain many files, so you can use files to break up code by purpose.
* Cyclic dependencies are not allowed, so a cannot import b which imports a.
* Package names are used to call package functions/data, so they should read together. 

Various patterns have been proposed for web application layout, the most popular ones are:

1. Grouping resources into one package called models or similar, and actions into another 
2. Isolating resources into their own packages, and actions with them in a subpackage 


### Grouping resources in one package 


#### Advantages
* Resources can share dependencies in this package (for example the database), and can freely refer to each other. 
* Avoids stutter on names by using generic model. prefix. 

#### Disadvantages
* The Model package becomes very large and complex with say > 10 resources 
* One package contains logic and data for unrelated concerns like page tags and user roles. 
* Resources become inextricably linked and can affect each other in unpredictable ways. 
* All function names have a generic prefix like models.NewUser etc. 


### Isolating resources in their own packages 

#### Advantages
* Resources cannot refer to each other 
* Resources are separated and complexity controlled as the package is only about one resource
* Resource packages can also include views, scripts, styles which are intimately tied to the resource. 
* Handlers have responsibility for collecting, mutating and linking resources, controlling complexity. 


#### Disadvantages
* Resources cannot refer to each other (cyclic dependencies). 
* Handlers must be kept in a separate package (I use a subpackage). 
* Model names can stutter (users.User). 

As you can see, both of these have attractions, neither is wrong in a fundamental way and the choice is somewhat subjective. This book is going to focus on the second approach, but it's not too important which you choose - choose whichever you feel most comfortable with. 


## Dependencies 

When your application grows to any significant size, you'll find your handlers need to know about some shared state in order to function. There are various approaches to this, none of which is obviously wrong, again this is quite a subjective choice. 


### Wrap handlers and inject dependencies as scoped variables 

```go
handler := NewFooHandler(db, etc...)

type FooHandler struct{
   db DB
}

func(f *FooHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
  f.db.Save("hello, world") 
}
```


### Inject into objects and attach handlers as methods

```go
handler := NewFooHandler(db, etc...)

type FooHandler struct{
   db DB
}

func(f *FooHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
  f.db.Save("hello, world") 
}
```


### Store in the importing package

For example in Buffalo models package there is a DB connection using the pop ORM, which is set up in init. This is used by all models to store themselves in the database. 

```go
package models
...
// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection
```

Testing with this approach would require testing with handlers. 

### Store in the imported package

So for example on startup the app would call 

```go
mail.Setup(config.Get("from"),config.Get("secret"))
```

and set up the mail package. 

The importers simply call:

```go
email := mail.New("recipient@example.com")
mail.Send(email)
```

with the information on the message they want to send, without worrying about how the message is actually sent. This is my preferred option as it scales with the number of imported packages, and keeps knowledge of setup and internals in the packages concerned, rather than spreading it around in calling packages. The callers to a log package or a mail or db package should not have to worry about which loggers or mail servers they are sending to in most circumstances. 

Testing with this approach just means calling setup methods before tests with either mock setup details or alternative test account details for the service concerned (for example a test database).



* [Style Guidelines for Packages](https://rakyll.org/style-packages/)
* [Standard Package Layout ](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1#.k8xq7h19a)