package scraper

import (
	"data-pirates-challenge-go-version/mock"
	"data-pirates-challenge-go-version/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScraperImp_GetUfListFromDoc(t *testing.T) {
	api := mock.NewMockCorreiosApi()
	writer := mock.NewMockDataWriter()
	svc := NewScraperSvc(api, writer)
	data, err := svc.GetUfListFromDoc()
	assert.Nil(t, err)
	assert.NotNil(t, []string{}, data)

	expect := []string{"AC", "AL", "AM", "AP"}
	assert.Equal(t, expect, data)
}

func TestScraperImp_GetZipCodeRangeFromDoc(t *testing.T) {
	api := mock.NewMockCorreiosApi()
	writer := mock.NewMockDataWriter()
	svc := NewScraperSvc(api, writer)
	data, nextPage, err := svc.GetZipCodeRangeFromDoc("SP", nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, false, nextPage)
	assert.NotNil(t, [][]string{}, data)

	expected := make([][]string, 0)
	localrionData := []string{"AcrelÃ¢ndia",
		" 69945-000 a 69949-999",
		"NÃ£o codificada por logradouros",
		"Total do municÃ\u00adpio"}

	expected = append(expected, localrionData)
	assert.Equal(t, expected, data)
}

func TestScraperImp_StartScraping(t *testing.T) {
	api := mock.NewMockCorreiosApi()
	writer := mock.NewMockDataWriter()
	svc := NewScraperSvc(api, writer)
	err := svc.StartScraping()
	assert.Nil(t, err)
}

func TestRemoveDuplicates(t *testing.T) {
	data := make([]model.ZipCoreRange, 0)

	locationA := model.ZipCoreRange{
		Location: "a",
		Range:    "b",
	}

	locationB := model.ZipCoreRange{
		Location: "a",
		Range:    "b",
	}
	data = append(data, locationA, locationB)
	result := RemoveDuplicates(data)

	assert.Equal(t, []model.ZipCoreRange{model.ZipCoreRange{
		Location: "a",
		Range:    "b",
	}}, result)
}
