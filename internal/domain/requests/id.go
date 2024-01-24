package requests

import (
	"strconv"
	"strings"
)

func ExtractIDFromPath(path string) (int, error) {
	parts := strings.Split(path, "/")

	idStr := parts[len(parts)-1]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	return id, nil
}
