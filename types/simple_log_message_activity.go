package types

import (
	"github.com/golang-devops/easy-workflow-manager/logging"
)

func NewSimplePrintMessageActivity(logger logging.Logger, name, msg string, nextNode Node) Activity {
	return NewSimpleActivity(
		name,
		func() error {
			logger.Info(msg)
			return nil
		},
		nextNode,
	)
}
