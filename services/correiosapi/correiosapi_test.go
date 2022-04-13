package correiosapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorreiosApiAimp_GetHomePage(t *testing.T) {
	svc := NewCorreiosApiService()
	data, err := svc.GetHomePage()
	assert.Nil(t, err)
	assert.NotNil(t, data)
}

func TestCorreiosApiAimp_GetData(t *testing.T) {
	svc := NewCorreiosApiService()
	data, err := svc.GetData("SP", nil, nil)
	assert.Nil(t, err)
	assert.NotNil(t, data)
}
