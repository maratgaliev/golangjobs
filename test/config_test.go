package test

import (
	"github.com/maratgaliev/golangjobs/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_NewConfig(t *testing.T) {
	c, err := config.NewConfig("../config/config.yaml")
	assert.Nil(t, err)
	assert.NotNil(t, c)
}
