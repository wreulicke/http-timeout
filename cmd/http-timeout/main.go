package main

import (
	noDefaultClient "github.com/wreulicke/http-timeout/no-default-http-client"
	noDefaultServer "github.com/wreulicke/http-timeout/no-default-http-server"
	noTimeoutClient "github.com/wreulicke/http-timeout/no-timeout-http-server"
	noTimeoutServer "github.com/wreulicke/http-timeout/no-timeout-http-server"

	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		noDefaultClient.Analyzer,
		noDefaultServer.Analyzer,
		noTimeoutClient.Analyzer,
		noTimeoutServer.Analyzer,
	)
}
