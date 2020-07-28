# bbolt-reader
read "bbolt" contents


# bbolt

https://github.com/etcd-io/bbolt

# Installation

```bash
$ git clone git@github.com:otajisan/bbolt-reader.git
$ cd cmd/bbolt-reader
$ go build
$ go install
```

# CLI Usage

## basic

```bash
$ bbolt-reader -h
  NAME:
     bbolt-reader - get bbolt value by specified key
  
  USAGE:
     bbolt-reader [global options] command [command options] [arguments...]
  
  VERSION:
     0.0.1
  
  COMMANDS:
     get, g, read  bbolt-reader get --path=<bbolt db path> --bucket=<bucket> --key=<key>
     help, h       Shows a list of commands or help for one command
  
  GLOBAL OPTIONS:
     --help, -h     show help (default: false)
     --version, -v  print the version (default: false)
```

## get

```bash
$ bbolt-reader get -h
NAME:
   bbolt-reader get - bbolt-reader get --path=<bbolt db path> --bucket=<bucket> --key=<key>

USAGE:
   bbolt-reader get [command options] key

OPTIONS:
   --path value, -p value    bbolt db file path
   --bucket value, -b value  bbolt bucket name
   --key value, -k value     bbolt key
   --help, -h                show help (default: false)
```

## example

```bash
$ bbolt-reader get --path=/tmp/foo.db --bucket=mybucket --key=example | jq
{
  "title": "example",
  "detail": "this is example command."
}
```