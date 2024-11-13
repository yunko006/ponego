package ponenotes_test

import (
	"testing"
	"testing/fstest"

	ponenotes "github.com/yunko006/ponego/internal/ponenotes"
)

func TestNewPoneNotes(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts := ponenotes.NewPoneNotesFromObsidian(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
