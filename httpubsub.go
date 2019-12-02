/**
 * Absolutely Inspired by PatchBay.Pub - https://news.ycombinator.com/item?id=21639066
 */

package main

import (
	"flag"
	"fmt"
	"strings"
	"sync"
	"io/ioutil"
	"net/http"
)

// type Publisher struct {
// }
//
// type Subscriber struct {
//
// }

type PS_Client struct {
	id string
	pump chan []byte
}

var ps_client_sync sync.RWMutex
var ps_client_list map[string]PS_Client

/**
 * POST/Publish Handler
 */
func pub(w http.ResponseWriter, r *http.Request, c PS_Client) {

	// Now, Publish to Subscribers
	body, err := ioutil.ReadAll(r.Body)
	if (err != nil) {
		// Error of some type
		// body = ""
	}
	c.pump <- body

}

/**
 * GET/Subscribe Handler
 */
func sub(w http.ResponseWriter, r *http.Request, c PS_Client) {

	// Wait for a write to this channel
	body := <- c.pump
	w.Write(body)

}

/**
 * Route the Path to a Channel via Map
 */
func dpsRouter(w http.ResponseWriter, r *http.Request) {

	// Find Which Channel It Is
	p := strings.Trim(r.URL.Path, "/");

	// Get this Channel from the Map
	ps_client_sync.RLock()
	c := ps_client_list[p]
	ps_client_sync.RUnlock()

	// or create, if not found
	if ("" == c.id) {
		c.id = p
		c.pump = make(chan []byte)

		ps_client_sync.Lock()
		ps_client_list[p] = c
		ps_client_sync.Unlock()

	}

	switch (r.Method) {
	// case "DELETE":
	// 	del(w, r, c)
	// 	break;
	case "GET":
		sub(w, r, c);
		break;
	case "POST":
		pub(w, r, c);
		break;
	default:
		// Error
		break;
	}

}

func main() {

	hpsPort := flag.String("port", "", "Listen Address (default is \":8080\" for standard, \":8443\" for TLS)")
	crtFile := flag.String("tls", "", "A PEM formatted SSL/TLS Certificate")
	keyFile := flag.String("tls-key", "", "A Key for the SSL/TLS Certificate")
	flag.Parse()

	ps_client_list = make(map[string]PS_Client)

	// http.HandleFunc("/info", viewInfo)
	// http.HandleFunc("/status", viewInfo)
	http.HandleFunc("/", dpsRouter)

	// SSL, we hope
	if (len(*crtFile) > 0) {
		if (0 == len(*hpsPort)) {
			*hpsPort = ":8443"
		}
		err := http.ListenAndServeTLS(*hpsPort, *crtFile, *keyFile, nil)
		if err != nil {
			panic(err)
		}
	} else {
		if (0 == len(*hpsPort)) {
			*hpsPort = ":8080"
		}
		err := http.ListenAndServe(*hpsPort, nil)
		if err != nil {
			panic(err)
		}
	}

}
