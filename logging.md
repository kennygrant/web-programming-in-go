# Logging in web apps

It's possible to completely skip logging in your web app, but it can be invaluable in tracking down. 

### Logging in development 

You'll probably want to do a lot more logging in development, when working on an app locally. You can do this just sending strings to stdout, and there is no need for more sophisticated techniques when logging locally. 

### Logging at scale

As more people use your service, logs become a serious resource sink, and are less and less valuable if they are spread around lots of different servers. In this case it's better to agreggate your logging, and send only the most important events (errors for example) to logs. 