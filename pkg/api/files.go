package api

import (
	"fmt"
	"net/http"

	"github.com/AdriDevelopsThings/latex-template-server/pkg/apierrors"
	"github.com/AdriDevelopsThings/latex-template-server/pkg/files"
	"github.com/gin-gonic/gin"
)

func GetFile(c *gin.Context) {
	name := c.Param("name")
	id := c.Param("id")
	encryption_key := c.Param("encryptionKey")
	content, err := files.ReadFile(id, name, encryption_key)
	if err != nil {
		serr, ok := err.(*apierrors.LatexTemplateServerError)
		if ok {
			serr.Abort(c)
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Length", fmt.Sprint(len(content)))
	c.Writer.Write(content)
}
