package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJson(t *testing.T) {

	book := Book{Title: "mybook", Author: "eoin", ISBN: "12345"}
	marshalled := book.ToJson()

	assert.Equal(t, `{"Title":"mybook","Author":"eoin","ISBN":"12345"}`, string(marshalled),
		"json marshalling wrong")

}

func TestFromJson(t *testing.T) {

	data := []byte(`{"Title":"mybook","Author":"eoin","ISBN":"12345"}`)
	createdBook := Book{Title: "mybook", Author: "eoin", ISBN: "12345"}
	book := FromJson(data)

	assert.Equal(t, book, createdBook, "fromjson failed")

}
