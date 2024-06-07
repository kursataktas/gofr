package main

import (
	"fmt"
	"time"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/cmd/terminal"
)

func main() {
	app := gofr.NewCMD()

	app.SubCommand("hello", func(ctx *gofr.Context) (interface{}, error) {
		return "Hello World!", nil
	})

	app.SubCommand("params", func(ctx *gofr.Context) (interface{}, error) {
		return fmt.Sprintf("Hello %s!", ctx.Param("name")), nil
	})

	app.SubCommand("stream", func(ctx *gofr.Context) (interface{}, error) {
		ch := make(chan string)

		go func() {
			ch <- "starting something"
			time.Sleep(1 * time.Second)
			ch <- "done task 1"
			ch <- "starting task 2"
			time.Sleep(2 * time.Second)
			ch <- "completed all tasks"

			close(ch)
		}()

		return ch, nil
	})

	app.SubCommand("progress", func(ctx *gofr.Context) (interface{}, error) {
		// Intialize a new TUI instance with os.Stdout as output medium
		o := terminal.NewOutput()
		pBar := terminal.NewProgressBar(o, 100)

		// Starting a process
		o.Println("Starting a time consuming process")
		for pBar.Incr(10) {
			time.Sleep(1 * time.Second)
		}

		return "Completed!", nil
	})

	app.SubCommand("spinner", func(ctx *gofr.Context) (interface{}, error) {
		o := terminal.NewOutput()
		s := terminal.NewGlobeSpinner(o)

		// Start the spinner
		s.Spin()
		// doing some background task
		time.Sleep(3 * time.Second)

		// Stop the spinner to denote completing of a process
		s.Stop()

		return "Completed!", nil
	})

	app.Run()
}
