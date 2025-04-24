package errs

import (
	"encoding/json"
	"fmt"
	"log"
)

type NotFoundErr struct{ error }

func FatalOnError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func Context(obj interface{}) string {
	if obj == nil {
		return "nil"
	}

	marshaled, err := json.Marshal(obj)
	if err != nil {
		return "<json_marshal_error>"
	}

	return string(marshaled)
}

func Wrap(err error, message string, contextObj ...interface{}) error {
	if err == nil {
		return nil
	}

	if len(contextObj) > 0 {
		contextStr := Context(contextObj[0])
		return fmt.Errorf("%s: %w | context: %s", message, err, contextStr)
	}

	return fmt.Errorf("%s: %w", message, err)
}
