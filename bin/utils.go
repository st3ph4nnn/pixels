package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func get_random(min, max int) int32 {
	return int32(rand.Intn(max-min+1) + min)
}

func get_random_float(min, max float32) float32 {
	return float32(float32(rand.Intn(int(max-min+1))) + min)
}

func get_random_color() rl.Color {
	col := get_random(1, 10)

	switch col {
	case 1:
		return rl.Black
	case 2:
		return rl.Red
	case 3:
		return rl.Blue
	case 4:
		return rl.Magenta
	case 5:
		return rl.DarkGreen
	case 6:
		return rl.Purple
	case 7:
		return rl.Yellow
	case 8:
		return rl.Lime
	case 9:
		return rl.DarkBlue
	case 10:
		return rl.Orange
	}

	return rl.Black
}
