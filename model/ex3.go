package model

import (
	"github.com/samber/lo"
)

type Room struct {
	ID     string
	Name   string
	Status bool
	HomeID string
}

type Light struct {
	ID         string
	Name       string
	Color      string
	Brightness int
	RoomID     string
}

// init sample with 2 rooms and 4 lights
var rooms = []*Room{
	{
		ID:     "1",
		Name:   "Living Room",
		Status: true,
		HomeID: "1",
	},
	{
		ID:     "2",
		Name:   "Bed Room",
		Status: true,
		HomeID: "1",
	},
}

var lights = []*Light{
	{
		ID:         "1",
		Name:       "Light 1",
		Color:      "#e41e3f",
		Brightness: 200,
		RoomID:     "1",
	},
	{
		ID:         "2",
		Name:       "Light 2",
		Color:      "#ffb7df",
		Brightness: 300,
		RoomID:     "1",
	},
	{
		ID:         "3",
		Name:       "Light 3",
		Color:      "#e1beff",
		Brightness: 500,
		RoomID:     "1",
	},
	{
		ID:         "4",
		Name:       "Light 4",
		Color:      "#5c3641",
		Brightness: 700,
		RoomID:     "2",
	},
}

func GetLightByRoomID(roomID string) []*Light {
	var lights []*Light
	// get lights by room id
	lights = lo.Filter(lights, func(l *Light, _ int) bool {
		return l.RoomID == roomID
	})
	// order with rainbow color

	return lights
}
