# OkraGo
Okra API wrapper in Go

<img src="https://pbs.twimg.com/profile_images/1199677745262989314/_D2jAMbu_400x400.jpg" alt="drawing" width="200"/> 

## Usage

### Install Package

```bash
go get github.com/Uchencho/OkraGo
```

### OKRA

```go
package main

import (
  	"fmt"
	 "os"
    
	okra "github.com/Uchencho/OkraGo"
)

token := os.Getenv("OKRA_TOKEN")
const sandboxUrl = "https://api.okra.ng/sandbox/v1/"
```

* Initialize a client

```go

func main () {
    okraClient := okra.New(token, sandboxUrl)
    
    //okraClient panics if token or sandboxUrl is empty
    
    body, err := okraClient.RetrieveTransaction()
    	
    if err != nil {
        log.Println(err)
    }
    fmt.Println(body.Data,"\n\n", body.StatusCode)
    
    /*
    okraClient has access to all the products OKRA API offers
    */
}
```
GIF coming soon...

>**NOTE**<br/>
>Check the `client` directory to see a sample implementation

## Contributors

* Nils Wogatzky - [GitHub](https://github.com/Oppodelldog)
