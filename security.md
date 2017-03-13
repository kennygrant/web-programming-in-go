# Security
Security in web apps is a difficult topic. Most web apps contain multiple undiscovered vulnerabilities, and most examples available don't go into any detail on how to mitigate them. While it is difficult to enumerate all vulnerabilities and this is a fraught subject, ignoring it and hoping for the best is not going to help. Below is an outline of the areas you'll need to think about. 

The concepts of Authentication and Authorisation are linked, but it is useful to distinguish them. 

* Authentication is how you know a user is who they say they are 
* Authorisation is how you control what they can access, based on their authenticated user account. 

So authentication is required in order to usefully control authorisation. 

### Authentication 

There are a few methods of authentication possible, the most useful ones are:

* Sessions stored in an encrypted cookie
* Sessions stored server-side in a database or other store (this may cause problems when scaling)
* Random tokens sent via email for passwordless login 
* OpenAuth (for example using twitter, github or facebook to control logins)
* Secret keys (API keys), usually sent in an Authorise: header with each request 

All of these are valid approaches and are suited to different types of app. For example when controlling API access secret keys which can be easily rotated are the normal way of controlling access. 

For this simple app we're going to explore using the first approach, storing sessions in cookies, which is commonly used for web application development. 

CODE

OAuth is another popular approach, let's try out providing oauth access:

OATH 1 (twitter)

OATH 2 (github)



### Authorisation 

We're going to explore simple role based authorisation using an integer role key. You can also use groups and assign users to groups so that they can have multiple roles, but this can get much more complex and has more edge cases, so carefully consider whether you require this or whether simple one-role per user will work for your app. 


CODE


### User Input 


### Form Validation

You should make sure that the user cannot pass unanticipated fields in forms to your update functions, and that the values which are passed are not out of range. An example of this would be passing in a user_id on a form to update a page which contains no such field, and reassigning the ownership of the page using the user_id key. Never trust that what your forms specify is what your handlers will receive, and whitelist which parameters you expect to process both for updating models and for display in the view. 



#### Example vulnerability 

homakov github vuln with setting owner of repo?



### CSRF 

Mention new same-site cookie attribute, now supported in Chrome and many other browsers.



#### Example vulnerability 

Something from any go app which doesn't actually handle CSRF properly (you could have a look at one to see, like the playground for example)?


### XSS 



#### Example vulnerability 

Something from php?


### SQLI 

Avoid using tainted input without extreme vetting (anything the client can influence), that includes header values, user agent values and any sort of parameters (path,form,query). Never to trust user input, and always parse from the strings coming in to the strictest type you can as soon as possible (for example convert what you believe to be an integer to an integer type immediately before passing to other functions). 

There are some great examples of SQLI over at the [Rails SQLI](http://rails-sqli.org/) site - these are extant issues, which arise from misuse of the Rails query builder and it's convenient param parsing. 

The most common pitfalls are:

* Using an order param directly, rather than using a whitelist of orders
* Allowing user input to set the table or column selected
* Not properly quoting user input for query strings (always use ? never string concat)
* Magically converting user input into a different type depending on its structure 
* Evaluating user input as code (for example json vulnerabilities)

All of these have in common allowing user input to influence your queries to the database. Always sanitise user input, and where possible use a whitelist of commands which are chosen according to the user input, rather than attempting to sanitise the strings that are fed in.


#### Example vulnerability 

Something from Rails 


### Encryption 

See this [talk](https://golangnews.com/stories/1469) and [Go Crypto examples](https://github.com/gtank/cryptopasta) on the web for guidance on which encryption options to use when and how to use them.
 
### References 

The author of this title is not a security expert, I welcome comments and corrections to the above from people who know more about network security than I do. 

Some references to example code for authorisation and crypto:


* [Go Crypto Explained](https://golangnews.com/stories/1469) A talk from George Tankersley on Crypto in Go
* [Go Crypto](https://github.com/gtank/cryptopasta) examples from George Tankersley at Coreos
* [Gorilla Sessions](https://github.com/gorilla/sessions)
* [Same-Site Cookies](https://www.netsparker.com/blog/web-security/same-site-cookie-attribute-prevent-cross-site-request-forgery/)
* [Gorilla CSRF](https://github.com/gorilla/csrf)
* [Preventing CSRF](https://elithrar.github.io/article/preventing-csrf-attacks-in-go/)
* [Fragmenta Auth (sessions, cookies, csrf, passwords)](https://github.com/fragmenta/auth) 
