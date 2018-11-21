// Package utils will contain different utility functions that are expected to be used throughout the project
package utils

import "fmt"

// HandleError standard function to print error and panic
func HandleError(message string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", message, err)
		panic(err)
	}
}
