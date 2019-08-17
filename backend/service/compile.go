package service

import (
	"errors"

	"github.com/k-nasa/code-hub/model"
)

type Compile struct{}

func NewCompileService() *Compile {
	return &Compile{}
}

func (c *Compile) Run(compile *model.Compile) (*model.CompileResult, error) {
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
}
