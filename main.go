package main

import (
	"github.com/higordiego/learning-golang-scanport/model"
)

func main() {

	model.GetOpenPorts("api.reportfy.com.br", model.Range{Start: 0, End: 85})
}
