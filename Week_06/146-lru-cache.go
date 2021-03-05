package Week_06

type LRUCache struct {
	size int
	cap int
	hashMap map[int]*doubleLikeNode
	head *doubleLikeNode
	tail *doubleLikeNode
}

type doubleLikeNode struct {
	key int
	value int
	pre *doubleLikeNode
	next *doubleLikeNode
}

func Constructor(capacity int) LRUCache {
	t := LRUCache{
		size: 0,
		cap: capacity,
		head: &doubleLikeNode{},
		tail: &doubleLikeNode{},
		hashMap: map[int]*doubleLikeNode{},
	}

	t.head.next = t.tail
	t.tail.pre = t.head

	return t
}


func (this *LRUCache) Get(key int) int {
	if _, has := this.hashMap[key]; has == false {
		return -1
	}

	node := this.takeNodeByKey(key)
	this.putToHead(node)

	return node.value
}


func (this *LRUCache) Put(key int, value int)  {
	var node *doubleLikeNode

	if _, has := this.hashMap[key]; has {
		node = this.takeNodeByKey(key)
		node.value = value
	} else {
		node = &doubleLikeNode{
			key: key,
			value: value,
			pre: &doubleLikeNode{},
			next: &doubleLikeNode{},
		}
	}

	this.putToHead(node)
}

// 取出指定节点
func (this *LRUCache) takeNodeByKey(key int) *doubleLikeNode {
	node := this.hashMap[key]
	node.pre.next = node.next
	node.next.pre = node.pre
	this.size--
	delete(this.hashMap, key)

	return node
}

// 将一个节点放到开始
func (this *LRUCache) putToHead(node *doubleLikeNode) {
	this.removeLastNodeIfTooLarge()
	tmp := this.head.next
	this.head.next = node
	node.pre = this.head
	node.next = tmp
	tmp.pre = node

	this.size++
	this.hashMap[node.key] = node
}

// 检查缓存，如果缓存过大就删除最后一个元素
func (this *LRUCache) removeLastNodeIfTooLarge() {
	if this.size < this.cap {
		return
	}

	for this.size >= this.cap {
		tmp := this.tail.pre
		tmp.pre.next = this.tail
		this.tail.pre = tmp.pre
		delete(this.hashMap, tmp.key)
		this.size--
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
