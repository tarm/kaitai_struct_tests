package spec

import "github.com/tarm/kaitai_struct_go_runtime/kaitai"

type NavParent struct {
	Header HeaderObj
	Index  IndexObj
}

func (n *NavParent) Unmarshal(s *kaitai.Stream) (err error) {
	if err := n.Header.Unmarshal(s); err != nil {
		return err
	}
	if err := n.Index.Unmarshal(s, n.Header.QtyEntries, n.Header.FilenameLen); err != nil {
		return err
	}
	return nil
}

type HeaderObj struct {
	QtyEntries  uint32
	FilenameLen uint32
}

func (h *HeaderObj) Unmarshal(s *kaitai.Stream) (err error) {
	if h.QtyEntries, err = s.ReadU4le(); err != nil {
		return err
	}
	if h.FilenameLen, err = s.ReadU4le(); err != nil {
		return err
	}
	return nil
}

type IndexObj struct {
	Magic   uint32
	Entries []Entry
}

func (o *IndexObj) Unmarshal(s *kaitai.Stream, qtyEntries uint32, filenameLen uint32) (err error) {
	if o.Magic, err = s.ReadU4le(); err != nil {
		return err
	}
	o.Entries = make([]Entry, qtyEntries)

	for i := range o.Entries {
		if err = o.Entries[i].Unmarshal(s, filenameLen); err != nil {
			return err
		}
	}
	return nil
}

type Entry struct {
	Filename string
}

func (e *Entry) Unmarshal(s *kaitai.Stream, filenameLen uint32) (err error) {
	// FIXME I'm unhappy about the conversion to int here
	buf, err := s.ReadBytes(int(filenameLen))
	if err != nil {
		return err
	}
	e.Filename = string(buf)
	return nil
}
