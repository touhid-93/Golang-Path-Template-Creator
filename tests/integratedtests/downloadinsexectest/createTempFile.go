package downloadinsexectest

import (
	"io/ioutil"
	"testing"
)

func createTempFile(t *testing.T) (string, []byte) {
	dir := t.TempDir()

	tempFile, err := ioutil.TempFile(dir, Test2MB)
	if err != nil {
		t.Fatal("creating temp file:", err)
	}

	randText := []byte("randText")
	if _, err := tempFile.Write(randText); err != nil {
		t.Fatal("write to temp file:", err)
	}

	defer tempFile.Close()
	// tempFile
	return tempFile.Name(), randText
}
