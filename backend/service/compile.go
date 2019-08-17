package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

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

	tmpFile, err := ioutil.TempFile(tmpDir, "ioutil")
	if err != nil {
		return nil, fmt.Errorf("error creating temp file: %v", err)
	}

	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	if _, err := tmpFile.Write([]byte(compile.Body)); err != nil {
		return nil, fmt.Errorf("error write code: %v", err)
	}

	filepath := tmpFile.Name()

	// ここでコマンド実行する
	switch compile.Language {
	case "golang":
		return nil, nil
	case "rust":
		return nil, nil
	case "ruby":
		return nil, nil
	default:
		return nil, errors.New("unsupported language")
	}

	// 実行結果を取得する

	// status みてOKか判定する

	// オブジェクトを作成して返す
}
