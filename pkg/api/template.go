package api

import (
	"fmt"
	"net/http"

	"github.com/AdriDevelopsThings/latex-template-server/pkg/apierrors"
	"github.com/AdriDevelopsThings/latex-template-server/pkg/config"
	"github.com/AdriDevelopsThings/latex-template-server/pkg/template"
	"github.com/gin-gonic/gin"
)

type SubmitTemplateArguments struct {
	Arguements map[string]string `json:"arguments"`
}

func SubmitTemplate(c *gin.Context) {
	templateName := c.Param("name")
	arguments := SubmitTemplateArguments{}
	if c.BindJSON(&arguments) != nil {
		return
	}

	fileInfos, err := template.BuildTemplate(templateName, arguments.Arguements)
	if err != nil {
		serr, ok := err.(*apierrors.LatexTemplateServerError)
		if ok {
			serr.Abort(c)
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}
	link := config.CurrentConfig.AppUrl + fmt.Sprintf("/file/%s/%s/%s", fileInfos.ID, fileInfos.EncryptionKey, fileInfos.Name)
	c.JSON(http.StatusOK, gin.H{"link": link})
}
