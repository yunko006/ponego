package ponenotes

import (
	"bufio"
	"io"
	"strings"
)

type PoneNote struct {
	Title       string
	Description string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

func newPoneNote(noteBody io.Reader) (PoneNote, error) {
	scanner := bufio.NewScanner(noteBody)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)

	}

	return PoneNote{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
	}, nil
}
