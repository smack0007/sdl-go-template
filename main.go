package main

import (
	"math"
	"os"

	SDL "github.com/smack0007/sdl-go/sdl"
)

const WINDOW_WIDTH = 1024
const WINDOW_HEIGHT = 768
const WINDOW_TITLE = "Snake"
const DESIRED_FPS = 60

var GAME_TICK_RATE = (uint64)(math.Floor(float64(1000) / float64(DESIRED_FPS)))

func main() {
	os.Exit(run())
}

func run() int {
	if SDL.Init(SDL.INIT_VIDEO) != 0 {
		SDL.LogError(SDL.LOG_CATEGORY_APPLICATION, "Failed initialize SDL.")
		return 1
	}
	defer SDL.Quit()

	SDL.LogSetPriority(SDL.LOG_CATEGORY_APPLICATION, SDL.LOG_PRIORITY_DEBUG)

	var window *SDL.Window
	var renderer *SDL.Renderer
	result := SDL.CreateWindowAndRenderer(WINDOW_WIDTH, WINDOW_HEIGHT, SDL.WINDOW_SHOWN, &window, &renderer)

	if result != 0 {
		SDL.LogError(SDL.LOG_CATEGORY_APPLICATION, "Failed to create window and renderer.")
		return 1
	}

	defer SDL.DestroyWindow(window)
	defer SDL.DestroyRenderer(renderer)

	SDL.SetWindowTitle(window, WINDOW_TITLE)

	shouldQuit := false
	event := SDL.Event{}

	currentTime := SDL.GetTicks64()
	lastTime := currentTime

	for !shouldQuit {
		for SDL.PollEvent(&event) > 0 {
			switch event.Type() {

			case SDL.QUIT:
				shouldQuit = true
			}
		}

		currentTime = SDL.GetTicks64()
		elapsedTime := currentTime - lastTime

		if elapsedTime >= GAME_TICK_RATE {
			update(float32(elapsedTime) / float32(1000))
			draw(renderer)

			lastTime = currentTime
		}

		SDL.Delay(1)
	}

	return 0
}

func update(elapsedTime float32) {
}

func draw(renderer *SDL.Renderer) {
	SDL.SetRenderDrawColor(renderer, 100, 149, 237, 255)
	SDL.RenderClear(renderer)

	SDL.RenderPresent(renderer)
}
