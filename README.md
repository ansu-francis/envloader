# envloader

A Go  project hich loads env vars from a .env file.

## Usage

Add your application configuration to your `config.env` file in the root of your project:

```shell
S3_BUCKET=YOURS3BUCKET
SECRET_KEY=YOURSECRETKEYGOESHERE
```

Then in your Go app you can do something like

```go
package main

import (
    "log"
    "os"

    "github.com/ansu-francis/envloader"
)

func main() {
  err := envloader.LoadEnv("config.env")
  if err != nil {
    log.Fatal("Error loading config.env file")
  }

  s3Bucket := os.Getenv("S3_BUCKET")
  secretKey := os.Getenv("SECRET_KEY")

  // now do something with s3 or whatever
}
```
