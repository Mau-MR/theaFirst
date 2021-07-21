package DB_test

import (
	"github.com/Mau-MR/theaFirst/DB"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewResponseWrapper(t *testing.T) {
	eWrapper := NewElasticWrapperTest(t)
	query := eWrapper.BuildSearchQueryByFields("Mauricio", []string{"name"})
	res, err := eWrapper.ESearchWithDefault("test", query)
	assert.NoError(t, err)
	defer res.Body.Close()
	rw, err := DB.NewResponseWrapper(res)
	assert.NoError(t, err)
	assert.NotNil(t, rw)
}
