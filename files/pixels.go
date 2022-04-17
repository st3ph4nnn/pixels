package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type pixel struct {
	x       float32
	y       float32
	id      uint32
	col     rl.Color
	alpha   float32
	nearest *pixel
	alive   int32
}

func (p *pixel) clamp(x_min, x_max, y_min, y_max float32) {
	if p.x > x_max {
		p.x = x_max
	}

	if p.x < x_min {
		p.x = x_min
	}

	if p.y > y_max {
		p.y = y_max
	}

	if p.y < y_min {
		p.y = y_min
	}
}

func (p *pixel) change_x(x float32) {
	p.x += x
}

func (p *pixel) change_y(y float32) {
	p.y += y
}

func (p *pixel) collides() bool {
	for i := range pixels {
		if pixels[i].id != p.id && pixels[i].col != p.col && collision_check(pixels[i].x, p.x, pixels[i].y, p.y) {
			return true
		}
	}

	return false
}

func collision_check(x1, x2, y1, y2 float32) bool {
	return x1 <= x2+12.0 && x1+12.0 >= x2 && y1 <= y2+12.0 && 12.0+y1 >= y2
}

func (p *pixel) get_distance(target *pixel) float32 {
	var max_x = math.Max(float64(p.x), float64(target.x))
	var max_y = math.Max(float64(p.y), float64(target.y))

	var min_x = math.Min(float64(p.x), float64(target.x))
	var min_y = math.Min(float64(p.y), float64(target.y))

	var delta_x = max_x - min_x
	var delta_y = max_y - min_y

	return float32(math.Sqrt((delta_x * delta_x) + (delta_y * delta_y)))
}

func (p *pixel) get_nearest() bool {
	var min_distance float32 = 9999999.0

	for i := range pixels {
		if &pixels[i] != p && pixels[i].id != p.id && pixels[i].col != p.col && min_distance > p.get_distance(&pixels[i]) {
			min_distance = p.get_distance(&pixels[i])
			p.nearest = &pixels[i]
		}
	}

	if min_distance == 9999999.0 {
		return false
	}

	return true
}

func (p *pixel) move_to(target *pixel, delta_frame_time float32) {
	if target.col == p.col || target.id == p.id || target == nil {
		return
	}

	var max_x = math.Max(float64(p.x), float64(target.x))
	var max_y = math.Max(float64(p.y), float64(target.y))

	var min_x = math.Min(float64(p.x), float64(target.x))
	var min_y = math.Min(float64(p.y), float64(target.y))

	var delta_x = ((max_x - min_x) / float64(p.get_distance(target)))
	var delta_y = ((max_y - min_y) / float64(p.get_distance(target)))

	if p.x > target.x {
		p.x -= float32(delta_x * float64(delta_frame_time) * 100)
	}

	if p.y > target.y {
		p.y -= float32(delta_y * float64(delta_frame_time) * 100)
	}

	if p.y < target.y {
		p.y += float32(delta_y * float64(delta_frame_time) * 100)
	}

	if p.x < target.x {
		p.x += float32(delta_x * float64(delta_frame_time) * 100)
	}
}
