# Substreams Entity Change

> Developer preview

`substreams-entity-change` crate contains all the definitions for the entity changes which can be emitted by a substream.

## Build proto manually

### Install protoc

To be able to build proto files manually using the command line, you have to have protoc installed on your machine. Visit [here](https://grpc.io/docs/protoc-installation/) to install.

For Linux, using apt
```bash
apt install -y protobuf-compiler
protoc --version  # Ensure compiler version is 3+
```

For macOS, using Homebrew
```bash
brew install protobuf
protoc --version  # Ensure compiler version is 3+
```

Instead of having a `build.rs` file which will build the proto automatically on every `cargo build --release` you have to build the proto files manually.

Simply run `./gen/generate.sh`
