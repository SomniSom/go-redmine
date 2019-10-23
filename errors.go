package redmine

import "github.com/ztrue/tracerr"

func printError(err error) {
	if err != nil {
		tracerr.PrintSourceColor(err)
	}
}
