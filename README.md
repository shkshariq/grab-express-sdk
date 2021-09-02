# GrabExpress SDK for Go [UNOFFICIAL]

An **UNOFFICIAL** Go SDK for integrating with the GrabExpress. 
Tested with Go 1.12+

⚠️ WARNING: This SDK is **NOT yet official**. What does this mean?
-   There is no formal Grab for this SDK at this point
-   Bugs may or may not get fixed
-   Not all SDK features may be implemented and implemented features may be buggy or incorrect
---

## Usage
First, you need to  [register](https://developer-beta.stg-myteksi.com/products/express-api)  to start making API requests. Once you have created an app, you can either use the SDK via an access token (useful for testing) or via the regular OAuth2 flow (recommended for production).

### Using OAuth token
Once you've created an app, you can get an access token from the [auth api](https://developer-beta.stg-myteksi.com/docs/grab-express/#authentication-method). Note that this token will only work for the GrabExpress account the token is associated with.

```golang
package main  
  
import (  
   ge "github.com/shkshariq/grab-express-sdk/grab-express"  
   "fmt"
)  
  
//https://developer-beta.stg-myteksi.com/docs/grab-express/#authentication-method  
var token = `GRAB_AUTH_TOKEN`
var isProduction = fasle

func main() {  
  
   client := ge.NewClient(isProduction, token)  
  
   // Preparing Request  
  req := ge.GrabExpressCancelDelivery{}  
   req.DeliveryID = "IN-2-00D2B1CBMWBO7JND48J5"  
  
  res, err := client.CancelDelivery(req)  
   if err != nil {  
      return  
  }  
  fmt.Println("RESULT", res)
}
```

