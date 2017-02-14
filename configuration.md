# Configuration

Your application needs to know certain facts about its environment when it executes, for example it might need to know if it is running in development or in production, which database to talk to, or which mail service to use. There are a few ways to feed this information to the app, and the most appropriate depends on where it runs and how many instances you have. 

## Flags 

The simplest way to pass preferences in to an application is via command line subcommands or flags. If you simply want to choose from a few alternatives. 


CODE


## Environment variables

You can read environment variables in Go. These can be set in your systemd service unit file. 


CODE


## Configuration file

If you don't wish to or cannot set environment variables when deploying, you may find a configuration file is a good balance. You will need to deploy the configuration file with your app. 



CODE


### Remote configuration 

If you're running multiple instances of your app in an automated way, it might make more sense to put your configuration into a central store, which you only need to update once. 



CODE?


## References 

* os - the package in the Go standard library which allows access to environment variables. 
* flag - parse command line flags 
* json - you can parse a json config file with this library 
* etcd - a key-value store designed for sharing configuration values