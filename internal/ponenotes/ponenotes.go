package ponenotes

import (
	"io/fs"
)

func NewPoneNotesFromObsidian(fileSystem fs.FS) ([]PoneNote, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var ponenotes []PoneNote
	for _, f := range dir {
		note, err := getPoneNote(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		ponenotes = append(ponenotes, note)
	}
	return ponenotes, nil
}

func getPoneNote(fileSystem fs.FS, fileName string) (PoneNote, error) {
	ponenoteFile, err := fileSystem.Open(fileName)
	if err != nil {
		return PoneNote{}, err
	}
	defer ponenoteFile.Close()
	return newPoneNote(ponenoteFile)
}
