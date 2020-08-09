package main

import (
	"encoding/json"
	"io"
	"fmt"
)
type FileSystemPlayerStore struct {
	database io.Reader
}

type ReadSeeker interface {
	Reader
	Seeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	league, _ := NewLeague(f.database)
	return league
}

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsin league, %v", err)

	}
	return league, err
}