package ical

import (
	"testing"

	"github.com/michael-gallo/simpleical/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAppendExtensionProperty(t *testing.T) {
	var xProp, ianaProp []model.ExtensionProperty

	params := map[string]string{"VALUE": "URI", "FMTTYPE": "audio/basic"}
	appendExtensionProperty(&xProp, &ianaProp, "X-ABC-MMSUBJ", "http://example.org/a.au", params)
	require.Len(t, xProp, 1)
	assert.Empty(t, ianaProp)
	assert.Equal(t, "X-ABC-MMSUBJ", xProp[0].Name)
	assert.Equal(t, "http://example.org/a.au", xProp[0].Value)
	assert.Equal(t, map[string]string{"VALUE": "URI", "FMTTYPE": "audio/basic"}, xProp[0].Params)

	// Mutating the source params must not affect the stored copy.
	params["VALUE"] = "TEXT"
	assert.Equal(t, "URI", xProp[0].Params["VALUE"])

	appendExtensionProperty(&xProp, &ianaProp, "DRESSCODE", "CASUAL", nil)
	require.Len(t, ianaProp, 1)
	assert.Equal(t, "DRESSCODE", ianaProp[0].Name)
	assert.Equal(t, "CASUAL", ianaProp[0].Value)
	assert.Nil(t, ianaProp[0].Params)

	appendExtensionProperty(&xProp, &ianaProp, "X-FOO", "one", nil)
	appendExtensionProperty(&xProp, &ianaProp, "X-FOO", "two", nil)
	require.Len(t, xProp, 3)
	assert.Equal(t, "one", xProp[1].Value)
	assert.Equal(t, "two", xProp[2].Value)
}

func TestIsXName(t *testing.T) {
	assert.True(t, isXName("X-FOO"))
	assert.True(t, isXName("x-foo"))
	assert.True(t, isXName("X-"))
	assert.False(t, isXName("DRESSCODE"))
	assert.False(t, isXName("XFOO"))
	assert.False(t, isXName("X"))
	assert.False(t, isXName(""))
}
