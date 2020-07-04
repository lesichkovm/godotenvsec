# GoDotEnvSec

The GoDotEnvSec package creates and loads env vars from an obfuscated .eenv file.

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
