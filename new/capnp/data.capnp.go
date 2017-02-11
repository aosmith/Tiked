package data

// AUTO GENERATED - DO NOT EDIT

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Command struct{ capnp.Struct }

// Command_TypeID is the unique identifier for the type Command.
const Command_TypeID = 0xcbdd1e426851d0b3

func NewCommand(s *capnp.Segment) (Command, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Command{st}, err
}

func NewRootCommand(s *capnp.Segment) (Command, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Command{st}, err
}

func ReadRootCommand(msg *capnp.Message) (Command, error) {
	root, err := msg.RootPtr()
	return Command{root.Struct()}, err
}

func (s Command) String() string {
	str, _ := text.Marshal(0xcbdd1e426851d0b3, s.Struct)
	return str
}

func (s Command) Id() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Command) HasId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Command) IdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Command) SetId(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Command) Text() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Command) HasText() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Command) TextBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Command) SetText(v string) error {
	return s.Struct.SetText(1, v)
}

func (s Command) Date() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Command) HasDate() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Command) DateBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Command) SetDate(v string) error {
	return s.Struct.SetText(2, v)
}

// Command_List is a list of Command.
type Command_List struct{ capnp.List }

// NewCommand creates a new list of Command.
func NewCommand_List(s *capnp.Segment, sz int32) (Command_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return Command_List{l}, err
}

func (s Command_List) At(i int) Command { return Command{s.List.Struct(i)} }

func (s Command_List) Set(i int, v Command) error { return s.List.SetStruct(i, v.Struct) }

// Command_Promise is a wrapper for a Command promised by a client call.
type Command_Promise struct{ *capnp.Pipeline }

func (p Command_Promise) Struct() (Command, error) {
	s, err := p.Pipeline.Struct()
	return Command{s}, err
}

const schema_a58dd07d3ddaf4bb = "x\xda\x12hv`2d\xcdgb`\x08\x94ae" +
	"\xfb\xbf\xf9B`\x86\x93\xdc\xdd\xd3\x0c\x82\xdc\x8c\xffw" +
	"\x7f\xb9e[{\xa1w)\x03+3;\x03\x83\xb0," +
	"\xe3,aUF\x10K\x91q=\x83\xf3\xff\x94\xc4\x92" +
	"D\xbd\xe4\xc4\x02\xc6\xbc\x02+\xe7\xfc\xdc\xdcD\xe6\xbc" +
	"\x94\x00F\xc6@\x1ef\x16\x06\x06\x16F\x06\x06AW" +
	")\x06\x86@\x07f\xc6@\x1f&FAFF\x11F" +
	"\x90\xa0\xa7\x16\x03C\xa0\x0b3c`\x00\x13\xa3 \x13" +
	"\x93\x08#\x13\x03\x83\xa0/H\xd0\x83\x9910\x84\x89" +
	"\x9193\x85\x91\x87\x81\x89\x91\x87\x81\x91\xbf$\xb5\xa2" +
	"\x04\xceII,I\x85q\x00\x01\x00\x00\xff\xffq\xd9" +
	"\"\x9b"

func init() {
	schemas.Register(schema_a58dd07d3ddaf4bb,
		0xcbdd1e426851d0b3)
}
