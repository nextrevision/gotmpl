# gotmpl [![Build Status](https://travis-ci.org/nextrevision/gotmpl.svg?branch=master)](https://travis-ci.org/nextrevision/gotmpl)

CLI tool for rendering templates that supports environment variables, variable files, and inline encryption

## Installation

Download a pre-built binary from the releases (see [Releases](https://github.com/nextrevision/gotmpl/releases)):

```
wget https://github.com/nextrevision/gotmpl/releases/download/0.1.0/gotmpl_darwin_amd64
chmod +x gotmpl_darwin_amd64
```

Or install with go:

```
go get -u github.com/nextrevision/gotmpl
```

## Usage

```
gotmpl is a template tool that supports encrypted data

Usage:
  gotmpl [flags]
  gotmpl [command]

Available Commands:
  decrypt     decrypts vars or files
  encrypt     takes a plain text value and encrypts it
  genpasswd   generates a compliant 32 character password
  help        Help about any command
  render      renders a template file

Flags:
  -h, --help   help for gotmpl

Use "gotmpl [command] --help" for more information about a command.
```

### encrypt

Encrypt the string "mysecret" to stdout using password "abcdefghijklmnopqrstuvwxyz012345":

```
$ gotmpl encrypt -p abcdefghijklmnopqrstuvwxyz012345 -v mysecret
MzY1YTYwODUzMTRiZTE0YWVhYzJiZDk4OjA1M2JlNmJjNzNlMWYyY2QwYzg4YjNhYjU3OTkyYTZiZDM1MzA1MjcwZGVjNzc1NA==
```

Encrypt the string "mysecret" to stdout to be added to a YAML file with key "mykey":

```
$ gotmpl encrypt -p abcdefghijklmnopqrstuvwxyz012345 -v mysecret -k mykey
mykey: ENC|N2E3NTJjNDI2NmViMTRjMjZhMWIxNmI2OmVjYTNjYjdmN2ZhYTVmMzk0ZDVhMjUxZGQ3YzNiMTIzYzRiMTE2ZTdlNTM1M2M3ZA==
```

Encrypting and inserting the result key into a file:

```
$ gotmpl encrypt -p abcdefghijklmnopqrstuvwxyz012345 -v mysecret -k mykey -y examples/vars.yml
Variable mykey inserted into examples/vars.yml
```

### decrypt

Decrypting a single value to STDOUT:

```
$ gotmpl decrypt -p abcdefghijklmnopqrstuvwxyz012345 -v N2E3NTJjNDI2NmViMTRjMjZhMWIxNmI2OmVjYTNjYjdmN2ZhYTVmMzk0ZDVhMjUxZGQ3YzNiMTIzYzRiMTE2ZTdlNTM1M2M3ZA==
mysecret
```

Decrypting a vars file containing encrypted values:

```
$ gotmpl decrypt -p abcdefghijklmnopqrstuvwxyz012345 -y examples/vars.yml
File decrypted to examples/vars.yml.unenc
```

### render

Render a template to STDOUT sourcing vars from the environment:

```
$ gotmpl render -t examples/template.env.tmpl
# Static Key
Static value
# Env Var
/usr/local/bin/bash
```

Render a template to STDOUT sourcing vars from the environment and a vars file w/ encrytped values:

```
$ gotmpl render -t examples/template.tmpl -y examples/vars.yml -p abcdefghijklmnopqrstuvwxyz012345
# Static Key
Static value
# Key1
value1
# Encrypted Key
encValue1
# Inline Encryped Key
# password: abcdefghijklmnopqrstuvwxyz012345
encInline
# Env Var
/usr/local/bin/bash
```

Render a template to a specific file:

```
$ gotmpl render -t examples/template.env.tmpl -o examples/template.env
```

### genpasswd

Generating a new 32 character password:

```
$ gotmpl genpasswd
myvN3Uno4IcXtPoa4gjlvPIXsLg20K2G
```

## Vars Files

Vars files contain key/value pairs of variables used when rendering templates. These files can be in YAML or env (key=value) format. There is no support for nested keys, the files must be in a flat hierarchy.

Values can be encrypted (see "Encrypting Values" section), but must be prefixed with `ENC|` in order to be decrypted by gotmpl.

## Encrypting Values

### In vars files (YAML or env)
When working with encrypted values, anything that is used in a vars file (YAML or ENV) must be prefixed with `ENC|`. This distinction lets gotmpl know when processing a vars file to decrypt that value. For example:

```
---
key1: value1
encKey1: ENC|N2E3NTJjNDI2NmViMTRjMjZhMWIxNmI2OmVjYTNjYjdmN2ZhYTVmMzk0ZDVhMjUxZGQ3YzNiMTIzYzRiMTE2ZTdlNTM1M2M3ZA==
```

### In templates
You can supply an encrypted value inline in a template by prefixing with the encrypted string with `ENC`. For example, if I wanted to specify an encrypted value inline in a template:

```
# Plain text key sourced from environment or vars file
{{ .key1 }}

# Encrypted key in a vars file or environment
# This will be decrypted before the template is rendered
{{ .encKey1 }}

# Inline encrypted key
# This will be decrypted at the time of the template render
{{ ENC "MzY1YTYwODUzMTRiZTE0YWVhYzJiZDk4OjA1M2JlNmJjNzNlMWYyY2QwYzg4YjNhYjU3OTkyYTZiZDM1MzA1MjcwZGVjNzc1NA==" }}
```

### As an environment variable
You can supply an encrypted key in an environment variable so long as it is prefixed with `ENC|`, same as in the vars files. For example:

```
# plaintext key
export key1=value1

# encrypted key
export encKey1=ENC|N2E3NTJjNDI2NmViMTRjMjZhMWIxNmI2OmVjYTNjYjdmN2ZhYTVmMzk0ZDVhMjUxZGQ3YzNiMTIzYzRiMTE2ZTdlNTM1M2M3ZA==
```

## Developing

Pull requests and issues are more than welcome. Clone this repo, then download dependencies:

```
go get -u github.com/kardianos/govendor
govendor sync
```

To run tests:

```
govendor test -v +l
```
