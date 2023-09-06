package names_test

import (
	"bytes"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tilotech/go-names"
)

func canonicalTestFixture(t *testing.T) *names.Canonical {
	data := `Neo,One,Keanu,Thomas A. Anderson
Smith,Agent,Hugo
Cypher,Reagan,Joe`

	buf := bytes.NewBufferString(data)

	fixture, err := names.NewCanonical(buf)
	require.NoError(t, err)

	return fixture
}

func TestNewCanonical(t *testing.T) {
	data := `Foo`
	_, err := names.NewCanonical(
		bytes.NewBufferString(data),
	)
	assert.Error(t, err)

	_, err = names.NewCanonical(
		iotest.ErrReader(assert.AnError),
	)
	assert.Error(t, err)
}

func TestOf(t *testing.T) {
	fixture := canonicalTestFixture(t)

	actual := fixture.Of("One")
	assert.Equal(t, "Neo", actual)

	actual = fixture.Of("Agent")
	assert.Equal(t, "Smith", actual)

	actual = fixture.Of("Neo")
	assert.Equal(t, "Neo", actual)

	actual = fixture.Of("Keanu")
	assert.Equal(t, "Neo", actual)

	actual = fixture.Of("Trinity")
	assert.Equal(t, "Trinity", actual)
}

func TestCanonicalPreset(t *testing.T) {
	fixture, err := names.NewCanonicalPreset("NICKNAMES")
	require.NoError(t, err)

	actual := fixture.Of("billy")
	assert.Equal(t, "bill", actual)

	_, err = names.NewCanonicalPreset("UNKNOWN")
	assert.Error(t, err)
}
