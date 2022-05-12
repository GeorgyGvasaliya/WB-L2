package main

import "L2/pattern/strategy/pkg"

func main() {
	lfu := &pkg.Lfu{}
	cache := pkg.InitCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")

	cache.Add("c", "3")

	lru := &pkg.Lru{}
	cache.SetEvictionAlgo(lru)

	cache.Add("d", "4")

	fifo := &pkg.Fifo{}
	cache.SetEvictionAlgo(fifo)

	cache.Add("e", "5")

}
