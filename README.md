# Fork of [tronprotocol/protocol](https://github.com/tronprotocol/protocol) for Tron Proto sharing within Firehose/Substreams ecosystem

We fork the https://github.com/tronprotocol/protocol because we wanted have a StreamingFast source for depending on Tron protocol protobufs.

This objectives of this repository are:
- Remain binary compatible with Tron protobuf
- Keep the Tron package ID `protocol` so that Google Any fits
- Rename Go package to points to github.com/streamingfast/tron-protocol/pb
- Publish a shareable and re-usable version of [buf.build/streamingfast/tron-protocol]

The changes we made:
- Replaced `option go_package = "github.com/tronprotocol/grpc-gateway/api";` by `option go_package = "github.com/streamingfast/tron-protocol/pb/api;pbtronapi";`.
- Replaced `option go_package = "github.com/tronprotocol/grpc-gateway/core";` by `option go_package = "github.com/streamingfast/tron-protocol/pb/core;pbtron";`.
- Moved `core/contract/*.proto` into `core/` directly to avoid Go import cycles (and adjusted Proto `import` directives).
- Removed empty `core/tron/*.proto` files.
- Generated and commit Go bindings at [pb](./pb/).

Check the [original README.md file](https://github.com/tronprotocol/protocol#protocol-) for details about the project