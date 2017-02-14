# Resources

Resources are the nouns which browsers create, retrieve, update and delete when visiting a URL (Uniform Resource Locator) - for a web application they are usually the most important concepts in your application, and which resources you choose to split has profound effects. They are therefore a top-level concern and the part of your app which is most exposed to visitors. They are usually referred to as models in MVC systems like Rails. 

You can reflect their importance in your Go app by making them your most important packages, as discussed in the previous chapter. If you follow this approach, you'll find that you can group everything to do with a resource in its package. You may want to split handlers which act on this resource into a sub-package, so that they can refer to other resources if necessary. 

### What should resources do? 

Debate rages (and will continue to rage) in communities like Rails as to where all your code should go, with many insisting it has to either live exclusively in a fat model or in a fat controller, with the other reduced to a supporting role. Other even more baroque architectures are suggested and discussions on responsibilities and structure can become heated and abstract when they should be practical. Adding lots of new files to your project can be satisfying but ultimately complexity doesn't go away, it just moves around. Even worse, when applied blindly to simple apps, architectural patterns like Hexagonal Architecture simply filter complexity through several layers of indirection - a baroque masterpiece of software architecture with hundreds of empty rooms, all alike. 

Instead of building abstractions, start with the simplest thing that will work, and work up from there, and consider the poor reader of your code - can they trace execution through one or two files and guess what a handler does? This poor reader will be you in 6 months when you return to your app, so you had best treat him well. 

### Splitting responsibilities

With this in mind you should allow resources control of data, persistence, and some helper functions for presenting or querying data (for example user.Admin or user.HasImage). 

This leaves processing params, authorisation, querying other models about relations, and responding to requests to the handlers. 


### Accessing the database 



Your resources should handle persisting themselves to the database, using whatever database you prefer. Some ORMs are available or you can simply use. I would personally recommend avoiding struct tags for structuring data, as they are magic comments and quickly become cluttered with fragile information on database access, validation and export into formats like json. So much so that [tools](https://github.com/fatih/gomodifytags) have been created just to manage them.

A better approach is to use functions on your models to marshal into and out of different representations (like the database, json, etc) performing validation as necessary. 

Below is an example of a simple function to populate the fields of an object from a database:


Here is an example of the same thing with struct tags from the x project (take from popular project):


Many people are happy with struct tags, so if you prefer them, or wish to use an ORM which requires them, you may want to use the tool above or a similar tool to manage them. 


### Representation in Views

Your resources might not know anything about views, but they will probably require some helper functions to ease presentation.  

EXAMPLES IN CODE FROM USER OBJECT


## The User Resource 

As an example of the above approaches, let's consider a resource to model website visitors (users). The full code for the package discussed is at this link. This includes:

Actions - handlers 
Assets - Javascript, Styles 
Views - html templates 
Resource - the resource package handles data + logic for users
  role.go 
  

I think give the whole listing and explain each file with excerpts. 



Give the full users code here linked, then pull out excerpts. The aim here is to cover everything you might need in a resource, and to demonstrate talking to the database a little (though we'll already have done that some in a basic way in the db chapter) - here we want to demonstrate why a query builder and the right functions in your model are important. 

Also we should discuss limitation of Go - because there are no generics, and no automatic unboxing of arrays, you need to either use reflection in the ORM (as in gorp for example), or add some functions to your model which explicitly convert from database rows to your model (as in frag or that sqlboiler perhaps?)


## References

* SQLBoiler - https://github.com/vattle/sqlboiler
* sqlx driver 
* Mysql driver
* BoltDB - a pure go key/value store
* qldb - a pure go sql database 

