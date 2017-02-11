@0xd2e84852a9e0a298;
using Go = import "/go.capnp";
$Go.package("network");
$Go.import("network/network");
struct Command {
  cmd @0 :Text;
  target @1 :Text;
  args @2 :Text;
  date @3 :Int64;
}
capnp compile -I$GOPATH/src/zombiezen.com/go/capnproto2/std -ogo *.capnp
