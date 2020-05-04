package feeder_test

import (
	"testing"

	"github.com/sayuthisobri/config/feeder"
	"github.com/stretchr/testify/assert"
)

func TestMap_Feed(t *testing.T) {
	f := feeder.Map{
		"key": "value",
	}

	c, err := f.Feed()
	assert.NoError(t, err)

	assert.Equal(t, "value", c["key"])
}
