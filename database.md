# Databases 

If your program does much of anything, it will probably need to store data. How you store the data is not a foregone conclusion, but you should try to pick a store which suits the sort of data you have, and will have in the future. For many apps a relational database is a good choice, and will cover a multitude of use cases adequately. 

There are drivers available for almost every mainstream database using go and the database/sql package as a basis, but research the driver as well as the database to ensure you're happy with it before committing. 

## Which database should you use?

Relational databases structure data into rows and are only accessible through the use of SQL. This query language can seem obtuse and old-fashioned, but is based on fundamental relationships between the data, and is well worth exploring in depth. Other databases use different data structures suitable for the particular sort of data they store, which tend to be inflexible if you want to relate records or perform complex queries. If you think you'll need complex queries or relations over your data, or you're just not sure what to choose, start with a relational database.

### SQL Databases
 
There are many databases available, but if you are looking for a relational database (and you should be, unless you have very specific requirements), Postgresql is reliable, performant and has a lot of features which rival NoSQL ones like jsonb - this would be my first choice. If you do only require document data, consider starting with a pure Go key/value store like BoltDB, but be wary of building half of a poorly specced relational database within your app. It's very hard to get storing data right. 

### Time Series 

If you're storing a huge number of events which are primarily indexed on timestamps (for example errors from another app, )

### Key Value store 

If you're just storing some simple values and don't need to relate them or query over large numbers of them in a sophisticated way, a key value store like Bolt is a good choice. Bold accesses all data with a byte key, this means BoltDB is very fast to read and write data but provides no support for joining data, so be sure that you won't want to relate items in your database (for example fetch all records which belong to a certain user).

### NoSQL

There are certain clearly defined uses where a NoSQL database makes sense. If you have a lot of records which are simple records, the contents of which change are heterogenous and change a lot over time, and don't need to query relations between them, a NoSQL database might make sense. 


## Connecting to the Database

```go

// Set up some options to open the database, these should be pulled from your config,
// not hardcoded as here. 
options := map[string]string{
  "adapter":  "postgres",
  "user":     "hunter",
  "password": "hunter2",
  "db":       "hunter_db",
  "host":     "localhost",                          // for unix instead of tcp use path - see driver
  "port":     "5432",                               // default PSQL port
  "params":   "sslmode=disable connect_timeout=60", // disable sslmode for localhost, set timeout
}

// Use an option string to connect to the database to allow passing additional parameters to psql
optionString := fmt.Sprintf("user=%s %s host=%s port=%s dbname=%s %s",
  options["user"],
  options["password"],
  options["host"],
  options["port"],
  options["db"],
  options["params"])

var err error
db, err = sql.Open(options["adapter"], optionString)
if err != nil {
  return err
}

// Call ping on the db to check it does actually exist!
err = db.Ping()
if err != nil {
  return err
}

if db != nil {
  fmt.Printf("Database %s opened using %s\n", options["db"], options["adapter"])
}
```

### Executing SQL 

To execute simple SQL is fairly straightforward, just pass the sql, along with any params to replace placeholders, to the database, and check for errors or results (if for example you're expecting a count in results).

```go
// Exec the given sql and args against the database directly
// Returning sql.Result (NB not rows)
uid := 1
sql := "delete from users where id=$1"
results, err := database.Exec(sql, uid)
if err != nil {
  log.Printf("query: error executing:%s error:%s",sql,err)
}
```

### Building Queries

To get data out of the database, you need to construct queries. These might be as simple as "select * from users where status=100", but typically you'll want to adjust the query based on the requirements of the particular request (for example depending on the filters applied or the order requested).
```go

// Fetch some records 
sql := "select * from users"

// If an order was requested, append it 
order := switch params.Get("order") {
case "name":
  " order by name asc, id asc"
case"status":
  " order by status desc, name asc"
default:
  " order by name asc, status desc"
}
sql = sql + order

// Append limit
sql = " limit 100"

// Then execute query below
rows, err := db.Query(sql, age)
if err != nil {
        log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
    // Scan a row into a resource
    model.NewWithRow(rows)
        var name string
        if err := rows.Scan(&name); err != nil {
                log.Fatal(err)
        }
        fmt.Printf("%s is %d\n", name, age)
}
if err := rows.Err(); err != nil {
        log.Fatal(err)
}

```

You may want to use a query builder to make this process a little less tedious for the many repetitive queries, and to avoid violating the strict syntax of sql (e.g. limit must be last) or store them as constants you can refer to. When constructing queries, be very careful about using user input, for example if selecting a column to sort on, don't use a parameter directly, use a switch statement selecting known good values. Some rules to avoid SQLi:

* Always use the driver parameter substitution, never build strings manually
* Always convert to known types - don't use params without asserting type
* Validate param values fall in known ranges - don't use params without validation 
* Test your application with common examples of sqli



### Fetching Records

```go
uid := 1
sql := "select * from users where id=$1"

// Fetch rows from the database
var users []*User
rows, err := database.Query(sql, uid)
if err != nil {
  log.Printf("query: error fetching results:%s",err)
}

// Close rows before returning
defer rows.Close()

// Fetch the columns
cols, err := rows.Columns()
if err != nil {
  log.Printf("Error fetching columns: %s\n", err)
}

// For each row, construct an entry in results with a map of column string keys to values
for rows.Next() {
  result, err := scanRow(cols, rows)
  if err != nil {
    log.Printf("Error fetching row: %s\n", err)
  }
  u := &User{
    ID: result["id"],
    Name:result["name"],
  }
  
  users = append(users,u)
}


func scanRow(cols []string, rows *sql.Rows) (Result, error) {

	// We return a map[string]interface{} for each row scanned
	result := Result{}
	values := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		var col interface{}
		values[i] = &col
	}

	// Scan results into these interfaces
	err := rows.Scan(values...)
	if err != nil {
		return nil, fmt.Errorf("Error scanning row: %s", err)
	}

	for i := 0; i < len(cols); i++ {
		v := *values[i].(*interface{})
		if values[i] != nil {
			switch v.(type) {
			default:
				result[cols[i]] = v
			case bool:
				result[cols[i]] = v.(bool)
			case int:
				result[cols[i]] = int64(v.(int))
			case []byte: // text cols are given as bytes
				result[cols[i]] = string(v.([]byte))
			case int64:
				result[cols[i]] = v.(int64)
			}
		}

	}

	return result, nil
}

```

You can use an ORM to attempt to map database columns to rows in your struct, but this doesn't gain you very much in Go, and is likely to be slower as it will use reflection and struct tags to map columns to fields in your struct. My preferred approach is to generate the code required to reify models in the resource objects, and extend or adjust it as necessary. 

## References 

* [Basic SQL](https://www.khanacademy.org/computing/computer-programming/sql)
* [Postgresql Relational DB](https://github.com/lib/pq)
* [MySQL  Relational DB](https://github.com/go-sql-driver/mysql/)
* [QL - an experimental Pure Go sql DB](https://godoc.org/github.com/cznic/ql/driver)
* [PGX - an alternative driver for Postgresql](https://github.com/jackc/pgx)
* [InfluxDB Time Series DB](https://github.com/influxdata/influxdb)
* [BoltDB Key Value Store](https://github.com/boltdb/bolt)
* [Tiedot Document Database](https://github.com/HouzuoGuo/tiedot)