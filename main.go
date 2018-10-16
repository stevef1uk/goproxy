package main

import (
    //"fmt"
    "time"
    "log"

    	//"github.com/tylerb/graceful"
    	"goproxy/client/operations"
    	apiclient "goproxy/client"
      "github.com/go-openapi/swag"
)

func main() {

  myparams := operations.NewGetAccounts2Params()
  myparams.ID = *swag.Int32(1)
  myparams.WithTimeout(time.Second * 2)
  //fmt.Printf("I am doing it")

  for i:= 0; i < 10000; i++ {
    //resp, err := apiclient.Default.Operations.GetAccounts2(myparams)
    _, err := apiclient.Default.Operations.GetAccounts2(myparams)
    if err != nil {
      log.Fatal(err)
    }
    //fmt.Println(resp)

  }
}
