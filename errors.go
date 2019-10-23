package redmine

import (
	"encoding/json"
	"errors"
	"github.com/ztrue/tracerr"
	"strings"
)

func printError(err error) {
	if err != nil {
		tracerr.PrintSourceColor(err)
	}
}

func checkStatusCode(statusCode int, body []byte) error {
	if statusCode == 403 {
		return errors.New("Access denied ")
	}

	var er errorsResult
	err := json.Unmarshal(body, &er)
	if err == nil {
		err = errors.New(strings.Join(er.Errors, "\n"))
	}

	return err
}
