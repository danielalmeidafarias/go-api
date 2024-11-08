package helpers

import (
	albums "github.com/danielalmeidafarias/go-api/db"

	"strconv"
)

func GetLastId() string {
	var id, err = strconv.Atoi(albums.Albums[len(albums.Albums)-1].ID)

	if err != nil {
		return err.Error()
	}

	newId := id + 1

	return strconv.Itoa(newId)
}
