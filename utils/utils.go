// Package utils will contain different utility functions that are expected to be used throughout the project
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// HandleError standard function to print error and panic
func HandleError(message string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", message, err)
		panic(err)
	}
}

// BodyLogger Gin middleware to log the body of a request
// Currently only works with json
func BodyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
			var prettyJSON bytes.Buffer
			json.Indent(&prettyJSON, bodyBytes, "", "    ")
			fmt.Fprintln(gin.DefaultWriter, string(prettyJSON.Bytes()))
		}

		// Point Body to new reader as the previous reader was consumed
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Next()
	}
}
