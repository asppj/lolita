package statics

import (
	"io"
	"io/ioutil"
	"os"
	"t-mk-opentrace/ext/log-driver/log"

	"github.com/gin-gonic/gin"
)

// ServerFile stream->file
func ServerFile(ctx *gin.Context) {
	log.Info(ctx.GetHeader("Context-Type"))
	fd, err := os.Open("C:\\Users\\asppj\\Pictures\\尘封\\1.png")
	if err != nil {
		log.Error(err)
	}
	body, err := ioutil.ReadAll(fd)
	if err != nil {
		log.Error(err)
	}
	// http.ServeContent(ctx.Writer, ctx.Request, "file_main.go", time.Now(), fd)
	ctx.Writer.Header().Set("Content-Disposition", "attachment; filename=1.png")
	ctx.Writer.Header().Set("Content-Type", ctx.Request.Header.Get("Content-Type"))
	_, err = ctx.Writer.Write(body)
	if err != nil {
		log.Error(err)
	}
}

// Stream  stream
func Stream(ctx *gin.Context) {
	fd, err := os.Open("./go.mod")
	if err != nil {
		log.Error(err)
	}
	ctx.Writer.Header().Set("Content-Disposition", "attachment; filename=go.mod")
	ctx.Writer.Header().Set("Content-Type", ctx.Request.Header.Get("Content-Type"))
	n, err := io.Copy(ctx.Writer, fd)
	log.Info(n)
}
