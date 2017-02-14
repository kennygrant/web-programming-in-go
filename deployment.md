# Deployment 

Deployment is often a source of confusion for beginners coming to go, as there are no clear guidelines for deploying applications, and the practice can differ greatly depending on the scale of your app. VPS is often discounted nowadays as being too complex or too expensive for beginners, but it's a good way to get a handle on exactly what is required for your app to run in more complex environments.  

## Guidelines

* Make regular backups, and test they work
* Run tests before deploy
* Always test on an identical staging server before deploys
* Deploy with a script, without manual steps
* Run a CI server so that whenever changes are made tests are run separately. 
* Cross-compile and deploy a binary, to avoid dependencies on the server 

## Backups 

Make regular backups of your data, and make sure your backups work by restoring them and testing the restore. One way to do this if your data is not sensitive is to restore the data into your dev environment periodically. 

## Cross-compile or compile on the target

Go makes cross-compilation incredibly simple, so it may be easier to simply compile your application before deploy, since you know your local machine has all the dependencies you have tested with. If you choose to compile on the server, you will need the same version of Go there, and the same dependencies, at the same version to ensure predictable builds. For this reason I prefer cross-compiling.  

## Testing 

Maintain a test suite which has good coverage, and run tests frequently, certainly before every deploy. Ideally run a continuous integration environment so that your tests are always run when changes are made, and you receive an email if anything breaks. 

## Staging 
If you're deploying an application into production, you should have a staging environment. This is simple to set up, and not as complex as production canaries. 


## Deployment Targets 

### VPS (Digital Ocean etc)

This is cheap and fast, and will get you familiar with operating servers, which is always a useful skill. So it's worth considering for any personal projects, and even for. golangnews.com is hosted this way. 

Systemd is now the default init system on Ubuntu LTS, CoreOS, and most other flavours of Linux. 


```service 

system d unit file here. 





```


### Docker 

The advantage of Docker is that it isolates dependencies, which can be a real problem when using languages which require a separate runtime (for example C, Ruby, Python, Java). For Go this is less of a problem because the runtime is compiled into your binary. 

The disadvantage of docker is that it adds considerable complexity to your deploys, and another part of the system which may fail. If you don't need to deploy more than a few instances of your app, it's probably not worth considering Docker unless you just want to learn about it. 

Mention there are other complexities like Kubernetes, and that this space is still evolving quickly and is only really useful if you are deploying at scale. 

Add a simple example of Docker deploy here.


### Platforms 

PAAS platforms allow you to. The disadvantages are that they tie you in to particular ways of doing things (storage, APIs etc), and it becomes harder to port your app to different platforms. For example if you start using GAE storage, it will be harder to port it over to AWS later. 

An example of this type of service (or rather a level above it, PAAS) is heroku, built on top of AWS. Below is an example of deploying a go app to heroku. 





## References 

* [Deploying with docker](https://blog.golang.org/docker) at the Go Blog
* [Deploying to AWS](https://aws.amazon.com/sdk-for-go/) at Amazon
* [Deploying to GAE](https://cloud.google.com/appengine/training/go-plus-appengine/hello-appengine) at GAE
