package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

type pairList []pair

func (p pairList) Len() int {
	return len(p)
}

func (p pairList) Less(i, j int) bool {
	return p[i].count < p[j].count
}

func (p pairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pairList) isExist(name string) int {
	for index, pair := range p {
		if pair.name == name {
			return index
		}
	}
	return -1
}

func MostAppearItem(items []string) []pair {
	var res pairList
	for _, item := range items {
		key := res.isExist(item)
		if 0 <= key {
			res[key].count += 1
		} else {
			res = append(res, pair{name: item, count: 1})
		}
	}
	sort.Sort(pairList(res))
	return res
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// golang->1 ruby->2 js->4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// C->1 D->2 B->3 A->4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
	// football->1 basketball->1 tenis->1
}
