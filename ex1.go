package main

import (
	"log"
	"sort"

	"github.com/samber/lo"
)

func GetLightList(lightBrightness []int, expectedBrightness int) [][]int {
	// check the list of lights that can provide minimumb rightness
	var (
		validLightBrightness []int
	)
	for _, item := range lightBrightness {
		if item <= expectedBrightness && !lo.Contains(validLightBrightness, item) {
			validLightBrightness = append(validLightBrightness, item)
		}
	}
	if len(validLightBrightness) == 0 {
		return [][]int{}
	}
	// sort the list
	sort.Slice(validLightBrightness, func(i, j int) bool {
		return validLightBrightness[i] < validLightBrightness[j]
	})
	return calculateLights(validLightBrightness, expectedBrightness)
}

func calculateLights(validLightBrightness []int, expectedBrightness int) [][]int {
	var res [][]int
	lightCases(validLightBrightness, expectedBrightness, []int{}, &res)
	return res
}

func lightCases(validLightBrightness []int, expectedBrightness int, lights []int, res *[][]int) {
	// add 1 case to result
	if expectedBrightness == 0 {
		*res = append(*res, lights)
	}
	for idx, lightKindValue := range validLightBrightness {
		// skip if current expected brightness < light kind
		if expectedBrightness < lightKindValue {
			continue
		}
		currentLights := append(lights, lightKindValue)
		lightCases(validLightBrightness[idx:], expectedBrightness-lightKindValue, currentLights, res)
	}
}

func main() {
	log.Println(GetLightList([]int{200, 300, 500}, 800))
	log.Println(GetLightList([]int{200, 300, 600, 700}, 800))
	log.Println(GetLightList([]int{200}, 100))
}
