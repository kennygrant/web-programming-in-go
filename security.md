# Security
Security in web apps is a difficult topic. Most web apps contain multiple undiscovered vulnerabilities, and most examples available don't go into any detail on how to mitigate them.



### Authentication 





### Authorisation 



### Form Validation

You should make sure that the user cannot pass unanticipated fields in forms to your update functions, and that the values which are passed are not. An example of this would be passing in a user_id on a form to update a page which contains no such field, and reassigning the ownership of the page using the user_id key. Never trust that what your forms specify is what your handlers will receive, and whitelist which paramaters you expect to process both for updating models and for display in the view. 

### CSRF 



### XSS 




### SQLI 

Avoid using tainted input without extreme vetting (anything the client can influence), that includes header values, user agent values and any sort of parameters (path,form,query). Never to trust user input, and always parse from the strings coming in to the strictest type you can as soon as possible (for example convert what you believe to be an integer to an integer type immediately before passing to other functions). 

There are some great examples of SQLI over at the [Rails SQLI](http://rails-sqli.org/) site - these are extant issues, which arise from misuse of the Rails query builder and it's convenient param parsing. 

The most common pitfalls are to use an order param, 

### Encryption 

See this [talk](https://golangnews.com/stories/1469) and [Go Crypto examples](https://github.com/gtank/cryptopasta) on the web for guidance on which encryption options to use when and how to use them. 
 
### References 

Some references to example code for authorisation and crypto:

* [Go Crypto](https://github.com/gtank/cryptopasta) examples from George Tankersley at Coreos
* [Gorilla sessions](https://github.com/gorilla/sessions)
* [Gorilla csrf](https://github.com/gorilla/csrf)
* [Fragmenta auth package](https://github.com/fragmenta/auth) 
