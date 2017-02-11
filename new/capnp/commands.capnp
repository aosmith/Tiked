@0xd2e84852a9e0a298;
using Go = import "/go.capnp";
$Go.package("commands");
$Go.import("capnp/commands");
struct Command {
  cmd @0 :Text;
  target @1 :Text;
  args @2 :Text;
}
//capnp compile -I$GOPATH/src/zombiezen.com/go/capnproto2/std -ogo *.capnp
