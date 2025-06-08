package errors

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

// RecoveryError represents an error recovered from a panic
type RecoveryError struct {
	PanicValue interface{}
	Stack      string
	Location   string
}

func (e *RecoveryError) Error() string {
	return fmt.Sprintf("panic recovered: %v at %s", e.PanicValue, e.Location)
}

// RecoverFromPanic recovers from a panic and converts it to an AppError
func RecoverFromPanic(recovered interface{}) *AppError {
	if recovered == nil {
		return nil
	}

	// get call stack information
	stack := string(debug.Stack())

	// get the location of the panic
	_, file, line, ok := runtime.Caller(2)
	location := "unknown"
	if ok {
		location = fmt.Sprintf("%s:%d", file, line)
	}

	recoveryErr := &RecoveryError{
		PanicValue: recovered,
		Stack:      stack,
		Location:   location,
	}

	return NewSystemError(
		ErrNameInternalError,
		"システムが重大なエラーを発生させました",
		recoveryErr,
		map[string]interface{}{
			"panic_value": fmt.Sprintf("%v", recovered),
			"location":    location,
		},
	)
}

// SafeExecute safely executes a function, captures panic and converts it to an error
func SafeExecute(fn func() error) (err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			err = RecoverFromPanic(recovered)
		}
	}()

	return fn()
}

// SafeExecuteWithResult safely executes a function with a return value, captures panic and converts it to an error
func SafeExecuteWithResult[T any](fn func() (T, error)) (result T, err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			var zero T
			result = zero
			err = RecoverFromPanic(recovered)
		}
	}()

	return fn()
}
