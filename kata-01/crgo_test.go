package kata_01

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRequest_CannotMakeARequestWithInvalidURL(t *testing.T) {
	err := Request("full")
	assert.EqualError(t, err, "Invalid URL")
}

func TestRequest_ShouldReturnNilForAValidURL(t *testing.T) {
	err := Request("http://www.google.com")
	assert.Equal(t, nil, err, "Url must return valid.")
}
