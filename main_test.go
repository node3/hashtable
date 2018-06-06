package main

import (
	"testing"
)

func TestCreateHashTable(t *testing.T) {
	tableSize := 3
	hashtable := CreateHashTable()
	if len(hashtable.table.records) != tableSize {
		t.Fail()
	}
}

func TestHashTable_Put(t *testing.T) {
	hashtable := CreateHashTable()
	hashtable.Put(1, 2)
	present, value := hashtable.Get(1)
	if present == false || value != 2 {
		t.Fail()
	}
	hashtable.Put(1, 3)
	present, value = hashtable.Get(1)
	if present == false || value != 3 {
		t.Fail()
	}
}

func TestHashTable_Get(t *testing.T) {
	hashtable := CreateHashTable()
	hashtable.Put(1, 2)
	present, value := hashtable.Get(1)
	if present == false || value != 2 {
		t.Fail()
	}
	present, value = hashtable.Get(2)
	if present == true {
		t.Fail()
	}
}

func TestHashTable_Del(t *testing.T) {
	hashtable := CreateHashTable()
	hashtable.Put(1, 2)
	result := hashtable.Del(1)
	if result == false {
		t.Fail()
	}
	result = hashtable.Del(1)
	if result == true {
		t.Fail()
	}
	result = hashtable.Del(3)
	if result == true {
		t.Fail()
	}
}

func TestMain(t *testing.T){
	h := CreateHashTable()
	h.Display()
	h.Put(1,2)
	h.Display()
	h.Put(2,3)
	h.Display()
	h.Put(3,4)
	h.Display()
	h.Put(4,5)
	h.Display()
	h.Put(5,6)
	h.Display()
	h.Del(1)
	h.Display()
	h.Del(2)
	h.Display()
	h.Del(3)
	h.Display()
	h.Put(3,4)
	h.Display()
	h.Put(4,5)
	h.Display()
	h.Put(5,6)
	h.Display()
	h.Del(4)
	h.Display()
	h.Del(5)
	h.Display()
	h.Put(11,12)
	h.Display()
	h.Put(12,13)
	h.Display()
	h.Put(13,14)
	h.Display()
	h.Put(14,15)
	h.Display()
	h.Put(15,16)
	h.Display()
}



