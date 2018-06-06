package main

import "fmt"

type record struct {
	key int
	value int
	next *record
}

var (
	minLoadFactor = 0.25
	maxLoadFactor = 0.75
	defaultTableSize = 3
)

type hashTable struct {
	table []*record
	nRecords *int
}

func createHashTable(tableSize int) hashTable {
	num := 0
	return hashTable{table: make([]*record, tableSize), nRecords:&num}
}

func CreateHashTable() hashTable {
	num := 0
	return hashTable{table: make([]*record, defaultTableSize), nRecords:&num}
}

func hashFunction(key int, size int) (int) {
	return key%size
}

func (h hashTable) Display() {
	fmt.Printf("----------%d elements-------\n", *h.nRecords)
	for i, node := range h.table {
		fmt.Printf("%d :", i)
		for node != nil {
			fmt.Printf("[%d, %d]->", node.key, node.value)
			node = node.next
		}
		fmt.Println("nil")
	}
}

func (h hashTable) put(key int, value int) (bool){
	index := hashFunction(key, len(h.table))
	iterator := h.table[index]
	node := record{key, value, nil}
	if iterator == nil {
		h.table[index] = &node
	} else {
		prev := &record{0, 0, nil}
		for iterator != nil {
			if iterator.key == key { // Key already exists
				iterator.value = value
				return false
			}
			prev = iterator
			iterator = iterator.next
		}
		prev.next = &node
	}
	*h.nRecords += 1
	return true
}

func (h hashTable) Put(key int, value int) {
	sizeChanged := h.put(key, value)
	if sizeChanged == true {
		h.checkLoadFactorAndUpdate()
	}
}

func (h hashTable) Get(key int) (bool, int) {
	index := hashFunction(key, len(h.table))
	iterator := h.table[index]
	for iterator != nil {
		if iterator.key == key {	// Key already exists
			return true, iterator.value
		}
		iterator = iterator.next
	}
	return false, 0
}

func (h hashTable) del(key int) (bool) {
	index := hashFunction(key, len(h.table))
	iterator := h.table[index]
	if iterator == nil {
		return false
	}
	if iterator.key == key {
		h.table[index] = iterator.next
		*h.nRecords--
		return true
	} else {
		prev := iterator
		iterator = iterator.next
		for iterator != nil {
			if iterator.key == key {
				prev.next = iterator.next
				*h.nRecords--
				return true
			}
			prev = iterator
			iterator = iterator.next
		}
		return false
	}
}

func (h hashTable) Del(key int) (bool) {
	sizeChanged := h.del(key)
	if sizeChanged == true {
		h.checkLoadFactorAndUpdate()
	}
	return sizeChanged
}

func (h hashTable) getLoadFactor() (float64) {
	return float64(*h.nRecords)/float64(len(h.table))
}

func (h hashTable) checkLoadFactorAndUpdate() {
	if *h.nRecords == 0 {
		return
	} else {
		loadFactor := h.getLoadFactor()
		//fmt.Println(loadFactor, minLoadFactor, maxLoadFactor)
		if loadFactor < minLoadFactor || loadFactor > maxLoadFactor {
			hash := createHashTable(*h.nRecords*2)
			for _, record := range h.table {
				for record != nil {
					hash.put(record.key, record.value)
					record = record.next
				}
			}
			fmt.Println("**-----------------------")
			copy(h.table, hash.table)
			hash.Display()
			h.Display()
		}
	}
}

func main(){
	h := CreateHashTable()
	//h.Display()
	h.Put(1,2)
	//h.Display()
	h.Put(2,3)
	//h.Display()
	h.Put(3,4)
	//h.Display()
	h.Put(4,5)
	//h.Display()
	h.Put(5,6)
	//h.Display()
}