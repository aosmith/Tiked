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

func (s Command) Date() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s Command) SetDate(v int64) {
	s.Struct.SetUint64(0, uint64(v))
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

const schema_d2e84852a9e0a298 = "x\xda<\xc8\xa1N\x03A\x14\x85\xe1s\xee\xedP\xd3" +
	"\xa4\x9dd\x1c\x0eI\x02\x09\xb6\xa6$\x18$\x17\x8d`" +
	"\xb2\xb3iI\x98\xb2\xd9\x1d\x89\xe2)\xd0\xa0\xd08\x0c" +
	"\x0a\x81\xc0\xf0\x0480X$K6\x04\xd4\x9f\xff\x9b" +
	"]\xee\xcb\x9e{$`\xc1m\xf4\xdf\x9b\xfay\x7f\xf5" +
	"u\x0b\x9b\x92\xfd\xf5\xcd\xdb\xdd\xf1\xe1\xc7+\x9c\x8e\x01" +
	"\xff\xf4\xe0_\x86>\xbfc\xa7\xaf.r\x8e\xeb\xd4q" +
	"\xb7\x8a\xcd\xba\x99\x1f,~\xe1\x88\xb4\x99\x8e\x80\x11\x01" +
	"\x1f\xb7\x00;Q\xdaJ\xe8\xc9\xc0\x01\xeb9`\xa7J" +
	";\x17z\x91@\x01\xfc\xd96`Ii\x8d\x90\x1a\xa8" +
	"\x80\xcf\x83\xad\x94V\x84\xe3*'N \x9c\x80\x8b\x12" +
	"\xdbe]\xfev\x1a\xdbe\xf7?)\x96\x9a\x0eB\x07" +
	"\xfe\x04\x00\x00\xff\xff\xc1\xf8.3"

func init() {
	schemas.Register(schema_d2e84852a9e0a298,
		0xa3f682b3ed031bfe)
}
