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

func NewVMap() VMap {
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

func (c *VMap) Range(from int, to int, call func(v interface{})) {
	c.v.Range(types.Int32(from), types.Int32(to), func(v interface{}) {
		call(v)
	})
}
func (c *VMap) RangeV(fromV []int, toV []int, call func(v interface{})) {
	from := fromV[0]
	to := toV[0]
	rcall := func(v interface{}) {
		nc := v.(*VMap)
		nc.RangeV(fromV[1:], toV[1:], call)
	}
	if len(fromV) == 1 {
		rcall = call
	}
	c.Range(from, to, rcall)
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
