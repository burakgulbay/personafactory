# personafactory

A simple persona factory including job, birthdate, gender with reasonable first name and last name using Go<br />
May use for mocking purposes<br />
I will keep updating by adding _email, address, social accounts,_ etc.

## Install in Your Project

```bash
go get github.com/burakgulbay/personafactory
```

## Update Existing Version in Your Project

```bash
go get -u github.com/burakgulbay/personafactory
```

## How to Use

```go
package main

import (
    "github.com/burakgulbay/personafactory"
)

func main() {
    personCreator := CreatePersonCreator()
    p := personCreator.Create()
    p.PrintJSON()
}
```