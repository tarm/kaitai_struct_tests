package spec

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tarm/kaitai_struct_go_runtime/kaitai"
)

func TestHelloWorld(t *testing.T) {
	f, err := os.Open("../../src/fixed_struct.bin")
	if err != nil {
		t.Fatal(err)
	}
	s := kaitai.NewStream(f)

	var h HelloWorld
	err = h.Unmarshal(s)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, h.One, byte(0x50), "They should be equal")
}
