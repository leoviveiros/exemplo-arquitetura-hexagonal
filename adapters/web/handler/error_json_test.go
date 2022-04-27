package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "erro message 0x00"
	result := jsonError(msg)

	require.Equal(t, []byte(`{"message":"erro message 0x00"}`), result)
}