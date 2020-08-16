# OkraGo
[Okra](https://okra.ng/) API wrapper in Go

<img src="https://pbs.twimg.com/profile_images/1199677745262989314/_D2jAMbu_400x400.jpg" alt="drawing" width="200"/> 

## Usage

### Install Package

```bash
go get github.com/Uchencho/OkraGo
```

Documentation
-------------
Please see https://docs.okra.ng/ for the most up-to-date documentation for the OKRA API.


#### OKRA

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
    okraClient, err := okra.New(token, sandboxUrl)
    
    if err != nil {
    	fmt.Println(err)
    }
    
    //okraClient returns an error if token or sandboxUrl is empty
    
    body, err2 := okraClient.RetrieveTransaction()
    	
    if err != nil {
        log.Println(err2)
    }
    fmt.Println(body.Data,"\n\n", body.StatusCode)
    
    /*
    okraClient has access to all the products OKRA API offers
    */
}
```
Use Product Name to Access Underlying Resources... ![](https://github.com/Uchencho/OkraGo/blob/master/client/okra.gif) 

>**NOTE**<br/>
>Check the `client` directory to see a sample implementation
