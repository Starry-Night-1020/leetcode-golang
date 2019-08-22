package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	set map[int]bool
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	var randomSet RandomizedSet = RandomizedSet{set: make(map[int]bool)}
	return randomSet
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if this.set[val] {
		return false
	}
	this.set[val] = true
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if this.set[val] {
		delete(this.set, val)
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	length := len(this.set)
	keys := make([]int, 0, length)
	for k, _ := range this.set {
		keys = append(keys, k)
	}
	return keys[rand.Intn(length)]
}

func main() {
	/**
	 * Your RandomizedSet object will be instantiated and called as such:
	 * obj := Constructor();
	 * param_1 := obj.Insert(val);
	 * param_2 := obj.Remove(val);
	 * param_3 := obj.GetRandom();
	 */
	obj := Constructor()
	param1 := obj.Insert(1)
	fmt.Println("param1: ", param1)
	param2 := obj.Insert(2)
	fmt.Println("param2: ", param2)
	param3 := obj.Remove(2)
	fmt.Println("param3: ", param3)
	fmt.Println("set", obj.set)
	param4 := obj.GetRandom()
	fmt.Println("param4: ", param4)
}
