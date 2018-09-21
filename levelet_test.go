package levelet

import (
	"testing"
	"strings"
	"io/ioutil"
	"os"
	"bytes"
	"path/filepath"
)

func TestMain(t *testing.T) {
	tmp, err := ioutil.TempDir("", "osplus-")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err := os.RemoveAll(tmp)
		if err != nil {
			t.Fatal(err)
		}
	}()

	dbPath := filepath.Join(tmp, "test.ldb")
	expected := "world"
	stdin = strings.NewReader(expected)
	os.Args = []string{"levelet", "-f", dbPath, "p", "hello"}
	Main()
	buf := bytes.Buffer{}
	stdout = &buf
	os.Args = []string{"levelet", "-f", dbPath, "g", "hello"}
	Main()
	actual := buf.String()
	if actual != expected {
		t.Fatalf("invalid value (expected: %s, actual: %s)", expected, actual)
	}
	os.Args = []string{"levelet", "-f", dbPath, "d", "hello"}
	Main()
}
