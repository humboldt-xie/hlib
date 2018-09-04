package vmap

import (
	"github.com/timtadh/data-structures/types"
)

type Vec struct {
	X []int
}

type VMap struct {
	v *AvlTree //map[int]interface{}
}

func New() VMap {
	v := VMap{}
	v.Init()
	return v
}

func (c *VMap) Init() {
	c.v = NewAvlTree() //make(map[int]interface{})
}
func (c *VMap) Get(x int) (v interface{}, ok error) {
	v, ok = c.v.Get(types.Int32(x))
	return
}

/*func (c *VMap) Set(x int, v interface{}) {
	c.v[x] = v
}*/
func (c *VMap) DeleteV(x []int) bool {
	if len(x) == 1 {
		c.v.Remove(types.Int32(x[0]))
		return c.v.Size() == 0
	}
	ci, ok := c.v.Get(types.Int32(x[0]))
	if ok != nil {
		return false
	}
	cc := ci.(*VMap)
	if cc.DeleteV(x[1:]) {
		cc.v.Remove(types.Int32(x[0]))
	}
	return cc.v.Size() == 0
}

func (c *VMap) SetV(x []int, v interface{}) {
	if len(x) == 1 {
		c.v.Put(types.Int32(x[0]), v)
		return
	}
	ci, ok := c.v.Get(types.Int32(x[0]))
	if ok != nil {
		cc := &VMap{}
		cc.Init()
		c.v.Put(types.Int32(x[0]), cc)
		ci = cc
	}
	cc := ci.(*VMap)
	cc.SetV(x[1:], v)
}

func (c *VMap) Range(from int, to int, call func(key int, v interface{})) {
	c.v.Range(types.Int32(from), types.Int32(to), func(key types.Hashable, v interface{}) {
		call(int(key.(types.Int32)), v)
	})
}

func (c *VMap) rangeV(prev []int, fromV []int, toV []int, call func(key []int, v interface{})) {
	from := fromV[0]
	to := toV[0]
	rcall := func(key int, v interface{}) {
		p := append(prev, key)
		if len(fromV) == 1 {
			call(p, v)
		} else {
			nc := v.(*VMap)
			nc.rangeV(p, fromV[1:], toV[1:], call)
		}
	}
	c.Range(from, to, rcall)
}
func (c *VMap) RangeV(fromV []int, toV []int, call func(key []int, v interface{})) {
	c.rangeV([]int{}, fromV, toV, call)
}
func (c *VMap) SizeRangeV(fromV []int, toV []int) int {
	count := 0
	c.RangeV(fromV, toV, func(key []int, v interface{}) {
		count++
	})
	return count
}

func (c *VMap) DeleteRangeV(fromV []int, toV []int) int {
	removeKeys := [][]int{}
	c.RangeV(fromV, toV, func(key []int, v interface{}) {
		removeKeys = append(removeKeys, key)
	})
	for _, key := range removeKeys {
		c.DeleteV(key)
	}
	return len(removeKeys)
}

/*func main() {
	c := VMap{}
	c.Init()
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			for z := 0; z < 10; z++ {
				v := []int{x, y, z}
				c.SetV(v, v)
			}
		}
	}
	c.RangeV([]int{5, 0, 5}, []int{7, 100000000, 6}, func(v interface{}) {
		fmt.Printf("%v\n", v)
	})
}*/
