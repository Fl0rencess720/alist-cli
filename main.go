package main

import (
	"github.com/Fl0rencess720/alist-cli/cmd"
	"github.com/Fl0rencess720/alist-cli/internal/config"
	"github.com/Fl0rencess720/alist-cli/internal/httpclient"
	"github.com/Fl0rencess720/alist-cli/internal/pkgs/logger"
)

func init() {
	config.Init()
	logger.Init()
	httpclient.InitClient()
}

func main() {
	cmd.Execute()
}
