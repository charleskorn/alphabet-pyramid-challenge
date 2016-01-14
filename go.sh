#!/usr/bin/env bash

set -e
set -u

GO_VERSION="1.5.3"
BINARY_NAME="./alphabet-pyramid-challenge"

function main {
  case "$1" in

  check)
    check
    ;;

  setup)
    setup
    ;;

  build)
    build
    ;;

  test)
    test
    ;;

  run)
    run "${@:2}"
    ;;

  *)
    help
    exit 1
    ;;

  esac
}

function help {
  echo "Usage:"
  echo " check        checks prerequisites are installed"
  echo " setup        fetches external dependencies"
  echo " build        builds the application"
  echo " test         builds and runs the test suite"
  echo " run <args>   run the application (<args> are passed through to the application, must build first)"
}

function check {
  path_to_go=$(which go || echo "")
  if [ ! -x "$path_to_go" ]; then
    echo "Go not found. (Try running 'brew update && brew install go'.)"
    exit 1
  fi

  version_string=$(go version)

  if [[ "$version_string" != "go version go$GO_VERSION "* ]]; then
    echo "Go version $GO_VERSION not found."
    echo "If you have a more recent version of Go installed, you can ignore this message. (Run 'go version' to check.)"
    exit 1
  fi

  echo "Everything looks OK!"
}
function setup {
  go get -t -v
}

function build {
  go build -o "$BINARY_NAME"
}

function test {
  go test
}

function run {
  if [ ! -f "$BINARY_NAME" ]; then
    echo "Application not found. Have you built it yet? (Try running './go.sh build'.)"
  fi

  "$BINARY_NAME" "$@"
}

main "$@"
