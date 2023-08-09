package lru

import "container/list"

type Cache struct {
	maxBytes  int64
	currBytes int64
	cache     map[string]*list.Element
	ll        *list.List
	OnExited  func(key string, val Val)
}

type entry struct {
	key string
	val Val
}

type Val interface {
	Len() int
}

func (c *Cache) Len() int {
	return len(c.cache)
}

func New(maxBytes int64, onExited func(key string, val Val)) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		cache:    make(map[string]*list.Element, 0),
		ll:       list.New(),
		OnExited: onExited,
	}

}

//添加或修改
func (c *Cache) Add(key string, v Val) {
	//key存在，修改value
	if ele := c.cache[key]; ele != nil {
		c.ll.MoveToFront(ele)
		enkv := ele.Value.(*entry)
		c.currBytes += int64(v.Len() - enkv.val.Len())
		ele.Value = v
	} else {
		//key不存在，存入键值对
		el := c.ll.PushFront(&entry{key, v})
		c.cache[key] = el
		c.currBytes += int64(len(key) + v.Len())
	}
	if c.maxBytes != 0 && c.maxBytes < c.currBytes {
		c.DeleteOldest()
	}
}

func (c *Cache) DeleteOldest() {
	ele := c.ll.Back()
	if ele != nil {
		enkv := ele.Value.(*entry)
		c.ll.Remove(ele)
		delete(c.cache, enkv.key)
		c.currBytes -= int64(len(enkv.key) + enkv.val.Len())
		if c.OnExited != nil {
			c.OnExited(enkv.key, enkv.val)
		}
	}
}

//查找
func (c *Cache) Get(key string) (v Val, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		enkv := ele.Value.(*entry)
		return enkv.val, true
	}
	return
}
