package api

import (
	"net/http"

	"fmt"

	"github.com/AdriDevelopsThings/latex-template-server/pkg/apierrors"
	"github.com/AdriDevelopsThings/latex-template-server/pkg/template"
	"github.com/gin-gonic/gin"
)

type SubmitTemplateArguments struct {
	Arguements []map[string]string `json:"arguments"`
}

func SubmitTemplate(c *gin.Context) {
	templateName := c.Param("name")
	arguments := SubmitTemplateArguments{}
	if c.BindJSON(&arguments) != nil {
		return
	}

	file, err := template.BuildTemplate(templateName, arguments.Arguements)
	if err != nil {
		serr, ok := err.(*apierrors.LatexTemplateServerError)
		if ok {
			serr.Abort(c)
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}
	c.Header("Content-Length", fmt.Sprint(len(file)))
	c.Data(http.StatusOK, "application/pdf", file)
}
