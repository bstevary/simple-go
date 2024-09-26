package password

import (
	"simple_go/test"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	//create random password
	password := test.RandomString(16)

	hashedPassword1, err := Hash(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = CompareHash(hashedPassword1, password)
	require.NoError(t, err)

	wrongPassword := test.RandomString(16)
	err = CompareHash(hashedPassword1, wrongPassword)
	require.Error(t, err)

	hashedPassword2, err := Hash(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
