package dynamicprogramming

import (
	"fmt"
	"testing"
)

type KnapsackItem struct {
	Name   string
	Weight int
	Value  int
}

type Key struct {
	w, v int
	list []KnapsackItem
}

func KnapsackBestValue(items []KnapsackItem, idx, weightLimit, currentWeight, currentValue int, mem [][]*Key) Key {

	if currentWeight > weightLimit {
		return Key{0, 0, []KnapsackItem{}}
	}

	if idx == len(items) {
		mem[idx][currentWeight] = &Key{currentWeight, currentValue, []KnapsackItem{}}
		return *mem[idx][currentWeight]
	}

	if k := mem[idx][currentWeight]; k != nil && k.v > currentValue {
		return *k
	}

	k_without := KnapsackBestValue(items, idx+1, weightLimit, currentWeight, currentValue, mem)
	k_within := KnapsackBestValue(items, idx+1, weightLimit, currentWeight+items[idx].Weight, currentValue+items[idx].Value, mem)

	if k_within.v > k_without.v {
		mem[idx][currentWeight] = &Key{k_within.w, k_within.v, append(k_within.list, items[idx])}
	} else {
		mem[idx][currentWeight] = &Key{k_without.w, k_without.v, k_without.list}
	}

	return *mem[idx][currentWeight]
}

func maxKey(a, b Key) Key {
	if a.v > b.v {
		return a
	}
	return b
}

func Test_Knapsack(t *testing.T) {
	fmt.Println("Hello world")
	totalCapacity := 10

	items := []KnapsackItem{
		{"A", 1, 20},
		{"B", 2, 6}, // test swapping value to {"B", 2, 4},
		{"C", 3, 10},
		{"D", 8, 40},
		{"E", 7, 15},
		{"F", 4, 25},
	}

	mem := make([][]*Key, len(items)+1)
	for i := range mem {
		mem[i] = make([]*Key, totalCapacity+1)
	}

	k := KnapsackBestValue(items, 0, totalCapacity, 0, 0, mem)
	fmt.Println(k.w, k.v, k.list)
}
