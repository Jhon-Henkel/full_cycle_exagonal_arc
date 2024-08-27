package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJsonError(t *testing.T) {
	msg := "hello json"
	result := JsonError(msg)

	require.Equal(t, []byte(`{"message":"hello json"}`), result)
}
