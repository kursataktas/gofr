package cmd

import (
	"fmt"
	"os"
)

type Responder struct{}

func (r *Responder) Respond(data interface{}, err error) {
	// TODO - provide proper exit codes here. Using os.Exit directly is a problem for tests.
	if data != nil {
		c, ok := data.(chan string)
		if ok {
			for {
				v, okC := <-c
				if !okC {
					break
				}

				fmt.Fprint(os.Stdout, v, "\n")
			}
		} else {
			fmt.Fprint(os.Stdout, data, "\n")
		}
	}

	if err != nil {
		fmt.Fprint(os.Stderr, err, "\n")
	}
}
