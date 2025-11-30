package main

import (
	"github.com/thebigmatchplayer/ton-exporter/cmd"
	"github.com/thebigmatchplayer/ton-exporter/utils"
)

func init() {
	utils.InitLogger()
}

func main() {
	cmd.Execute()
}
