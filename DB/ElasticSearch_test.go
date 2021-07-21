package DB_test

import (
	"github.com/Mau-MR/theaFirst/DB"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func NewElasticWrapperTest(t *testing.T) *DB.ElasticWrapper {
	eUser := os.Getenv("EUSER")
	ePassword := os.Getenv("EPASSWORD")
	l := log.New(os.Stdout, "[Keybons-Test] ", log.LstdFlags)
	eWrapper, err := DB.NewElasticWrapper("https://localhost:9200", eUser, ePassword, l)
	assert.NoError(t, err)
	assert.NotNil(t, eWrapper)
	return eWrapper
}

func TestNewElasticWrapper(t *testing.T) {
	eUser := os.Getenv("EUSER")
	ePassword := os.Getenv("EPASSWORD")
	l := log.New(os.Stdout, "[Keybons-Test] ", log.LstdFlags)
	eWrapper, err := DB.NewElasticWrapper("https://localhost:9200", eUser, ePassword, l)
	assert.NoError(t, err)
	assert.NotNil(t, eWrapper)
}

func TestElasticWrapper_InsertStructTo(t *testing.T) {
	eWrapper := NewElasticWrapperTest(t)
	type Test struct {
		Name string `json:"name"`
	}
	test := &Test{
		Name: "Mauricio",
	}
	_, err := eWrapper.InsertStructTo("test", "someID", test)
	assert.NoError(t, err)
	_, err = eWrapper.DeleteDocumentByID("test", "someID")
	assert.NoError(t, err)

}

func TestElasticWrapper_DeleteDocumentByID(t *testing.T) {
	eWrapper := NewElasticWrapperTest(t)
	_, err := eWrapper.DeleteDocumentByID("test", "SomeErroneousID")
	assert.Error(t, err)
}

func TestElasticWrapper_BuildSearchQueryByFields(t *testing.T) {
	eWrapper := NewElasticWrapperTest(t)
	fields := []string{"name"}
	query := eWrapper.BuildSearchQueryByFields("Probando algo", fields)
	assert.NotNil(t, query)
	t.Log(query)
}
func TestElasticWrapper_eSearchWithDefault(t *testing.T) {
	eWrapper := NewElasticWrapperTest(t)
	query := eWrapper.BuildSearchQueryByFields("Mauricio", []string{"Mauricio"})
	res, err := eWrapper.ESearchWithDefault("test", query)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestElasticWrapper_SearchStructIn(t *testing.T) {
	eWrapper := NewElasticWrapperTest(t)
	query := eWrapper.BuildSearchQueryByFields("Mauricio", []string{"name"})
	res, err := eWrapper.SearchIn("costumers", query)
	for _, v := range res.Hits.Hits {
		log.Println(v)
	}
	assert.NoError(t, err)
}
