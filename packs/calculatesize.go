package packs

import (
	"log"
	"sort"
)

func CalculatePackSizes(packs []int, target int) map[int]int {
	// first if the exact size exists send that
	if findPack(packs, target) {
		return map[int]int{
			target: 1,
		}
	}

	packsToSend := make(map[int]int)

	// sort desc
	sort.Slice(packs, func(i, j int) bool {
		return packs[i] > packs[j]
	})

outer:
	for getCurrentTotal(packsToSend) < target {
		for i, p := range packs {
			ct := getCurrentTotal(packsToSend)

			if p+ct <= target {
				packsToSend[p] += 1
				continue outer
			}

			// lowest number in slice
			if i+1 == len(packs) {
				packsToSend[p] += 1
			}
		}
	}

	for {
		updated := false
		for k, v := range packsToSend {
			total := v * k
			log.Printf("current total: %v", total)
			if v > 1 && findPack(packs, total) {
				log.Println("found better pack")
				delete(packsToSend, k)
				packsToSend[total] += 1
				updated = true
			}
		}

		if !updated {
			break
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

func findPack(packs []int, size int) bool {
	for _, item := range packs {
		if item == size {
			return true
		}
	}
	return false
}
