# paste-master

Filtering Paste From StdIn.

## Description

processing paste from stdin

## Usage

```console
$ pbpaste | paste-master
```

### Options
#### RegexFilter

```console
$ pbpaste | paste-master --regex="^\d$"
```

#### Implode Processing
> LineEnd Replacer 
```console
$ pbpaste | paste-master --implode=", "
```


## Install

### homebrew
```console
$ brew tap aozora0000/homebrew-tap
$ brew install paste-master
or
$ brew install aozora0000/homebrew-tap/paste-master
```

## Contribution

1. Fork ([https://github.com/aozora0000/paste-master/fork](https://github.com/aozora0000/paste-master/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[aozora0000](https://github.com/aozora0000)
