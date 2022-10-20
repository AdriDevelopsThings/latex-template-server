package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/AdriDevelopsThings/latex-template-server/pkg/config"
)

func StartAutoDeleteFiles() {
	for {
		time.Sleep(10 * time.Second)
		files, err := ioutil.ReadDir(config.CurrentConfig.FileServePath)
		if err != nil {
			fmt.Printf("Error while autodeleting files: %v\n", err)
			continue
		}
		for _, file := range files {
			s := strings.Split(file.Name(), "_")
			if len(s) == 2 {
				i, err := strconv.Atoi(s[1])
				if err == nil && time.Now().Unix() > int64(i+config.CurrentConfig.DeleteFileAfter) {
					fmt.Printf("Removing file %s\n", file.Name())
					os.RemoveAll(path.Join(config.CurrentConfig.FileServePath, file.Name()))
				}
			}
		}
	}
}
