package main

import "fmt"

type record struct {
	key int
	value int
	//left *record
	//right *record
	next *record
	isnull bool
}

var (
	minLoadFactor = 1/4
	maxLoadFactor = 3/4
	tableSize = 2
)


// Binary tree implementation
//func (head *record) insert(node record) (*record){
//	if head == nil {
//		return &node
//	} else {
//		var iterator = head
//		for true {
//			if iterator.key > node.key {
//				if iterator.left == nil {
//					iterator.left = &node
//					break
//				} else {
//					iterator = iterator.left
//				}
//			} else {
//				if iterator.right == nil {
//					iterator.right = &node
//					break
//				} else {
//					iterator = iterator.right
//				}
//			}
//		}
//		return head
//	}
//}

type hashTable struct {
	table []record
	nRecords int			// Number of records
}

func createHashTable() hashTable {
	hashtable := hashTable{table: make([]record, tableSize), nRecords:tableSize}
	for i:=0; i< tableSize; i++ {
		hashtable.table[i].isnull = true
	}
	return hashtable
}

func (h hashTable) hashFunction(key int) (int) {
	return key%len(h.table)
}

func (h hashTable) put(key int, value int) {
	index := h.hashFunction(key)
	iterator := h.table[index]
	prev := record{0, 0, nil, true}
	for !iterator.isnull {
		if iterator.key == key {	// Key already exists
			iterator.value = value
			prev.isnull = false
			return
		}
		prev = iterator
		iterator = *iterator.next
	}
	node := record{key, value, nil, true}
	iterator.next = &node
	iterator.isnull = false
	//h.checkLoadFactorAndUpdate()
}

func (h hashTable) get(key int) (bool, int) {
	index := h.hashFunction(key)
	iterator := h.table[index]
	for !iterator.isnull {
		if iterator.key == key {	// Key already exists
			return true, iterator.value
		}
		iterator = *iterator.next
	}
	return false, -1
	//h.checkLoadFactorAndUpdate()
}

func (h hashTable) del(key int) (bool) {
	index := h.hashFunction(key)
	iterator := h.table[index]
	prev := record{-1, -1, nil, true}
	for !iterator.isnull {
		if iterator.key == key {
			if prev.isnull == true {
				h.table[index] = *iterator.next
			} else {
				prev.next = iterator.next
				prev.isnull = iterator.isnull
			}
			return true
		} else {
			prev = iterator
			iterator = *iterator.next
		}
	}
	return false
	//h.checkLoadFactorAndUpdate()
}

//func (h hashTable) checkLoadFactorAndUpdate() {
//	loadFactor := h.loadFactor()
//	if loadFactor < minLoadFactor {
//		// Reduce the size
//	}
//
//	if loadFactor > maxLoadFactor {
//		// Increase the size
//	}
//}

//func (h hashTable) loadFactor() (int) {
//	return h.nRecords/len(h.table)
//}

func main() {
	hashtable := createHashTable()
	hashtable.put(1, 2)
	fmt.Println(hashtable.get(1))
	hashtable.del(1)
}