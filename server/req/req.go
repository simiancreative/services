package req

import "github.com/gin-gonic/gin"

type Context struct {
	ID      string
	Parsed  Parsed
	Context *gin.Context
}

type Parsed struct {
	Headers map[string]string
	Params  map[string]string
	Body    []byte
	Input   interface{}
}
