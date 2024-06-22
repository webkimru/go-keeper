package grpcserver

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNew(t *testing.T) {

	t.Run("positive", func(t *testing.T) {
		_, err := New("localhost:5001")
		assert.NoError(t, err)
		time.Sleep(2 * time.Second)
	})
	t.Run("positive", func(t *testing.T) {
		_, err := New("localhost:5001")
		assert.Error(t, err)
	})
}
