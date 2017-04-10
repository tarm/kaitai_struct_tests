package spec

import "github.com/tarm/kaitai_struct_go_runtime/kaitai"

type HelloWorld struct {
	One uint8
}

func (h *HelloWorld) Unmarshal(s *kaitai.Stream) (err error) {
	if h.One, err = s.ReadU1(); err != nil {
		return err
	}
	return err
}
