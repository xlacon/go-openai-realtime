package openairt_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	openairt "github.com/xlacon/go-openai-realtime"
)

func TestPermanent(t *testing.T) {
	original := errors.New("test")
	err := openairt.Permanent(original)
	require.ErrorIs(t, err, original)
	var permanent *openairt.PermanentError
	require.ErrorAs(t, err, &permanent)
	require.ErrorIs(t, permanent.Err, original)
	require.EqualError(t, err, original.Error())
}
