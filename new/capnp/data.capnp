@0xa58dd07d3ddaf4bb;
using Go = import "/go.capnp";
$Go.package("data");
$Go.import("capnp/data");
struct Command {
  id @0 :Text;
  text @1 :Text;
  date @2 :Text;
}
