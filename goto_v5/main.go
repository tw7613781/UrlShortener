package main

import (
	"fmt"
	"net/http"
	"flag"
	"net/rpc"
)

const AddForm = `
<html>
<form method="post" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="add">
</form>
</html>
`

var (
	listenAddr = flag.String("http", ":8080", "http listen address")
	dataFile = flag.String("file", "store.json", "data store file name")
	hostname = flag.String("host", "localhost:8080", "host name and port")
	rpcEnabled = flag.Bool("rpc", false, "enable RPC server")
	masterAddr = flag.String("master", "", "RPC master address")
)

var store Store

func main() {
	flag.Parse()
	if *masterAddr != "" {
		store = NewProxyStore(*masterAddr)
	} else{
		store = NewURLStore(*dataFile)
	}
	if *rpcEnabled {
		rpc.RegisterName("Store", store)
		rpc.HandleHTTP()
	}
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(*listenAddr, nil)
}

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		fmt.Fprintf(w, AddForm)
		return
	}
	var key string
	if err := store.Put(&url, &key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "http://%s/%s", *hostname, key)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	var url string
	if err := store.Get(&key, &url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	http.Redirect(w, r, url, http.StatusFound)
}
