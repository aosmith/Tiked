package commands

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

const schema_d2e84852a9e0a298 = "x\xda\x12hv`2d\xcdgb`\x08\x94ae" +
	"\xfb\xffO\x9a\xf9\xed\xe6\xa6o\x8b\x19\x04\xf9\x19\xff\xcf" +
	"X\xf4`e\x90\xc7\x8bK\x0c\xac\xcc\xec\x0c\x0c\xc2\xb2" +
	"\x8c\xbb\x84U\x19A,E\xc6\xf5\x0c\xce\xff\x93\xf3s" +
	"s\x13\xf3R\x8a\x19\xf5\x92\x13\x0b\xf2\x0a\xac\x9c\xed!" +
	"\x02\x01\x8c\x8c\x81<\xcc,\x0c\x0c,\x8c\x0c\x0c\x82\xae" +
	"J\x0c\x0c\x81\x0e\xcc\x8c\x81>L\x8c\x82\x8c\x8c\"\x8c" +
	" AO+\x06\x86@\x17f\xc6\xc0\x00&FA&" +
	"&\x11F&\x06\x06A_-\x06\x86@\x0ff\xc6\xc0" +
	"\x10&F\xf6\xe4\xdc\x14F\x1e\x06&F\x1e\x06F\xfb" +
	"\x92\xc4\xa2\xf4\xd4\x12\x18\x97?\xb1(\xbd\x18\xc6\x01\x04" +
	"\x00\x00\xff\xff\x93.&\x9c"

func init() {
	schemas.Register(schema_d2e84852a9e0a298,
		0xa3f682b3ed031bfe)
}
