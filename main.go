package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}

}

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if nil != err {
		return fmt.Errorf("could not initialize SDL: %v", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("could not inititlize TTF: %v", err)
	}
	defer ttf.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if nil != err {
		return fmt.Errorf("could not create window: %v", err)
	}

	defer w.Destroy()

	if err := drawTitle(r, "Flappy Gopher"); err != nil {
		return fmt.Errorf("could not draw title: %v", err)
	}

	// temporary hack: window does not appear without this two calls to PumpEvents
	// this hack becomes redundant when sdl.WaitEvent() is introduced later in the tutorial
	// sdl.PumpEvents()
	// sdl.PumpEvents()

	time.Sleep(1 * time.Second)

	s, err := newScene(r)

	if err != nil {
		return fmt.Errorf("could not create scene: %v", err)
	}

	defer s.destroy()

	events := make(chan sdl.Event)
	errc := s.run(events, r)

	runtime.LockOSThread()
	for {
		select {
		case events <- sdl.WaitEvent():
		case err := <-errc:
			fmt.Printf("error found %v", err)
			return err
		}
	}

}

func drawTitle(r *sdl.Renderer, text string) error {
	r.Clear()

	f, err := ttf.OpenFont("res/fonts/Flappy.ttf", 20)
	if err != nil {
		return fmt.Errorf("could not load font: %v", err)
	}
	defer f.Close()

	c := sdl.Color{R: 255, G: 100, B: 0, A: 255}
	s, err := f.RenderUTF8Solid(text, c)
	if err != nil {
		return fmt.Errorf("could not render title: %v", err)
	}
	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("could not create texture: %v", err)
	}
	defer t.Destroy()

	if err := r.Copy(t, nil, nil); err != nil {
		return fmt.Errorf("could not copy texture %v", err)
	}

	r.Present()

	return nil
}
