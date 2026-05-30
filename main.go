package main

import "fmt"

type Node struct {
	key   string
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	items    map[string]*Node
	head     *Node
	tail     *Node
}

func NewLRUCache(capacity int) *LRUCache {
	cache := &LRUCache{
		capacity: capacity,
		items:    make(map[string]*Node),
		head:     &Node{},
		tail:     &Node{},
	}

	cache.head.next = cache.tail
	cache.tail.prev = cache.head

	return cache
}

func (c *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next = c.tail
}

func (c *LRUCache) addToTail(node *Node) {
	node.prev = c.tail.prev
	node.next = c.tail

	c.tail.prev.next = node
	c.tail.prev = node
}

func (c *LRUCache) Get(key string) int {
	node, exists := c.items[key]

	if !exists {
		return -1
	}

	c.remove(node)
	c.addToTail(node)

	return node.value
}

func (c *LRUCache) Put(key string, value int) {
	if node, exists := c.items[key]; exists {
		node.value = value
		c.remove(node)
		c.addToTail(node)
		return
	}

	if len(c.items) >= c.capacity {
		oldest := c.head.next
		c.remove(oldest)
		delete(c.items, oldest.key)
	}

	newNode := &Node{key: key, value: value}
	c.addToTail(newNode)
	c.items[key] = newNode
}

func main() {
	// Создаем кеш с максимальной емкостью в 2 элемента
	cache := NewLRUCache(2)

	cache.Put("A", 1) // Кеш: [A]
	cache.Put("B", 2) // Кеш: [A, B]

	fmt.Println("Get A:", cache.Get("A")) // Выведет 1. Кеш: [B, A] (А стал самым свежим)

	cache.Put("C", 3) // Превышен лимит! Выкидываем "B" (он самый старый). Кеш: [A, C]

	fmt.Println("Get B:", cache.Get("B")) // Выведет -1 (не найдено)
	fmt.Println("Get A:", cache.Get("A")) // Выведет 1. Кеш: [C, A]

	cache.Put("D", 4) // Превышен лимит! Выкидываем "C". Кеш: [A, D]

	fmt.Println("Get C:", cache.Get("C")) // Выведет -1 (не найдено)
	fmt.Println("Get D:", cache.Get("D")) // Выведет 4
}
