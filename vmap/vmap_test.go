package vmap

import (
	"fmt"
	"testing"
)

func TestVM(t *testing.T) {
	vm := New()
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			for z := 0; z < 10; z++ {
				v := []int{x, y, z}
				vm.SetV(v, v)
			}
		}
	}
	vm.RangeV([]int{5, 0, 5}, []int{7, 10000000, 6}, func(key []int, v interface{}) {
		fmt.Printf("%v %v\n", key, v)
	})
	fmt.Printf("====================\n")
	vm.DeleteV([]int{5, 0, 5})
	size := vm.SizeRangeV([]int{0, 0, 0}, []int{10, 10, 10})
	if size != 999 {
		t.Fatalf("delete error %d", size)
	}
	removeKeys := [][]int{}
	vm.RangeV([]int{0, 0, 0}, []int{10000, 10000, 10000}, func(key []int, v interface{}) {
		if key[0]%2 == 0 {
			fmt.Printf("%v %v\n", key, v)
			removeKeys = append(removeKeys, key)
		}
	})
	for _, key := range removeKeys {
		vm.DeleteV(key)
	}

	res := []interface{}{}
	vm.RangeV([]int{0, 0, 0}, []int{10000, 10000, 10000}, func(key []int, v interface{}) {
		fmt.Printf("%v %v\n", key, v)
		res = append(res, v)
	})
	if len(res) != 499 {
		t.Fatalf("delete error %d", len(res))
	}

	vm.DeleteRangeV([]int{0, 0, 0}, []int{1000, 1000, 1000})
	size = vm.SizeRangeV([]int{0, 0, 0}, []int{1000, 1000, 1000})
	if size != 0 {
		t.Fatalf("delete error %d", size)
	}

}
