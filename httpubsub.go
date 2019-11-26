/**
 * Absolutely Inspired by PatchBay.Pub - https://news.ycombinator.com/item?id=21639066
 */

package main

import (
	// "fmt"
	"strings"
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

	// var p = r.URL.Path
	body := <- c.pump
	w.Write(body)

}

/**
 *
 * @type {[type]}
 */
func dpsRouter(w http.ResponseWriter, r *http.Request) {

	// Find Which Channel It Is
	p := strings.Trim(r.URL.Path, "/");

	// Get this Channel from the Map
	// or create, if not found
	c := ps_client_list[p]
	if ("" == c.id) {
		c.id = p
		c.pump = make(chan []byte)
		ps_client_list[p] = c
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

	ps_client_list = make(map[string]PS_Client)

	// http.HandleFunc("/info", viewInfo)
	// http.HandleFunc("/status", viewInfo)
	http.HandleFunc("/", dpsRouter)
	http.ListenAndServe(":8080", nil)

}
