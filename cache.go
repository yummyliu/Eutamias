package Eutamias

import (
	"fmt"
	"time"
)

type Entry struct {
	Object     interface{}
	Expiration int64
}

func (ent Entry) Expired() bool {
	if ent.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > ent.Expiration
}

type Cache struct {
	ents map[string]Entry
}
func (c *Cache) Put(k,v string, oot time.Duration) error{
	c.ents[k] = Entry{
		Object : v,
		Expiration : int64(oot),
	}
	return nil
}

type MyError struct {
	When time.Time
	what string
}

func (e *MyError) Error() string {
	return fmt.Sprint("at %v, %s", e.When, e.what)
}

func (c *Cache) Get(k string) (Entry,error) {
	v := c.ents[k]
	err := &MyError{
		time.Now(),
		"map error",
	}
	return v,nil
}
