package main

import (
	"examination2304/model"
	"log"
)

func main() {
	log.Println(model.GetLightList([]int{200, 300, 500}, 800))
	log.Println(model.GetLightList([]int{200, 300, 600, 700}, 800))
	log.Println(model.GetLightList([]int{200}, 100))
}
