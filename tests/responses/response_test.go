package tests

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/akecel/gabbro/responses"
)

func TestPrintImageResponse(t *testing.T) {
	url := "https://www.example.com/image.jpg"
	captured := captureOutput(func() {
		responses.PrintImageResponse(url)
	})
	if captured == "" {
		t.Error("Expected output, but got nothing")
	}
}

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestPrintResponse(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	origStdout := os.Stdout
	os.Stdout = w

	testStruct := struct {
		Foo string
		Bar int
	}{
		Foo: "Hello",
		Bar: 123,
	}
	expectedOutput := "Foo: Hello\nBar: 123\n"

	responses.PrintResponse(testStruct)

	w.Close()
	os.Stdout = origStdout

	outputBytes, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	outputStr := string(outputBytes)

	if outputStr != expectedOutput {
		t.Errorf("PrintResponse() = %q, expected %q", outputStr, expectedOutput)
	}
}
