package cmd

import (
	"bytes"
	"io"
	"testing"
)

func TestPing(t *testing.T) {
	cmd := rootCmd
	b := new(bytes.Buffer)
	cmd.SetOut(b)
	cmd.SetArgs([]string{"ping"})
	cmd.Execute()
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatalf("Expected \"%s\" got \"%s\"", "Pong", string(out))
	}
}
