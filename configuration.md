# Configuration

Your application needs to know certain facts about its environment when it executes, for example it might need to know if it is running in development or in production, which database to talk to, or which mail service to use. There are a few ways to feed this information to the app, and the most appropriate depends on where it runs and how many instances you have. Sometimes your app might support several ways to set the same information. 

## Flags 

The simplest way to pass preferences in to an application is via command line subcommands or flags. If you just need to choose the port your app serves on (for dev or production), or set which mode it runs in, flags are a simple way to change this behaviour at runtime. 

```go
  import "flag"
  
  // Config holds our app config
  type Config struct {
  	Port int
  }
  var config = &Config{Port: 3000}
  
  func main() {
    // Read the port with the flag package into a pointer to config.Port
    flag.IntVar(&config.Port, "port", 3000, "The port the server listens on (default 3000)")
    flag.Parse()
    fmt.Println("Starting server on port:", config.Port)
  }
  
```

## Environment variables

You can also read environment variables in Go. These can be set in your systemd service unit file for example, which might be more convenient than flags. Be aware that environment variables can leak information, and can also be set in some circumstances by web servers (for example the HTTP_USER_AGENT variable). This means they may contain hostile data. For more detail on this read up on the [Shellshock](https://en.wikipedia.org/wiki/Shellshock_(software_bug)) attacks of 2014.


```go

  // Read the port from the environment variable, if set
  portEnv := os.Getenv("MY_SERVER_PORT")
  if len(portEnv) > 0 {
    // Convert string to an int
    config.Port, err = strconv.Atoi(portEnv)
    if err != nil {
      log.Print(err)
    }
  }

```


## Configuration file

As your configuration grows, you might decide to store it in a file. If this file includes secrets like keys or api user data, you should probably keep it out of your version control system, just in case it leaks. If you choose to use a config file you may not need to set environment variables or command line flags, or you may choose to let the user override some settings (say the port) with a flag. 

If your app is likely to be started by end users, it's probably friendlier to use something like flags, if it is to be started solely as a service (hopefully in an automated way), it's easier to provide config with a file or external service. 

### File Formats

There are many possible file formats for configuration files, and which is popular often depends more on the tools available than the inherent qualities of the file format. Most of them are flawed in some way (for example JSON doesn't support comments and doesn't allow trailing commas, YAML is white-space/indent sensitive). The Go stdlib provides json parsing, so that is often used by Go apps, we're going to explore that here. 

### Loading JSON 

To load a JSON config file, since you know it will be a small file, you can simply load it all into memory. 

```go

	type Config struct {
		Port int
	}
	config := Config{Port: 80}

	// Load the config file at this path
	configFilePath := "./config.json"
	file, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("error loading:%s %s", configFilePath, err)
	}

  // Read the data into the config struct 
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("error reading:%s %s", configFilePath, err)
	}

  // Print the contents
	log.Printf("Config read:%v", config)
```


You can load the config either into an untyped map (stringly typed), or into a struct designed for that purpose. A map is more flexible as if the config is missing keys or has malformed keys the app will not refuse to start, but it requires discipline in using the keys - you should always assert the type when using them. 

For reasons of security and data sanity it's important that whenever you receive data from external input, you assert the type as soon as possible and convert them to a real Go type. This goes for input from any source which might be untrusted as you build out your app (user names, config files, parameters)- a significant number of vulnerabilities in web apps are caused by passing around data which is loosely typed and forgetting to assert that it conforms to the type expected. 

For example, do this:

```go
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
```

not this:

```go
err = http.ListenAndServe(config.Get("my_port_key"), nil)
```


## Remote configuration 

If you're running multiple instances of your app in an automated way, it might make more sense to put your configuration into a central store, which you only need to update once. This is the sort of problem you'll only run into if your app is actually successful, so for the vast majority of projects it's not a huge concern, at least initially. 


### etcd


```go

	etcd := etcdClient()

	key := "/foo"
	value := "bar"

	// Set a key
	resp, err := etcd.Set(context.Background(), key, value, nil)
	if err != nil {
		log.Fatal(err)
	}

	// print metdata to confirm set
	log.Printf("Set is done. Metadata is %q\n", resp)

	// Get the same key from etcd instance
	res, err := etcd.Get(context.Background(), key, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("%q key has %q value\n", res.Node.Key, res.Node.Value)

  ...

  // etcd returns a new etcd client
  func etcdClient() client.KeysAPI {
  	cfg := client.Config{
  		Endpoints:               []string{"http://127.0.0.1:2379"},
  		Transport:               client.DefaultTransport,
  		HeaderTimeoutPerRequest: time.Second,
  	}
  	c, err := client.New(cfg)
  	if err != nil {
  		log.Fatal(err)
  	}
  	return client.NewKeysAPI(c)
  }

```


### Vault 

Vault handles keys, tokens, passwords and certs transparently and in a secure way. It offers an API for accessing an encrypted Key/Value store with data in it. It also offers generation and rollover of tokens. This tool is too complex to cover in an introductory text, but you can head over to the [website](https://www.vaultproject.io) and read more about the project there. 

### Viper 

You can see an example of all these different approaches in the Viper project, 
and you may want to import it (or copy bits of it to your project). However beware of reaching for complex dependencies like this without a good understanding of what they are buying you, and the tradeoffs involved (maintenance, security, understanding vs saving time initially by hiding complexity). This library is a very good extended example of all the different ways to read configuration though. 


## References 


* [os](https://golang.org/pkg/os)  - the package in the Go standard library which allows access to environment variables. 
* [flag](https://golang.org/pkg/flag) - parse command line flags 
* [json](https://golang.org/pkg/encoding/json/) - you can parse a json config file with this library 
* [etcd](https://github.com/coreos/etcd)  - a key-value store designed for sharing configuration values
* [vault](https://github.com/hashicorp/vault)  - a tool for managing secrets
* [Viper](https://github.com/spf13/viper) - Viper lets you read config from many sources at once.