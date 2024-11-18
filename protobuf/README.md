# protobuf

### How to install:
- Install "protocol buffers"
    ```bash
    https://github.com/protocolbuffers/protobuf
    ```
    or

    (MacOS)
    ```bash
    brew install protobuf
    ```

- Install plugins to generate .go
    ```bash    
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```

### How to upgrade (MacOS):
    brew upgrade protobuf
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    
### How to run:
- Add these 2 important lines to your ~/.bash_profile or ~/.zshrc:
    ```bash
    export GO_PATH=~/go
    export PATH=$PATH:/$GO_PATH/bin
    ```
- Then close your terminal and re-open it, after that run:
    ```bash
    source ~/.bash_profile (source ~/.zshrc)
    ```
- Run:
    Single file
    ```bash
    protoc -I. --go_out=plugins=grpc:. your_file_want_to_build.proto
    ```
    Build all
    ```bash
    make
    ```

### How to test:
- Install prototool (https://github.com/uber/prototool#installation) (MacOS)
    ```bash
    brew install prototool
    prototool lint
    ```

### How to format .proto file(MacOS):
- Install clang-format
    ```bash
    brew install clang-format
    ```
- Generate .clang-format file
    ```bash
    clang-format -assume-filename=<any file *.proto you want> -style=llvm -dump-config > .clang-format
    ```
    EX: clang-format -assume-filename=tnx.proto -style=llvm -dump-config > .clang-format
- Config clang format:
    Add or update below line to .clang-format which had been generated prev step.
    ```bash
    AlignConsecutiveAssignments: true
    AlignConsecutiveDeclarations: true
    ```
- Add the following to your settings.json file:
    ```bash
    "[proto3]": {
        "editor.formatOnSave": true,
    },
    ```
- Double save(Cmd + S) to see the magic