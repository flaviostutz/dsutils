package dsutils

import "time" 

type Element struct {
	item interface{}
	time time.Time
}

type TTLCollection struct {
    Items []Element
    TtlSeconds float32
}

func NewTTLCollection(ttlSeconds0 float32) TTLCollection {
	a := TTLCollection {
		TtlSeconds: ttlSeconds0,
		Items: make([]Element, 0),
	}
	return a
}

func (c *TTLCollection) Add(element interface{}) {
	elem := Element {
		time: time.Now(),
		item: element,
	}
	c.Items = append(c.Items, elem)
}

func (c *TTLCollection) List() []interface{} {
	list := make([]Element, 0)
	result := make([]interface{}, 0)
	for i:=0; i<len(c.Items); i++ {
		item0 := c.Items[i]
		if c.TtlSeconds <=0 || float32(time.Now().Sub(item0.time).Seconds()) <= c.TtlSeconds {
			list = append(list, item0)
			result = append(result, item0.item)
		}
	}
	c.Items = list
	return result
}

