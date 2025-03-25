package apierrors

import "github.com/gin-gonic/gin"

type LatexTemplateServerError struct {
	Name       string
	HttpStatus int
}

func (err *LatexTemplateServerError) Error() string {
	return err.Name
}

func (err *LatexTemplateServerError) Abort(c *gin.Context) {
	c.AbortWithStatusJSON(err.HttpStatus, gin.H{"error": err.Name})
}

var TemplateDoesNotExist = &LatexTemplateServerError{Name: "template_does_not_exist", HttpStatus: 404}
