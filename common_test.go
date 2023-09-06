package names_test

import (
	"bytes"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tilotech/go-names"
)

func commonTestFixture(t *testing.T) *names.Common {
	data := `Neo,0.213
Morpheus,0.171
Trinity,0.157
Smith,0.143
Oracle,0.064
Niobe,0.057
Cypher,0.043
Seraph,0.036
Architect,0.021`
	buf := bytes.NewBufferString(data)

	fixture, err := names.NewCommon(buf)
	require.NoError(t, err)

	return fixture
}

func TestNewCommon(t *testing.T) {
	data := `Name,0.1234,Foo`
	_, err := names.NewCommon(
		bytes.NewBufferString(data),
	)
	assert.Error(t, err)

	data = `Name,Foo`
	_, err = names.NewCommon(
		bytes.NewBufferString(data),
	)
	assert.Error(t, err)

	_, err = names.NewCommon(
		iotest.ErrReader(assert.AnError),
	)
	assert.Error(t, err)
}

func TestTopN(t *testing.T) {
	fixture := commonTestFixture(t)

	actual := fixture.Top(1)
	assert.Equal(t, []string{"Neo"}, actual)

	actual = fixture.Top(3)
	assert.Equal(t, []string{"Neo", "Morpheus", "Trinity"}, actual)

	actual = fixture.Top(100)
	assert.Equal(t, []string{"Neo", "Morpheus", "Trinity", "Smith", "Oracle", "Niobe", "Cypher", "Seraph", "Architect"}, actual)

	actual = fixture.Top(0)
	assert.Equal(t, []string{}, actual)

	actual = fixture.Top(-1)
	assert.Equal(t, []string{}, actual)
}

func TestFrequency(t *testing.T) {
	fixture := commonTestFixture(t)

	actual := fixture.Frequency("Neo")
	assert.Equal(t, float32(0.213), actual)

	actual = fixture.Frequency("Oracle")
	assert.Equal(t, float32(0.064), actual)

	actual = fixture.Frequency("John")
	assert.Equal(t, float32(0), actual)
}

func TestCommonPreset(t *testing.T) {
	fixture, err := names.NewCommonPreset("US_FIRST_NAME")
	require.NoError(t, err)

	actual := fixture.Top(1)
	assert.Len(t, actual, 1)

	_, err = names.NewCommonPreset("UNKNOWN")
	assert.Error(t, err)
}
