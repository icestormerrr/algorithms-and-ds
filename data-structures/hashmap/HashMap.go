package hashmap

import "fmt"

type Pair struct {
	Key   string
	Value int
}

type HashMap struct {
	table []Pair
}

func NewHashMap(size int) *HashMap {
	return &HashMap{table: make([]Pair, size)}
}

func (this *HashMap) PrintTable() {
	fmt.Println(this.table)
}

func (this *HashMap) PrintMap() {
	for _, pair := range this.table {
		fmt.Println(pair.Key, pair.Value)
	}
}

func (this *HashMap) GetLength() int {
	var counter = 0
	for i := range this.table {
		if this.table[i].Key != "" {
			counter++
		}
	}
	return counter
}

func (this *HashMap) Set(key string, value int) {
	if key == "" {
		return
	}

	if len(this.table) == 0 {
		this.increaseTableSize(1)
	}

	var freeIndex = this.findFreeIndexForKey(key)
	if freeIndex != -1 {
		this.table[freeIndex] = Pair{key, value}
	} else {
		this.increaseTableSize(len(this.table) * 2)
		this.Set(key, value)
	}
}

func (this *HashMap) increaseTableSize(size int) {
	var oldMap = this.table
	this.table = make([]Pair, size)
	for _, pair := range oldMap {
		this.Set(pair.Key, pair.Value)
	}
}

func (this *HashMap) getHash(s string) int {
	hash := 0
	prime := 31
	for _, char := range s {
		hash = hash*prime + int(char)
	}
	return hash % len(this.table)
}

func (this *HashMap) findFreeIndexForKey(key string) int {
	var index = this.getHash(key)

	if this.table[index].Key == "" {
		return index
	}

	for index < len(this.table) {
		if this.table[index].Key == key {
			return index
		}

		if this.table[index].Key == "" {
			return index
		}
		index++
	}

	return -1
}

func (this *HashMap) Get(key string) int {
	if key == "" {
		return 0
	}

	if len(this.table) == 0 {
		return 0
	}

	var index = this.getHash(key)

	if this.table[index].Key == key {
		return this.table[index].Value
	}

	for i := index; i < len(this.table); i++ {
		if this.table[i].Key == key {
			return this.table[i].Value
		}
	}

	return 0
}
