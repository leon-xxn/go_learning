package src

import "testing"

//least recently used
//缓存淘汰算法,历史访问记录淘汰数据

func TestLru(t *testing.T) {

}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}
