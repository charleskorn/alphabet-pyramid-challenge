#!/usr/bin/env bash

set -e

GO_VERSION="1.5.2"

function main {
  case "$1" in

  setup)
    setup
    ;;

  build)
    build
    ;;

  test)
    test
    ;;

  *)
    help
    exit 1
    ;;

  esac
}

function help {
  echo "Usage:"
  echo " setup  checks prerequisites are installed and fetches external dependencies"
  echo " build  compiles the application"
  echo " test   runs the test suite"
}

function setup {
  VERSION_STRING=$(go version)

  if [[ "$VERSION_STRING" != "go version go$GO_VERSION "* ]]; then
    echo "Go version $GO_VERSION not found. Try running 'brew install go'."
    exit 1
  fi

  go get -t -v
}

function build {
  go build
}

function test {
  go test
}

main $1
