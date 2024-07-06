#! /bin/bash

rm -rf ./bin

build_binary()
{
  GOOS=$(go env GOOS)
  GOARCH=$(go env GOARCH)
  BIN=$GOOS-$GOARCH

  if (( "$GOARCH" == "windows" )); then
    BIN="$BIN.exe"
  fi

  GOOS=$GOOS GOARCH=$GOARCH go build -o ./bin/"$BIN" ./cmd/main.go
  echo -e "build binary file \e[1;92mSUCCESSFULLY\e[0m"
}

build_binary
