# cloudlinux-go

# Installation
```
go get github.com/umtaktpe/cloudlinux-go
```

# Usage

To bind IP license with particular product you must provide valid product type:

- 1 - "CloudLinux OS"
- 4 – "CloudLinux Solo"
- 16 - "KernelCare"
- 17 – "KernelCare Plus"
- 40 - "ImunifyAV+"
- 41 - "Imunify360 single user"
- 42 - "Imunify360 up to 30 users"
- 43 - "Imunify360 up to 250 users"
- 49 - "Imunify360 unlimited"

```go
package main

import (
	"fmt"
	"log"

	"github.com/umtaktpe/cloudlinux-go"
	"github.com/umtaktpe/cloudlinux-go/client"
	"github.com/umtaktpe/cloudlinux-go/model"
)

func main() {
	config := cloudlinux.NewConfig("username", "secret_key")
	client := client.NewClient(config)

	params := &model.LicenseRegisterParams{
		Ip: "1.1.1.1",
		Type: "16",
	}

	response, err := client.LicenseRegister(params)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(response)
}
```