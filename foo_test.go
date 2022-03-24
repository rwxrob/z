package foo

import (
	"bytes"
	"log"
	"os"
	"testing"
)

// Unlike other Go projects, Bonzai commands don't really benefit from
// Go's example-based tests (which normally would be in package
// foo_test). Instead, testing should be against the first-class Call
// functions directly (with nil callers) or the high-level functions
// from a package library into which they call themselves (which might
// have their own example and other tests).

func TestBar(t *testing.T) {

	// capture the output
	buf := new(bytes.Buffer)
	log.SetFlags(0)
	log.SetOutput(buf)
	defer log.SetFlags(log.Flags())
	defer log.SetOutput(os.Stderr)

	Bar.Call(nil)

	t.Log(buf)
	if buf.String() != "would bar stuff\n" {
		t.Fail()
	}
}
