package helpers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func CreatePostContext(msg ...string) *gin.Context {
	var body *bytes.Buffer
	gin.SetMode(gin.TestMode)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	if len(msg) == 0 {
		body = bytes.NewBufferString("Fetch binary post data")
	} else {
		body = bytes.NewBufferString(msg[0])
	}
	c.Request, _ = http.NewRequest("POST", "/", body)
	c.Request.Header.Add("Content-Type", gin.MIMEPOSTForm)

	return c
}
