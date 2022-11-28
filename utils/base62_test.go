package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBase62(t *testing.T) {
	var value uint64 = 642557611423047429

	//  ub4EELBXnk
	hash := Encode(value)
	result := Decode(hash)

	require.NotEmpty(t, hash)
	require.NotEmpty(t, result)

	require.Equal(t, result, value)
	require.Equal(t, hash, "ub4EELBXnk")
}
