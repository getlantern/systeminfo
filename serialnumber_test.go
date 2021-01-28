package systeminfo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSerialNumber(t *testing.T) {
	serial, err := SerialNumber()
	require.NoError(t, err)
	require.NotEmpty(t, serial)
	t.Log(serial)
}
