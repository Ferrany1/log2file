# Golang Log2File

Module that initialize mainLogFile in a repository and renames previous one into backupLogFile. (Both names are set on creating options)

## Installation
```
$ GO111MODULE=on go get github.com/Ferrany1/log2file
```

```go
...
import (
  "github.com/Ferrany1/log2file" // imports as package "log2file"
)
...
```



##Examples:
1. [Empty Options](/examples/example1/example1_empty.go)
2. [Custom Options](/examples/example2/example2_custom.go)