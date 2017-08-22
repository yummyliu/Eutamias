package Eutamias

import (
	"testing"
)

func TestGet(t *testing.T) {
	ca := &Cache{
		ents : make(map[string]Entry),
	}
	if v,e := ca.Get("a"); e != nil {
		t.Error("Get has error")
	} else {
		t.Logf("Get Pass {}",v)
	}
}

func TestPut(t *testing.T) {
	t.Log("TestPut")
	ca := &Cache{
		ents : make(map[string]Entry),
	}
	if e := ca.Put("a","a1",0); e != nil {
		t.Error("Put has error")
	} else {
		t.Log("Put pass")
	}
}

