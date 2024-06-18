// Un-used so far

package errors

import (
    "fmt"
    "os"
)

var (
    ErrNoSuchFileOrDir  = fmt.Errorf("no such file or directory")
    ErrNotADirectory    = fmt.Errorf("not a directory")
    ErrPermissionDenied = fmt.Errorf("permission denied")
    ErrTooManyArgs      = fmt.Errorf("too many arguments")
    ErrHomeNotSet       = fmt.Errorf("HOME not set")
)

type ErrorHandler struct{}

// New creates a new ErrorHandler instance
func New() *ErrorHandler {
    return &ErrorHandler{}
}

// HandleError prints the error message and exits the program
func (e *ErrorHandler) HandleError(err error, path string) {
    switch err {
    case ErrNoSuchFileOrDir:
        fmt.Printf("cd: %s: %s\n", err, path)
        os.Exit(1)
    case ErrNotADirectory:
        fmt.Printf("cd: %s: %s\n", err, path)
        os.Exit(1)
    case ErrPermissionDenied:
        fmt.Printf("cd: %s: %s\n", err, path)
        os.Exit(1)
    case ErrTooManyArgs:
        fmt.Printf("cd: %s\n", err)
        os.Exit(1)
    case ErrHomeNotSet:
        fmt.Printf("cd: %s\n", err)
        os.Exit(1)
    default:
        fmt.Printf("cd: %v\n", err)
        os.Exit(1)
    }
}