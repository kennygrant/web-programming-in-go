# Go Middleware

## Router support 


### Authentication 




### Logging 

Use fragmenta server/log as an example here of middleware to log and add a request identifier. 


### Overuse

There is a performance overhead to using middleware on every request, and it should not be overused for that reason, but there is also a complexity overhead. It adds implicit hidden behaviour to your application, which is difficult to reason about as it is not easy to trace execution through lots of different handlers. If you can do something explicitly in your handler, and it's not something you want to do on every handler, it's probably better to handle it explicitly rather than building middleware which is sometimes applied, or has conditional logic in it. 


### References 

* Mat article 
