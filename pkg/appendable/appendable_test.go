package appendable

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testWriteFile(t *testing.T, name string, index int) {
	file, err := NewFile(name)

	assert.NoError(t, err)
	assert.NotNil(t, file)

	defer file.Close()

	n, err := fmt.Fprintf(file, "%d\n", index+1)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, n)
}

func TestCreate(t *testing.T) {
	t.Parallel()

	name := filepath.Join(os.TempDir(), "create.txt")

	defer os.Remove(name)

	for i := 0; i < 5; i++ {
		testWriteFile(t, name, i)
	}

	data, err := ioutil.ReadFile(name)

	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, "1\n2\n3\n4\n5\n", string(data))
}

func TestOpen(t *testing.T) {
	t.Parallel()

	name := filepath.Join(os.TempDir(), "open.txt")

	defer os.Remove(name)

	assert.NoError(t, ioutil.WriteFile(name, []byte("Hello\n"), 0644))

	for i := 0; i < 5; i++ {
		testWriteFile(t, name, i)
	}

	data, err := ioutil.ReadFile(name)

	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, "Hello\n1\n2\n3\n4\n5\n", string(data))
}

func TestFaileWrite(t *testing.T) {
	t.Parallel()

	// Able to open because not created yet.
	name := filepath.Join(os.TempDir(), "directory")

	defer os.Remove(name)

	file, err := NewFile(name)

	assert.NoError(t, err)
	assert.NotNil(t, file)

	// Create a directory with the same name.
	os.Mkdir(name, 0755)

	n, err := fmt.Fprintln(file, "Hello")

	assert.NoError(t, err)
	assert.NotEqual(t, 0, n)
	assert.Error(t, file.Close())
}
