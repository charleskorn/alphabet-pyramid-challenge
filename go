#!/usr/bin/env bash

set -e
set -u

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
  echo " setup        checks prerequisites are installed and fetches external dependencies"
  echo " build        builds the application"
  echo " test         runs the test suite"
  echo " run <args>   run the application (<args> are passed through to the application, must build first)"
}

function setup {
  VERSION_STRING=$(go version)

  if [[ "$VERSION_STRING" != "go version go$GO_VERSION "* ]]; then
    echo "Go version $GO_VERSION not found. (Try running 'brew install go'.)"
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

function run {
  if [ ! -f ./alphabet-pyramid-challenge ]; then
    echo "Application not found. Have you built it yet? (Try running './go build'.)"
  fi

  ./alphabet-pyramid-challenge "$@"
}

main "$@"
