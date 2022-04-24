package ArrayLib

import (
	"math/rand"
	"time"
)

type ArrayINTF interface {
	InSlice(val interface{}, slice []interface{}) bool
	Shuffle(s []interface{})
	Unique(s []interface{}) []interface{}
	Udiff(a []interface{}, b []interface{}) []interface{}
}
type Array struct {
}

var ArrayObj ArrayINTF = (*Array)(nil)

//InSlice checkout val in slice
func (ArrayObj *Array) InSlice(val interface{}, slice []interface{}) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// Shuffle
func (ArrayObj *Array) Shuffle(s []interface{}) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//Unique
func (ArrayObj *Array) Unique(s []interface{}) []interface{} {
	if len(s) <= 0 {
		return s
	}
	checkMap := map[interface{}]struct{}{}
	for _, id := range s {
		checkMap[id] = struct{}{}
	}
	var ret []interface{}
	for id, _ := range checkMap {
		if id != 0 {
			ret = append(ret, id)
		}
	}
	return ret
}

// Udiff returns the elements in `a` that aren't in `b`.
func (ArrayObj *Array) Udiff(a []interface{}, b []interface{}) []interface{} {
	var c []interface{}
	temp := map[interface{}]struct{}{}
	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}
	for _, val := range a {
		if _, ok := temp[val]; !ok {
			c = append(c, val)
		}
	}
	return c
}
