# Databases 

While there are some limitations, database/sql provides a good baseline interface for drivers, and has seen improvements in Go 1.8. Go has many pure go database drivers for popular databases, so you don't need to be limtied in your choice. 

### Which database should you use?

There are many databases available, if you are looking for a relational database (and you should be, unless you have very specific requirements), Postgresql is reliable, performant and has a lot of features which rival NoSQL ones. If you do only require document data, consider starting with a pure Go key/value store like BoltDB, but be wary of unknowingly and over time building half a relational database without the ACID constraints within your app.

# Connecting to the Database



