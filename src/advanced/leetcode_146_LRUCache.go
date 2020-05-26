package advanced

type Data struct {
	Key   int
	Value int
	Left  *Data
	Right *Data
}

type LRUCache struct {
	Data     map[int]*Data
	Head     *Data
	Tail     *Data
	Capacity int
}

func Constructor(capacity int) LRUCache {
	var cache LRUCache
	cache.Data = make(map[int]*Data)
	cache.Capacity = capacity
	return cache
}

func (this *LRUCache) Get(key int) int {
	var date *Data
	if value, ok := this.Data[key]; !ok {
		return -1
	} else {
		date = value
	}
	this.Head.Left = this.Tail
	this.Tail = this.Head
	this.Head = this.Head.Right
	this.Tail.Right = nil
	this.Head.Left = nil
	return date.Value
}

func (this *LRUCache) Put(key int, value int) {

	data := Data{Key:key, Value: value, Left: this.Tail, Right: nil}
	if _, ok := this.Data[key]; !ok {
		if len(this.Data) < this.Capacity {
			this.Capacity++
		} else {
			this.Head = this.Head.Right
			delete(this.Data, this.Head.Left.Key)
			this.Head.Left = nil
		}
	}
	this.Data[key] = &data
	this.Tail.Right = &data
	this.Tail = this.Tail.Right
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
