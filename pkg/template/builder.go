package template

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/AdriDevelopsThings/latex-template-server/pkg/apierrors"
	"github.com/AdriDevelopsThings/latex-template-server/pkg/config"
	"github.com/AdriDevelopsThings/latex-template-server/pkg/files"
)

func BuildTemplate(name string, arguments map[string]string) (*files.FileInfos, error) {
	filepath := path.Join(config.CurrentConfig.TemplatePath, name+".tex")
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return nil, apierrors.TemplateDoesNotExist
	}
	dir, err := ioutil.TempDir(config.CurrentConfig.TmpDirectory, "template-"+name)
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)

	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	s := string(b)
	for key, value := range arguments {
		value = strings.Replace(value, "\n", " \\\\\n", -1)
		s = strings.Replace(s, "__"+strings.ToUpper(key)+"__", value, -1)
	}

	file, err := os.Create(path.Join(dir, "latex.tex"))
	if err != nil {
		return nil, err
	}
	file.WriteString(s)
	file.Close()

	cmd := exec.Command("pdflatex", "latex.tex")
	cmd.Dir = dir
	_, err = cmd.Output()
	if err != nil {
		return nil, err
	}
	pdf, err := os.ReadFile(path.Join(dir, "latex.pdf"))
	if err != nil {
		return nil, err
	}
	return files.WriteFile(name+".pdf", pdf)
}