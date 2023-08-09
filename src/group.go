package src

import "sync"

type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

var (
	mu       = sync.RWMutex{}
	groupmap = map[string]*Group{}
)

func NewGroup(name string, getter Getter) *Group {
	if getter == nil {
		panic("getter is nil")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:   name,
		getter: getter,
	}

	groupmap[name] = g
	return g

}

func GetGroup(name string) *Group {

	mu.RLock()
	defer mu.RUnlock()
	gr := groupmap[name]
	return gr
}
