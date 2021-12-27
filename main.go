package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/higordiego/learning-golang-scanport/model"
)

func main() {
	fmt.Println(model.ShowBanner())
	host := flag.String("h", "", "Domains and ip for scan. (Required)")
	limit := flag.String("l", "int", "Limit port scan. (Required)")
	flag.Parse()

	if *host == "" || *limit == "" {

		model.Helpers()
		flag.PrintDefaults()
		model.Signature()
		return
	}

	i, err := strconv.Atoi(*limit)

	if err != nil {
		model.Helpers()
		flag.PrintDefaults()
		model.Signature()
		return
	}

	model.GetOpenPorts(*host, model.Range{Start: 0, End: i})

	model.Signature()
}
