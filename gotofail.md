#goto fail 

This chapter lists a few problems encountered by programmers new to Go. First, if you haven't already, you should go and read Effective Go, it covers many of the topics below in greater detail.  


### Why can't I use inheritance?

If you're trying to make inheritance work in Go, please give up now. The language has been designed without it and is better for it. James Gosling, who created Java, when asked what he would change if he could do Java over again, memorably said:

<blockquote>"I'd leave out classes"</blockquote>

Go allows you to compose objects, but this is not the same as inheritance. If find yourself creating an Abstract Base Class, superclass or an AbstractFactoryBean, step away from the keyboard. 

### Errors 

Never ignore errors (either with _ or by omitting to assign them). Never use panic in place of error handling (except in very simple command line apps). 

Don't use panic as a poor substitute for exceptions, it was included for truly exceptional circumstances. Go doesn't have exceptions by design, which is probably a good thing - as an indication of current thinking on this current Google guidelines for C++ disallow the use of exceptions in that language as well. 

### Why can't I ignore GoPath

This may change in future, but for now the tools rely on GOPATH in order to make installs easier and find your code. So put your code in GOPATH and be happy. 

### Why is each folder a package?

By convention, in Go each folder is a package of code. Code within the package may refer to any unexported identifiers, so you can split your code up into files within the package. You can't package code any other way, so structure your code into folders along with other code which is intimately connected, and export the minimum number of identifiers possible. 


### Why do I see panic: assignment to entry in nil map?

If you declare an array, you don't need to use make, but if you declare a map, you must initiallise it first, so either use:

```go
// Create an empty map 
m := make(map[string]string)

// Create a map with values 
m := map[string]string {
  "foo":"bar",
}
```

### How to check if a map contains a key in go?

Use the comma ok idiom to check if a key exists in a map on one line. 

```go
m := map[string]string{"foo":"bar"}
if v, ok := m["foo"]; ok {
    //val has a value here
    // as ok was true
}
```

### Why are my map keys out of order?

Map key order is not guaranteed, and is shuffled in newer versions of Go. This means if you want sorted keys and values from your map you'll need to sort them yourself and then grab the values in a for loop. Unfortunately the collection types in Go are somewhat limited. 


### Reading a file 

The simplest method to read a file is with ioutils:

```go
// Load the config file at this path
configFilePath := "./config.json"
file, err := ioutil.ReadFile(configFilePath)
if err != nil {
  log.Fatalf("error loading:%s %s", configFilePath, err)
}
```

This loads the entire file into memory. Obviously for large files that's not desirable, so you should instead read them in chunks. A good example of this is in the io package, if you look at the source of io.Copy you'll see it copies a file in chunks. 

```go

    // open the input file 
    r, err := os.Open("in.txt")
    if err != nil {
        panic(err)
    }
    defer r.Close()

    // open the output file 
    w, err := os.Create("out.txt")
    if err != nil {
        panic(err)
    }
    defer w.Close()

    // copy data between them 
    // note w doesn't have to be a file, it could be any io.Writer
    n, err := io.Copy(w, r)
    if err != nil {
        panic(err)
    }
    log.Printf("Copied %v bytes\n", n)
    
```

### Strings and Bytes

You should generally try to work with []byte where you have data in that format, and avoid conversions between strings and bytes. Often you will find you need to convert between the two though. This is fairly straightforward. 

```go
// Convert bytes to a string 
s := string([]byte{1,2})

// Convert a string to bytes 
b := []byte("hello")

// Convert a string to a byte array 
var b [32]byte
copy(b[:],"hello")

// To range over the runes in a string (note this is not the bytes)
for i,r := range s {
  //...
}
```

Write strings with "", runes with '', and multiline strings with ``

Watch out for strings.Trim - this takes a cutset (a set of characters to cut), not a string to trim. So you use it like this:

```go
strings.TrimRight("wat????hello?","?eloh")
// output: "wat" 
```

### Dates 

Beware time.AddDate() - if you're adding months you may get unexpected results due to rollover rules. Give a few examples here and a mitigation - workaround always add to 1st of month, then add days in a separate operation. Actually fixing is harder and requires a new library.   


### Strings and Ints 

Converting an int to a string using the strconv package:

```go
  var s string 
  var i int 

  i = 99

  // Convert an int to a string
  s = strconv.Itoa(i)
  fmt.Printf("%q",s)

  // Convert an int to a string
  i,err := strconv.Atoi(s)
  if err != nil {
    panic(err)
  }
  
  fmt.Printf("%d",i)
  
  // You can also use Sprintf 
  fmt.Sprintf("%d->%s",i,s)
  
```

### Struct Tags

Struct tags can be used to define the translation from Go fields to json, xml and database columns. You don't have to use them and I'd recommend choosing an ORM which doesn't rely on them, as declaring relations in string tags attached to fields is error prone and fragile.

### Why are comments used as directives?

It's true, comments can be used to: 

* //go:generate bar - for a code generator
* //import "foo" - set the canonical import path
* //Output: foo - less crucially, document the result for an example function.

This is ugly and fragile, and doesn't even have a consistent syntax, but you're unlikely to add comments with exactly this syntax and break something. Perhaps in Go 2.0 they'll tidy some of this up. 

