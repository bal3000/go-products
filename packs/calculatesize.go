package packs

import (
	"log"
	"sort"
)

func CalculatePackSizes(packs []int, target int) map[int]int {
	packsToSend := make(map[int]int)

	// sort desc
	sort.Slice(packs, func(i, j int) bool {
		return packs[i] > packs[j]
	})

outer:
	for {
		for i, p := range packs {
			ct := getCurrentTotal(packsToSend)
			log.Printf("current pack: %v", p)
			log.Printf("current total: %v", ct)

			if ct >= target {
				break outer
			}

			if p+ct <= target {
				packsToSend[p] += 1
				continue outer
			}

			// lowest number in slice
			if i+1 == len(packs) {
				packsToSend[p] += 1
				continue outer
			}
		}
	}

	return packsToSend
}

func getCurrentTotal(packs map[int]int) int {
	ct := 0
	for k, v := range packs {
		ct += k * v
	}
	return ct
}
