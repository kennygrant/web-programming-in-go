package main

import (
	"log"
	"time"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

func main() {

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

}

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
