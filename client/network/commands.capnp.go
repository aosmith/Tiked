package network

// AUTO GENERATED - DO NOT EDIT

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Command struct{ capnp.Struct }

// Command_TypeID is the unique identifier for the type Command.
const Command_TypeID = 0xa3f682b3ed031bfe

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
	str, _ := text.Marshal(0xa3f682b3ed031bfe, s.Struct)
	return str
}

func (s Command) Cmd() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Command) HasCmd() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Command) CmdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Command) SetCmd(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Command) Target() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Command) HasTarget() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Command) TargetBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Command) SetTarget(v string) error {
	return s.Struct.SetText(1, v)
}

func (s Command) Args() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Command) HasArgs() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Command) ArgsBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Command) SetArgs(v string) error {
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

const schema_d2e84852a9e0a298 = "x\xda\x12\xc8u`2d\xdd\xcf\xc8\xc0\x10(\xc2\xca" +
	"\xf6\xff\x9f4\xf3\xdb\xcdM\xdf\x163\x08\xf23\xfe\x9f" +
	"\xb1\xe8\xc1\xca \x8f\x17\x97\x18X\x99\xd9\x19\x18\x04\x8f" +
	"\xee\x12<\x0b\xa2O\xaeg\xd0\xfd\x9f\x9c\x9f\x9b\x9b\x98" +
	"\x97R\xcc\xa8\x97\x9cX\x90W`\xe5l\x0f\x11\x08`" +
	"d\x0c\xe4afa``ad`\x10tUb`" +
	"\x08t`f\x0c\xf4ab\x14dd\x14a\x04\x09z" +
	"Z10\x04\xba03\x06\x0601\x0a21\x890" +
	"210\x08\xfaj10\x04z03\x06\x8601" +
	"\xb2'\xe7\xa60\xf2001\xf200\xda\x97$\x16" +
	"\xa5\xa7\x96\xc0\xb8\xfc\x89E\xe9\xc50\x0e \x00\x00\xff" +
	"\xff\x8e\xa7(\xa6"

func init() {
	schemas.Register(schema_d2e84852a9e0a298,
		0xa3f682b3ed031bfe)
}
