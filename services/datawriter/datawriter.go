package datawriter

import (
	"data-pirates-challenge-go-version/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type (
	DataWriterSvc interface {
		SaveJsonlFile(fileNamePrefix string, data []model.ZipCoreRange) error
	}

	dataWriterImp struct {
	}
)

func NewDataWriter() DataWriterSvc {
	return &dataWriterImp{}
}

func (d dataWriterImp) SaveJsonlFile(fileNamePrefix string, data []model.ZipCoreRange) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fmt.Sprintf("%s.jsonl", fileNamePrefix),
		file, 0644); err != nil {
		return err
	}
	return nil
}
