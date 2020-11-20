/**
 * Manages a Channel/Path
 */

package main

import (
	"errors"
	"fmt"
	"sync"
)


// Channel Manager
type PubSub_Channel_List struct {
	channel_list map[string]PubSub_Channel
	sync sync.RWMutex
}

/**
 */
func (cl PubSub_Channel_List) Create(p string) PubSub_Channel {

	fmt.Printf("PubSub_Channel_List.Create(p=%s)\n", p)

	ch := PubSub_Channel{}
	ch.id = p
	ch.list = make(map[string]PS_Client)

	cl.channel_list[p] = ch

	return ch
}

/**
 */
func (cl PubSub_Channel_List) Delete(p string) {
	delete(cl.channel_list, p)
}


/**
 */
func (cl PubSub_Channel_List) Find(p string) (PubSub_Channel, error) {

	fmt.Printf("PubSub_Channel_List.Find(%s)\n", p)

	cl.sync.RLock()
	ch, ok := cl.channel_list[p]
	cl.sync.RUnlock()

	if (ok) {
		return ch, nil
	}

	return ch, errors.New("Channel Not Found")

}
func New_PubSub_Channel_List() *PubSub_Channel_List {
	r := new(PubSub_Channel_List)
	r.channel_list = make(map[string]PubSub_Channel)
	return r
}
