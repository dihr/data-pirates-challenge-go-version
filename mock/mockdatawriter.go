package mock

import (
	"data-pirates-challenge-go-version/model"
	"data-pirates-challenge-go-version/services/datawriter"
)

type mockDataWriter struct {
}

func (m mockDataWriter) SaveJsonlFile(fileNamePrefix string, data []model.ZipCoreRange) error {
	return nil
}

func NewMockDataWriter() datawriter.DataWriterSvc {
	return &mockDataWriter{}
}
