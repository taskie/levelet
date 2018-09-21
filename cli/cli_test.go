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
	key := "hello"
	value := "world"
	{
		stdin = strings.NewReader(value)
		os.Args = []string{"levelet", "-f", dbPath, "p", key}
		err = mainImpl()
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		buf := bytes.Buffer{}
		stdout = &buf
		os.Args = []string{"levelet", "-f", dbPath, "g", key}
		err = mainImpl()
		actual := buf.String()
		if actual != value {
			t.Fatalf("invalid value (expected: %s, actual: %s)", value, actual)
		}
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		buf := bytes.Buffer{}
		stdout = &buf
		os.Args = []string{"levelet", "-f", dbPath, "l", "h"}
		err = mainImpl()
		actual := buf.String()
		if actual != key + "\n" {
			t.Fatalf("invalid value (expected: %s, actual: %s)", key + "\n", actual)
		}
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		os.Args = []string{"levelet", "-f", dbPath, "d", key}
		err = mainImpl()
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		os.Args = []string{"levelet", "-f", dbPath, "g", key}
		err = mainImpl()
		if err == nil {
			t.Fatal("get must fail")
		}
	}
}
