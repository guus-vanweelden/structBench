package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("vim-go")
}

type saverLoader interface {
	Save(*item)
	Load(key) *item
}

type item struct {
	ID     string
	Dump   string
	Number int64
	Floati float64
	Flag   bool
}

func (i *item) Key() key {
	return key{
		ID:   i.ID,
		Flag: i.Flag,
	}
}

type key struct {
	ID   string
	Flag bool
}

type mapStringBool map[string]map[bool]*item

func (m mapStringBool) Save(i *item) {
	if _, ok := m[i.ID]; !ok {
		m[i.ID] = make(map[bool]*item)
	}
	m[i.ID][i.Flag] = i
}
func (m mapStringBool) Load(k key) *item {
	return m[k.ID][k.Flag]
}

type mapBoolString map[bool]map[string]*item

func (m mapBoolString) Save(i *item) {
	if _, ok := m[i.Flag]; !ok {
		m[i.Flag] = make(map[string]*item)
	}
	m[i.Flag][i.ID] = i

}
func (m mapBoolString) Load(k key) *item {
	return m[k.Flag][k.ID]
}

type mapStringSlice map[string][2]*item

func (m mapStringSlice) Save(i *item) {
	s := m[i.ID]
	s[boolToInt(i.Flag)] = i
	m[i.ID] = s
}
func (m mapStringSlice) Load(k key) *item {
	return m[k.ID][boolToInt(k.Flag)]
}

type sliceMapString [2]map[string]*item

func (m sliceMapString) Save(i *item) {
	s := m[boolToInt(i.Flag)]
	s[i.ID] = i
	m[boolToInt(i.Flag)] = s
}
func (m sliceMapString) Load(k key) *item {
	return m[boolToInt(k.Flag)][k.ID]
}

type mapStruct map[key]*item

func (m mapStruct) Save(i *item) {
	m[i.Key()] = i
}
func (m mapStruct) Load(k key) *item {
	return m[k]
}

func generateBool(n int64) bool {
	return n%2 == 0
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func generateItems(num int) []*item {
	rand.Seed(time.Now().UnixNano())

	result := make([]*item, num)

	for i := 0; i < num; i++ {
		result[i] = &item{
			ID:     generateString(20),
			Dump:   generateString(rand.Intn(80)),
			Number: rand.Int63(),
			Floati: rand.Float64(),
			Flag:   generateBool(rand.Int63()),
		}
	}

	return result
}
