# tempura
A Fast and Flexible Template Fill Tool.

[![](https://img.shields.io/badge/docker-ghcr.io%2Fkorosuke613%2Ftempura-blue)](https://github.com/users/korosuke613/packages/container/package/tempura) [![](https://img.shields.io/github/v/release/korosuke613/tempura)](https://github.com/korosuke613/tempura/releases) ![](https://img.shields.io/github/go-mod/go-version/korosuke613/tempura) 
## Install

### go install
```
go install github.com/korosuke613/tempura@latest
```

### docker
```
docker pull ghcr.io/korosuke613/tempura
```

### binary
Download from [Releases](https://github.com/korosuke613/tempura/releases).

## Getting Started
using binary
```
❯ tempura --input-string '{"Name": "John", "Message": "Good"}' --template-string 'Hello {{.Name}}, {{.Message}}'
Hello John, Good
```

or using docker

```
❯ docker run --rm ghcr.io/korosuke613/tempura --input-string '{"Name": "John", "Message": "Good"}' --template-string 'Hello {{.Name}}, {{.Message}}'
Hello John, Good
```


## Example

### input multiline template
```
❯ tempura \
 --input-string '{"Name": "John", "Message": "Good"}' \
 --template-string \
'Hello {{.Name}},
{{.Message}}'
Hello John,
Good
```

### read template and inputs

`input.json`
```json
{
  "Name": "John",
  "Message": "Good"
}
```

`template.txt`
```text
Hello {{.Name}},
{{.Message}}
```

```
❯ tempura -i input.json -t template.txt
Hello John,
Good
```

### output to file
```
❯ tempura -o ./output.txt
❯ cat ./output.txt
Hello John,
Good
```

### Options
```
❯ tempura -h
A Fast and Flexible Template Fill Generator built with love by korosuke613 in Go.

Usage:
  tempura [flags]

Flags:
  -h, --help                       help for tempura
  -i, --input-filepath string      input file name (default "input.json")
      --input-string string        input string
  -o, --output string              output file name
  -t, --template-filepath string   template file name (default "template.txt")
      --template-string string     template string
  -v, --version                    show version
```
