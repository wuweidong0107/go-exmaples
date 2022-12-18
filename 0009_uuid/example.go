package main

import (
	"fmt"
	"strings"

	"github.com/pborman/uuid"
)

func main() {
	uuidWithHypen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHypen.String(), "-", "", -1)
	fmt.Println(uuid)
}
