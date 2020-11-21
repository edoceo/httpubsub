/**
 *
 */

package main

import (
	// "fmt"
	"net/http"
	"sync"
)

/**
 * A Channel "class"
 */
type PubSub_Channel struct {
	id string
	sync sync.RWMutex
	list map[string]PS_Client
}

/**
 * Send Stuff to this Channel
 */
func (ch PubSub_Channel) Send(t []byte) {

	// fmt.Printf("PubSub_Channel.Send(t) (%d subs)\n", len(ch.list));

	// @todo Make this Non-Blocking to Instantly Respond if there are no listeners
	// And we need to remove the channel once we've done this
	// So, maybe we need to lock?

	// Should Delete While Iterating?
	// key_list := keys ch.list // how
	// https://bitfieldconsulting.com/golang/map-iteration
	// Then I can keys and delete that item after I pump
	// https://stackoverflow.com/questions/21362950/getting-a-slice-of-keys-from-a-map
	ch.sync.Lock()
	defer ch.sync.Unlock()

	for cl, sub := range ch.list {
		sub.pump <- t
		delete(ch.list, cl)
	}


}

/**
 * Add Subscriber to this Channel
 */
func (ch PubSub_Channel) Sub(w http.ResponseWriter) {

	// fmt.Printf("Sub(w)\n")

	// Subscriber Client
	c := PS_Client{}
	c.id = create_ulid()
	c.pump = make(chan []byte)

	// Add to Channel
	ch.sync.Lock()
	defer ch.sync.Unlock()

	ch.list[c.id] = c

	// Wait for a write to this channel
	body := <-c.pump
	w.Write(body)

	// fmt.Printf("PubSub_Channel.Sub() <<<\n")

}
