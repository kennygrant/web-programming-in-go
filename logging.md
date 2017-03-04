# Logging in web apps

It's possible to completely skip logging in your web app, but it can be invaluable in tracking down bugs, and it's useful to have some form of persistent log. A common approach is to log to a file, however this has limitations if you want to run your server in a container or across several servers, so you might want to consider either adopting structured logging and sending logs to a separate time series database (like influxdb), or simply sending logs to stdout and using other tools to process the ephemeral data your app generates. 

### Logging in development 

You'll probably want to do a lot more logging in development, when working on an app locally. You can do this just sending strings to stderr, and there is no need for more sophisticated techniques when logging locally. Often the best way to track down a bug is to think about why the problem might be occurring (e.g. perhaps your data is not in the correct state, and you're not sure where it changes), and insert some lines of logging at points to verify your assumptions. Nothing more sophisticated than log.Printf() is required for this, here is an example of logging in a web handler. 

```go
func HandleShowBlub(w http.ResponseWriter, r *http.Request) error {
    
    // DEBUG: Verify the request URL
    log.Printf("#debug params:%v",r.URL)

    // Normal code
    u,err := users.Find(1)
    if err != nil {
      return err // log the error and inform the user
    }
    
    // DEBUG: Verify the username for user 1
    log.Printf("#debug username:%s",u.Name)
}
```


### Logging across boundaries 

Across service boundaries, or even just to trace execution in handlers and goroutines, you may want to assign a request id to your requests with logging middleware which can then be used to trace them when logging. This is relatively simple to do with some middleware. 


```go

// RequestID is but a simple token for tracing requests.
type RequestID struct {
	id []byte
}

// String returns a string formatting for the request id.
func (r *RequestID) String() string {
	return fmt.Sprintf("%X-%X-%X-%X", r.id[0:2], r.id[2:4], r.id[4:6], r.id[6:8])
}

// NewRequestID returns a new random request id.
func newRequestID() *RequestID {
	r := &RequestID{
		id: make([]byte, 8),
	}
	rand.Read(r.id)
	return r
}

// ctxKey is used as a unique key for the context
type ctxKey struct{}


// Middleware adds a logging wrapper and request tracing to requests.
func Middleware(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		
    // Store a unique request id for handlers of the request.
    requestID := newRequestID()
  	ctx := context.WithValue(r.Context(), &ctxKey{}, requestID)
  	r = r.WithContext(ctx)

    // Log before handling, and store the time
    log.Printf("Request %s trace:%s",r.RequestURI,requestID)
		start := time.Now()
        
    // Execute the handler which we wrap
    h(w, r)

    // Log after the request 
    log.Printf("Response %s in %s trace:%s",r.RequestURI,time.Now().UTC().Sub(start),requestID)

	}

}

```

Then in your handler, you can extract this stored request id (or another value stored on context), to use it to. 

```go
func HandleShowBlub(w http.ResponseWriter, r *http.Request) error {
    
    // Log the request id for tracing
    requestID := log.GetRequestID(r)
    log.Printf("request:%s",requestID)

    // Perform a task, sending request details onward
    go doSomething(r)

    // Respond to request by writing response 
    ...
}
```

### Logging at scale

As more people use your service, logs become a serious resource sink, and are less and less valuable if they are spread around lots of different servers. You may find it is better to centralise your logging and log to a central database. This problem only affects larger applications though, so you should worry about it if you reach this scale and start to split your app across servers, not before, unless you just want to. 

To store your data, you might consider a time series database like InfluxDB (written in Go)


```go
  
  data := map[string]interface{}{
      "idle":   10.1,
      "system": 53.3,
      "user":   46.6,
  }

  err := stats.Send(data)
  if err != nil{
    log.Errorf("stats: %s",err)
  }

  ...

  // Send sends data to influxdb
  Send(fields map[string]interface{}) error {
  // Create a new HTTPClient
   c, err := client.NewHTTPClient(client.HTTPConfig{
       Addr:     "http://influxhost:8086",
       Username: config.Get("influxuser"),
       Password: config.Get("influxpass"),
   })
   if err != nil {
       return err
   }

   // Create a new batch
   bp, err := client.NewBatchPoints(client.BatchPointsConfig{
       Database:  "mydb",
       Precision: "s",
   })
   if err != nil {
       return err
   }

   // Create a point and add to batch
   tags := map[string]string{"my": "cpu-total"}
   pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
   if err != nil {
      return err
   }
   bp.AddPoint(pt)

   // Write the batch
   err := c.Write(bp)
   if err != nil {
       return err
   }
   
   
  }
```

### Distributed Logging 

If you're interested in logging, you should read the [article](https://peter.bourgon.org/ok-log/) introducing OK Log by Peter Bourgon, which talks in depth about logging at scale (for microservices in particular or just for clusters of many servers), which is a much deeper topic than we have time for here. 

The good news is you won't have to worry about this until your app is successful, so unless you're already operating at scale, it's probably of academic interest only.


### References 

[Standard Library Log](https://golang.org/pkg/log/)
[Logrus](https://github.com/Sirupsen/logrus)
[ok log](https://github.com/oklog/oklog) 
[influxdb](https://github.com/influxdata/influxdb) 
[grafana open source graphing](http://grafana.org/)
[12 Factor logging](https://github.com/agnivade/funnel)