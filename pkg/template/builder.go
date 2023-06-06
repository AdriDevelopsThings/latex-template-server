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

const MAX_ARGUMENTS = 5
const MAX_VALUE_LENGTH = 2048

func validateLatexValue(value string) string {
	if len(value) > MAX_VALUE_LENGTH {
		value = value[:MAX_VALUE_LENGTH]
	}
	value = strings.ReplaceAll(value, "\\", "\\textbackslash")
	value = strings.ReplaceAll(value, "\n", "\\\\")
	value = strings.ReplaceAll(value, "_", "\\_")
	return value
}

func argumentsToCSV(filepath string, arguments []map[string]string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	if len(arguments) > MAX_ARGUMENTS {
		arguments = arguments[:MAX_ARGUMENTS-1]
	}
	argumentLength := len(arguments[0])
	argumentIndex := make(map[string]int, 0)
	i := 0
	for key, _ := range arguments[0] {
		argumentIndex[key] = i
		file.WriteString(key + ",")
		i++
	}
	for _, argument := range arguments {
		keys := make([]string, len(argumentIndex))
		if len(argument) != argumentLength {
			continue
		}
		file.WriteString("\n")
		for key, _ := range argument {
			keys[argumentIndex[key]] = key
		}
		for _, key := range keys {
			value := argument[key]
			file.WriteString("\"" + validateLatexValue(value) + "\",")
		}
	}
	file.Close()
	return nil
}

func BuildTemplate(name string, arguments []map[string]string) (*files.FileInfos, error) {
	if strings.Contains(name, "/") || strings.Contains(name, "\\") || strings.Contains(name, ".") {
		return nil, apierrors.TemplateDoesNotExist
	}
	filepath := path.Join(config.CurrentConfig.TemplatePath, name+".tex")
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return nil, apierrors.TemplateDoesNotExist
	}
	dir, err := ioutil.TempDir(config.CurrentConfig.TmpDirectory, "template-"+name)
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	s := string(b)

	file, err := os.Create(path.Join(dir, "latex.tex"))
	if err != nil {
		return nil, err
	}
	file.WriteString(s)
	file.Close()

	//  create arguments csv
	argumentsToCSV(path.Join(dir, "data.csv"), arguments)

	cmd := exec.Command("pdflatex", "latex.tex", "-no-shell-escape")
	cmd.Dir = dir
	_, err = cmd.Output()
	if err != nil {
		return nil, err
	}
	pdf, err := os.ReadFile(path.Join(dir, "latex.pdf"))
	if err != nil {
		return nil, err
	}
	os.RemoveAll(dir)
	return files.WriteFile(name+".pdf", pdf)
}
