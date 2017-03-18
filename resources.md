# Resources

Resources are the nouns which browsers create, retrieve, update and delete when visiting a URL (Uniform Resource Locator) - for a web application they are usually the most important concepts in your application, and which resources you choose to split has profound effects. They are therefore a top-level concern and the part of your app which is most exposed to visitors. They are usually referred to as models in MVC systems like Rails. 

You can reflect their importance in your Go app by making them your most important packages, as discussed in the previous chapter. If you follow this approach, you'll find that you can group everything to do with a resource in its package. If you take this approach you will want to split handlers which act on this resource into a sub-package, so that they can refer to other resources if necessary. This isolates resource code in its own package, forcing you not to tie resources together into one big ball of mud to avoid cyclic dependencies. 

### What should resources do? 

Debate rages (and will continue to rage) in communities like Rails as to where all your code should go, with many insisting it has to either live exclusively in a fat model or in a fat controller, with the other reduced to a supporting role. Other even more baroque architectures are suggested and discussions on responsibilities and structure can become heated and abstract when they should be practical. Adding lots of new files to your project can be satisfying but ultimately complexity doesn't go away, it just moves around. Even worse, when applied blindly to simple apps, architectural patterns like Hexagonal Architecture simply filter complexity through several layers of indirection - a baroque masterpiece of software architecture with hundreds of empty rooms, all alike. 

Instead of building abstractions, start with the simplest thing that will work, and work up from there, and consider the poor reader of your code (who may well be you in 6 months) - can they trace execution through one or two files and guess what a handler does?

### Splitting responsibilities

You can divide responsibilities in many ways, and as discussed above this is often a matter of shifting complexity around rather than getting rid of it, so try to choose the simplest breakdown that will work for your app. If writing a small service for example, you might not need packages at all, for larger apps, organisation becomes more important. So here is a proposed breakdown for a complex web application:

* Views - rendering what the user sees
* Handlers - http, talking to resources, rendering views
* Resources - data, persistence, most logic associated with that data

### Representation in Views

Your resources might not know anything about views, but they will probably require some helper functions to ease presentation. You might for example offer a StatusDisplay method to display the string value of a status for a user rather than an integer value (e.g. Live or Suspended), or an URLShow method to provide the url for the user profile page.

```html
<div class="user_profile">
  <a href="{{ .user.URLShow }}">
  <h4>{{ .user.Name }}</h4>
  <h5>{{ .user.Summary }}</h5>
  <h5>{{ .user.StatusDisplay }}</h5>
  </a>
</div>
```


### Handlers

Your resources might not know anything about views, but they will probably require some helper functions to ease presentation. You might for example offer a StatusDisplay method to display the string value of a status for a user rather than an integer value (e.g. Live or Suspended), or an URLShow method to provide the url for the user profile page.

```go
// HandleShow displays a single user.
func HandleShow(w http.ResponseWriter, r *http.Request) error {

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

	// Authorise access to this page
	err = can.Show(user, session.CurrentUser(w, r))
	if err != nil {
		return server.NotAuthorizedError(err)
	}

	// Render the template
	view := view.NewRenderer(w, r)
	view.CacheKey(user.CacheKey())
	view.AddKey("user", user)
	return view.Render()
}
```

## The User Resource 

The resource package itself for User defines a user resource, the methods that act on it, and any associated functions. This package may also contain folders for associated assets like templates, javascripts and stylesheets, to keep everything grouped by resource. 


```go 

// User handles saving and retreiving users from the database
type User struct {
	ID      int64
  Status  int64
	Role    int64
	Email   string
	Name    string
	Summary string

	PasswordHash    string
	PasswordResetAt time.Time
}

// NewWithColumns creates a new user instance and fills it with data from the database cols provided.
func NewWithColumns(cols map[string]interface{}) *User {
  user := New()
	user.ID = resource.ValidateInt(cols["id"])
  ...
}


// Find fetches a single user record from the database by id.
func Find(id int64) (*User, error) {
  // Query is a query builder library used in this application, 
  // which creates and executes a database query. 
	result, err := Query().Where("id=?", id).FirstResult()
	if err != nil {
		return nil, err
	}
	return NewWithColumns(result), nil
}

```

Because Go does not have generics, and no automatic unboxing of arrays, you need to either use reflection in the ORM (as in gorp for example), or add some functions to your model which explicitly convert from database rows to your model. I prefer to explicitly generate code in the models which reads from the database, as stating this explicitly makes it clearer what is going on and avoids the use of struct tags or reflection. 
