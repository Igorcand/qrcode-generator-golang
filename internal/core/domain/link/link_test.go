package link_test

import (
	"qrcode-generator/internal/core/domain/link"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestValidadeIfLinkIsEmpty(t *testing.T) {
	link := link.NewLink()
	err := link.Validate()
	require.Error(t, err)
}

func TestLinkIdIsNotUuid(t *testing.T) {
	link := link.NewLink()
	link.ID = "abc"
	link.Url = "abc"
	link.CreatedAt = time.Now()
	err := link.Validate()
	require.Error(t, err)
}

func TestLinkIsValid(t *testing.T) {
	link := link.NewLink()
	link.Url = "abc"
	err := link.Validate()
	require.Nil(t, err)
}
