package service

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/k-nasa/code-hub/model"
)

type Compile struct{}

func NewCompileService() *Compile {
	return &Compile{}
}

func (c *Compile) Run(compile *model.Compile) (*model.CompileResult, error) {
	tmpDir, err := ioutil.TempDir("", "sandbox")
	if err != nil {
		return nil, fmt.Errorf("error creating temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	tmpFile, err := os.Create(tmpDir + "/exec.go")
	if err != nil {
		return nil, fmt.Errorf("error creating temp file: %v", err)
	}

	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	if _, err := tmpFile.Write([]byte(compile.Body)); err != nil {
		return nil, fmt.Errorf("error write code: %v", err)
	}

	filename := path.Base(tmpFile.Name())
	cmd, err := languageCmd(compile, filename)

	result := createResult(cmd, compile, tmpDir)

	return result, nil
}

func createResult(cmd *exec.Cmd, compile *model.Compile, tmpDir string) *model.CompileResult {
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Dir = tmpDir

	err := cmd.Run()

	if err != nil {
		log.Println(fmt.Sprint(err) + ": " + stderr.String())
		result := &model.CompileResult{
			Language:     compile.Language,
			Ok:           false,
			ErrorMessage: stderr.String(),
		}

		return result
	}

	return &model.CompileResult{
		Language: compile.Language,
		Ok:       true,
		Output:   out.String(),
	}

}

func languageCmd(compile *model.Compile, filename string) (*exec.Cmd, error) {
	switch compile.Language {
	case "golang":
		return exec.Command("go", "run", filename), nil
	case "rust":
		return exec.Command("cargo", "script", filename), nil
	case "ruby":
		return exec.Command("ruby", filename), nil
	default:
		return nil, errors.New("unsupported language")
	}

}
