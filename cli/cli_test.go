package cli

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	tmp, err := ioutil.TempDir("", "levelet-")
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
	{
		stdin = strings.NewReader(expected)
		os.Args = []string{"levelet", "-f", dbPath, "p", "hello"}
		err = mainImpl()
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		buf := bytes.Buffer{}
		stdout = &buf
		os.Args = []string{"levelet", "-f", dbPath, "g", "hello"}
		err = mainImpl()
		actual := buf.String()
		if actual != expected {
			t.Fatalf("invalid value (expected: %s, actual: %s)", expected, actual)
		}
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		os.Args = []string{"levelet", "-f", dbPath, "d", "hello"}
		err = mainImpl()
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		os.Args = []string{"levelet", "-f", dbPath, "g", "hello"}
		err = mainImpl()
		if err == nil {
			t.Fatal("get must fail")
		}
	}
}
