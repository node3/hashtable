package main

import (
	"testing"
)

func TestCreateHashTable(t *testing.T) {
	tableSize := 3
	hashtable := CreateHashTable(tableSize)
	if len(hashtable.table) != tableSize {
		t.Fail()
	}
}

func TestHashTable_Put(t *testing.T) {
	tableSize := 3
	hashtable := CreateHashTable(tableSize)
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
	tableSize := 3
	hashtable := CreateHashTable(tableSize)
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
	tableSize := 3
	hashtable := CreateHashTable(tableSize)
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


