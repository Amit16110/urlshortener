package utils

import (
	"testing"
)

func TestBase62(t *testing.T) {

	// 642557611423047429 ub4EELBXnk
	hash := Encode(642557611423047429)
	result := Decode(hash)

	requires.NotEmpty(hash)
	requires.NotEmpty(result)

	requires.Equal(result, 642557611423047429)
	requires.Equal(hash, "ub4EELBXnk")
}
