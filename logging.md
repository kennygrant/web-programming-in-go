# Logging in web apps

It's possible to completely skip logging in your web app, but it can be invaluable in tracking down bugs, and it's useful to have some form of persistent log. A common approach is to log to a file, however this has limitations if you want to run your server in a container or across several servers, so you might want to consider either adopting structured logging and sending lots to a separate time series database (like influxdb), or simply sending logs to stdout. 

### Logging in development 

You'll probably want to do a lot more logging in development, when working on an app locally. You can do this just sending strings to stdout, and there is no need for more sophisticated techniques when logging locally. Often the best way to track down a bug. 

### Logging across boundaries 

Across service boundaries, or even just to trace execution in handlers and goroutines, you may want to assign a request id to your requests with logging middleware which can then be used to trace them when logging. 


### Logging at scale

As more people use your service, logs become a serious resource sink, and are less and less valuable if they are spread around lots of different servers. You may find it is better to centralise your logging and log far less except in case of errors. This problem only affects larger applications though, so you should worry about it if you reach this scale and start to split your app across servers, not before. 


See logging articles by Peter Bourgon etc. 


References 

Logrus 
gokit log 
new project logging 
influxdb 
grafana 
[12 Factor logging](https://github.com/agnivade/funnel)