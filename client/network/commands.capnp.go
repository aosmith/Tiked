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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return Command{st}, err
}

func NewRootCommand(s *capnp.Segment) (Command, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
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

func (s Command) Date() int32 {
	return int32(s.Struct.Uint32(0))
}

func (s Command) SetDate(v int32) {
	s.Struct.SetUint32(0, uint32(v))
}

// Command_List is a list of Command.
type Command_List struct{ capnp.List }

// NewCommand creates a new list of Command.
func NewCommand_List(s *capnp.Segment, sz int32) (Command_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
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

const schema_d2e84852a9e0a298 = "x\xda<\xc8\xa1N\x03A\x14\x85\xe1s\xee\x9d\xa5\xa6" +
	"I;\xc98\x1c\x92\x04\x12lMI0H.\x1a\xc1" +
	"dg\xd3\x920e\xd3\x1d\x89\xe2)\xd0\xa0\xd08\x0c" +
	"\x0a\x81\xc0\xf0\x0480X$K6\x04\xd4\x9f\xff\x9b" +
	"^\xee\xcb^\xf5H\xc0B\xb5\xd1\x7fo\xea\xe7\xfd\xd5" +
	"\xd7-lB\xf6\xd77ow\xc7\x87\x1f\xaf\xa8t\x04" +
	"\xf8\xa7\x07\xff2\xf4\xf9\x1d;}}\x91s\\\xa5\x8e" +
	"\xbbulW\xed\xec`\xfe\x0bG\xa4M\xd5\x01\x8e\x80" +
	"\x8f[\x80\x9d(m)\xf4d\xe0\x80\xcd\x0c\xb0S\xa5" +
	"\x9d\x0b\xbdH\xa0\x00\xfel\x1b\xb0\xa4\xb4VH\x0dT" +
	"\xc0\xe7\xc1\x96J+\xc2Q\x9d\x13\xc7\x10\x8e\xc1y\x89" +
	"\xebES\xfev\x12\xd7\x8b\xee\x7fR,\x0d\x1d\x84\x0e" +
	"\xfc\x09\x00\x00\xff\xff\xc1\xee.1"

func init() {
	schemas.Register(schema_d2e84852a9e0a298,
		0xa3f682b3ed031bfe)
}
