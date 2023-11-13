package Error

import (
	"fmt"
)

func CustomErrorHandler(message string, code int) error {
	return fmt.Errorf("custom error handler (message: %s) %d", message, code)
}
