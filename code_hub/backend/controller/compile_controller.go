package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/k-nasa/code-hub/model"
	"github.com/k-nasa/code-hub/service"
)

type Compile struct{}

func NewCompile() *Compile {
	return &Compile{}
}

func (c *Compile) Run(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	compile := &model.Compile{}

	if err := json.NewDecoder(r.Body).Decode(&compile); err != nil {
		log.Println("decode error")
		return http.StatusBadRequest, nil, err
	}

	if compile.Language == "" || compile.Body == "" {
		log.Printf("empty: %v", compile)
		return http.StatusBadRequest, nil, nil
	}

	compileService := service.NewCompileService()
	compileResult, err := compileService.Run(compile)

	if err != nil {
		return http.StatusBadRequest, map[string]string{"message": err.Error()}, err
	}

	return http.StatusOK, compileResult, nil
}
