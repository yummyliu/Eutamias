package cache

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
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
