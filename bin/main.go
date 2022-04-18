package main

import (
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var pixels []pixel
var color rl.Color = rl.Black
var close = false

var stop = false
var key int32 = 0
var input bool = false

var info bool = true

func modify(width, height int32) {
	var delta_frame_time float32 = 1.0
	for !close {
		start := time.Now()

		if input {
			switch {
			case key == 1:
				pixels = nil
			case key == 2 || key == 3:
				pos := rl.GetMousePosition()

				if len(pixels) == 0 {
					pixels = append(pixels, pixel{x: float32(pos.X), y: float32(pos.Y), id: 1, col: color, alpha: 1.0})
				} else {
					pixels = append(pixels, pixel{x: float32(pos.X), y: float32(pos.Y), id: pixels[len(pixels)-1].id + 1, col: color, alpha: 1.0})
				}
			case key == 4:
				color = get_random_color()
			case key == 5:
				stop = !stop
			}
			input = false
		}

		if stop {
			continue
		}

		for i := range pixels {
			if stop {
				continue
			}

			if pixels[i].collides() {
				pixels = append(pixels[:i], pixels[i+1:]...)
				break
			}

			if pixels[i].get_nearest() {
				if get_random(0, 1) == 1 {
					if get_random(0, 1) == 1 {
						if get_random(0, 1) == 1 {
							pixels[i].change_x(delta_frame_time * 100)
						} else {
							pixels[i].change_x(-(delta_frame_time * 100))
						}
					} else {
						if get_random(0, 1) == 1 {
							pixels[i].change_y(delta_frame_time * 100)
						} else {
							pixels[i].change_y(-(delta_frame_time * 100))
						}
					}
				} else {
					pixels[i].move_to(pixels[i].nearest, delta_frame_time)
				}
			} else {
				if get_random(0, 1) == 1 {
					if get_random(0, 1) == 1 {
						pixels[i].change_x(delta_frame_time * 100)
					} else {
						pixels[i].change_x(-(delta_frame_time * 100))
					}
				} else {
					if get_random(0, 1) == 1 {
						pixels[i].change_y(delta_frame_time * 100)
					} else {
						pixels[i].change_y(-(delta_frame_time * 100))
					}
				}
			}

			go pixels[i].clamp(0, float32(width-10), 0, float32(height-10))
		}
		delta_frame_time = float32(time.Since(start).Seconds())
	}
}

func main() {
	rl.SetConfigFlags(rl.FlagWindowAlwaysRun | rl.FlagMsaa4xHint | rl.FlagWindowMaximized)
	rl.InitWindow(0, 0, "pixels")
	rl.ToggleFullscreen()
	rl.SetTargetFPS(244)

	width := int32(rl.GetScreenWidth())
	height := int32(rl.GetScreenHeight())
	rand.Seed(time.Now().UnixNano())

	go modify(width, height)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		close = false

		rl.ClearBackground(rl.White)

		input = rl.GetKeyPressed() != 0 || rl.IsMouseButtonDown(rl.MouseRightButton) || rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.GetMouseWheelMove() != 0

		if input {
			switch {
			case rl.IsKeyPressed(rl.KeyDelete):
				key = 1
			case rl.IsMouseButtonDown(rl.MouseRightButton):
				key = 2
			case rl.IsMouseButtonPressed(rl.MouseLeftButton):
				key = 3
			case rl.GetMouseWheelMove() != 0:
				key = 4
			case rl.IsKeyPressed(rl.KeyEnter):
				key = 5
			case rl.IsKeyPressed(rl.KeyX):
				info = !info
			case rl.IsKeyPressed(rl.KeyF11):
				rl.ToggleFullscreen()
			default:
				input = false
				key = 0
			}
		}

		for _, pixel := range pixels {
			rl.DrawRectangle(int32(pixel.x), int32(pixel.y), 15, 15, rl.ColorAlpha(pixel.col, pixel.alpha))
			rl.DrawRectangleLines(int32(pixel.x), int32(pixel.y), 15, 15, rl.ColorAlpha(rl.Black, pixel.alpha))
		}

		rl.DrawText(strconv.FormatInt(int64(len(pixels)), 10), 30, 60, 50, rl.SkyBlue)
		rl.DrawRectangle(30, 120, 50, 50, color)
		rl.DrawRectangleLines(29, 119, 51, 51, rl.Black)

		if info {
			rl.DrawRectangle(0, 0, width, height, rl.ColorAlpha(rl.Black, 0.5))

			rl.DrawText("F11 - fullscreen (may cause issues)", width/2-rl.MeasureText("F11 - fullscreen (may cause issues)", 40)/2, height/2-250, 40, rl.Black)
			rl.DrawText("ESC - exit", width/2-rl.MeasureText("ESC - exit", 40)/2, height/2-200, 40, rl.Black)
			rl.DrawText("DELETE - clear screen", width/2-rl.MeasureText("DELETE - clear screen", 40)/2, height/2-150, 40, rl.Black)
			rl.DrawText("LMB - drop 1 pixel", width/2-rl.MeasureText("LMB - drop 1 pixel", 40)/2, height/2-100, 40, rl.Black)
			rl.DrawText("RMB (hold) - drop multiple pixels", width/2-rl.MeasureText("RMB (hold) - drop multiple pixels", 40)/2, height/2-50, 40, rl.Black)
			rl.DrawText("SCROLL - change pixel color", width/2-rl.MeasureText("SCROLL - change pixel color", 40)/2, height/2, 40, rl.Black)
			rl.DrawText("(press X to hide)", width/2-rl.MeasureText("(press X to hide)", 60)/2, height/2+125, 60, rl.Black)

			rl.DrawText("pixels - made by stephan", width/2-int32(rl.MeasureText("pixels - made by stephan", 60))/2, height-100, 60, rl.White)
		}

		rl.DrawFPS(30, 30)
		rl.EndDrawing()
	}

	close = true

	rl.CloseWindow()
}
