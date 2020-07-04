# GoDotEnvSec

The GoDotEnvSec package loads env vars from an obfuscated .eenv file. It will also help to create an .eenv file and update it later when something changes

# Install

go get -u github.com/lesichkovm/godotenvsec

## Usage

```go
package main

import (
    "github.com/lesichkovm/godotenvsec"
)

func main() {
  godotenvsec.Init()

  s3Bucket := os.Getenv("S3_BUCKET")
  secretKey := os.Getenv("SECRET_KEY")
}

```

## Generating .eenv File From .env

- Use to obfuscate your .env file

```
go run main.go -envenc yes
```

Result:
```
==================================
== START: Encoding .env file    ==
==================================
1. Reading .env file...
2. Encoding contents with random key ...
3. Writing encoded content to .eenv file...
==================================
== END: Encoding .env file      ==
==================================
```

## Generating .denv File From .eenv

- Use to deobfuscate your .eenv file so that you can update the content

```
go run main.go -envdec yes
```

Result:
```
==================================
== START: Decoding .eenv file   ==
==================================
1. Reading .eenv file...
2. Decoding contents ...
3. Writing decoded content to .denv file...
==================================
== END: Decoding .eenv file     ==
==================================
```
