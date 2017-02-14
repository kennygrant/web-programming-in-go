# Glossary

This glossary contains terms used in Go web programming which you might not be instantly familiar with. 

* MVC - a popular approach to web programming (borrowed from desktop apps), which splits the app into Models (data, logic), View which presents the data, and Controllers (accepts input, converts data for view).
* Resource - the nouns which web browsers manipulate via http verbs. 
* Handlers - respond to URLS on your server, normally each endpoint is assigned a handler. 
* Router - connects endpoints (URLs) to handlers in your app. 
* TLS - a cryptographic protocol for the net, used by the https protocol. 
* Mux - another term for Router. 
* Actions - a term used in this book for handlers. 
* Models - a component of MVC, referred to as resources in this book. 
* Controllers - a component of MVC, analagous to handlers in this book. 
* Postgresql - a popular SQL relational database, more rigorous than MySQL 
* MySQL - a popular SQL relational database
* Sqlite - a popular SQL relational database, unfortunately C-only so requires CGO
* Middleware 
* Structured logging - sending log entries in the form of key/value pairs, rather than simply strings
* AES - Advanced Encryption Standard used for encryption
 