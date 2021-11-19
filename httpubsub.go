/**
 * Absolutely Inspired by PatchBay.Pub - https://news.ycombinator.com/item?id=21639066
 */

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/oklog/ulid"
)

// I like these as IDs
func create_ulid() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	u, _ := ulid.New(ulid.Timestamp(t), entropy)
	return u.String()
	// Output: 0000XSNJG0MQJHBF4QX1EFD6Y3
}

//
type PS_Client struct {
	id   string
	pump chan []byte
}

var pubsub_channel_list = New_PubSub_Channel_List()

/**
 * POST/Publish Handler
 */
func pub(w http.ResponseWriter, r *http.Request, p string) {

	fmt.Printf("pub(w, r, %s)\n", p)

	ch, err := pubsub_channel_list.Find(p)
	if err != nil {
		w.Write([]byte("No Subscribers\n"))
		return
	}

	// Now, Publish to Subscribers
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// Error of some type
		// body = ""
		w.Write([]byte("Read Error"))
		return
	}

	ch.Send(body)
	// pubsub_channel_list.Delete(ch.id)

}

/**
 * GET/Subscribe Handler
 */
func sub(w http.ResponseWriter, r *http.Request, p string) {

	// Try to see if this Channel is already register
	fmt.Printf("sub(w, r, %s)\n", p)

	ch, _ := pubsub_channel_list.Find(p)
	if "" == ch.id {
		ch = pubsub_channel_list.Create(p)
	}

	ch.Sub(w)

	// Should Drop Channel & Client Here

	fmt.Printf("sub() <<<\n")
}

/**
 * Route the Path to a Channel via Map
 */
func dpsRouter(w http.ResponseWriter, r *http.Request) {

	// Find Which Channel It Is
	p := strings.Trim(r.URL.Path, "/")

	if 0 == len(p) {
		// Write HTML And Exit
		var html = "<html><head><title>httpubsub</title></head><body><h1>httppubsub</h1><p>Specify a path to Publish or Subscibe to</p><pre>curl http://localhost:8080/sub123\n  ** waiting **</pre><p>Then in another terminal:</p><pre>curl -X POST http://localhost:8080/sub123</pre><p>And watch the subscribe side emit any posted data</p>"
		w.Write([]byte(html))
		return
	}

	// PUBSUBing
	switch r.Method {
	// case "DELETE":
	// 	del(w, r, c)
	// 	break;
	case "GET":
		sub(w, r, p)
		break
	case "POST":
		pub(w, r, p)
		break
	default:
		// Error
		break
	}

}

func main() {

	port := flag.String("port", "", "Listen Address (default is \":8080\" for standard, \":8443\" for TLS)")
	// secret := flag.String("secret", "", "A Secret Value")
	cert := flag.String("cert", "", "A PEM formatted SSL/TLS Certificate")
	certKey := flag.String("cert-key", "", "A Key for the SSL/TLS Certificate")
	flag.Parse()

	// http.HandleFunc("/info", viewInfo)
	// http.HandleFunc("/status", viewInfo)
	http.HandleFunc("/", dpsRouter)

	// SSL, we hope
	if len(*cert) > 0 {
		if 0 == len(*port) {
			*port = ":8443"
		}
		err := http.ListenAndServeTLS(*port, *cert, *certKey, nil)
		if err != nil {
			panic(err)
		}
	} else {
		if 0 == len(*port) {
			*port = ":8080"
		}
		err := http.ListenAndServe(*port, nil)
		if err != nil {
			panic(err)
		}
	}

}
