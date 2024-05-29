package cmd

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"gofr.dev/pkg/gofr/testutil"
)

func TestResponder_Respond(t *testing.T) {
	r := Responder{}

	out := testutil.StdoutOutputForFunc(func() {
		r.Respond("data", nil)
	})

	err := testutil.StderrOutputForFunc(func() {
		r.Respond(nil, errors.New("error")) //nolint:goerr113 // We are testing if a dynamic error would work.
	})

	assert.Equal(t, "data\n", out, "TEST Failed.\n", "Responder stdout output")

	assert.Equal(t, "error\n", err, "TEST Failed.\n", "Responder stderr output")
}

func TestResponder_ChannelResponse(t *testing.T) {
	r := Responder{}

	out := testutil.StdoutOutputForFunc(func() {
		ch := make(chan string)

		go func() {
			for i := 0; i < 2; i++ {
				ch <- fmt.Sprintf("string %d ", i)
			}

			close(ch)
		}()

		r.Respond(ch, nil)

	})

	assert.Equal(t, "string 0 \nstring 1 \n", out, "TEST Failed.\n", "Responder stdout output")
}
