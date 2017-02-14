# Deployment 

Deployment is often a source of confusion for beginners coming to go, as there are no clear guidelines for deploying applicaitons, and the practice can differ greatly. 

If you're deploying an application into production, you should have a staging environment. 

Backups 


## Cross-compile or compile on the target

Go makes cross-compilation incredibly simple, so it may be easier to simply compile your application before deploy, since you know your local machine has all the dependencies you have tested with. 


## Testing 

## Systemd init system 

Systemd is now the default init system on Ubuntu LTS, CoreOS, and most other flavours of Linux. 


## Docker 

The advantage of Docker is that it isolates dependencies, which can be a real problem when using languages which require a separate runtime (for example C, Ruby, Python, Java). For Go this is less of a problem because the runtime is compiled into your binary. 

The disadvantage of docker is that it adds considerable complexity to your deploys, and another part of the system which may fail. If you don't need to deploy more than a few instances of your app, it's probably not worth considering Docker yet. 

Add a simple example of Docker deploy here.


## Platforms 




## PAAS 


## GAE


## Common Pitfalls

If you choose to compile on the server, you will need the same version of Go there, and the same dependencies, at the same version to ensure predictable builds.  

Use reproducible deploys with a script, you should never have any part of your deployment process be manual. Run your tests as part of this deployment process. 

If possible look into using a CI server to continuously test your changes against a clean environment, to ensure you're not introducing problems. 

Make regular backups of your data, and make sure your backups work by restoring them and testing the restore. One way to do this if your data is not sensitive is to restore the data into your dev environment periodically. 

## References 

* [Deploying with docker](https://blog.golang.org/docker) at the Go Blog
