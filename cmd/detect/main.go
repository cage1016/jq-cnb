package main

import (
	"github.com/cloudfoundry/packit"

	"github.com/cage1016/jq-cnb/jq"
)

func main() {
	packit.Detect(jq.Detect())
}
