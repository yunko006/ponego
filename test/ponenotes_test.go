package ponenotes_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	ponenotes "github.com/yunko006/ponego/internal/ponenotes"
)

func TestNewPoneNotes(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go`
		secondBody = `Title: Post 2 
Description: Description 2 
Tags: rust, borrow-checker`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}
	notes, err := ponenotes.NewPoneNotesFromObsidian(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(notes) != len(fs) {
		t.Errorf("got %d notes, wanted %d notes", len(notes), len(fs))
	}

	got := notes[0]
	want := ponenotes.PoneNote{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

	assertPoneNote(t, notes[0], ponenotes.PoneNote{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
	})
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func assertPoneNote(t *testing.T, got ponenotes.PoneNote, want ponenotes.PoneNote) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
