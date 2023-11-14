package main

import (
	"github.com/codekirei/webchat-poc/backend/config"
	"github.com/codekirei/webchat-poc/cmd"
)

func init() {
	config.Configure()
}

func main() {
	cmd.Execute()
}
