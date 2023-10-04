// `package main`  This line indicates that this Go file belongs to the "main" package. In Go, the main package is special and is used for creating executable programs.
package main

import (
   //"log"
  //"fmt"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
  "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)
// `import "fmt"` This line indicates that this Go file belongs to the "main" package. In Go, the main package is special and is used for creating executable programs.

func main() {
       plugin.Serve(&plugin.ServeOpts{
	          ProviderFunc: Provider,

})
}	
func Provider() *schema.Provider {
var p *schema.Provider
p = &schema.Provider{
  ResourcesMap: map[string]*schema.Resource{

  },
  DataSourcesMap: map[string]*schema.Resource{

  },
  Schema: 	map[string]*schema.Schema{
	"endpoint": {
	  Type: schema.TypeString,
	  Required: true,
	  Description: "The endpoint for the external service",	
	},
	"token": {
	  Type: schema.TypeString,
	  Sensitive: true,
	  Required: true,
	  Description: "Bearer token for authorization",

	},
	"user_uuid": {
	  Type: schema.TypeString,
	  Required: true,
	  Description: "UUID for configuration",
	  //ValidateFunc: validateUUID,

	},

  },
}

//p.ConfigureContextFunc = providerConfigure(p)
return p
}



//func validateUUID(v interface{}, k string) (ws []string, errors []error) {
//	log.Print('validateUUID:start')
//    value, ok := v.(string)
//    if !ok {
//        errors = append(errors, schema.TypeErrorf("%q must be a string", k))
//        return
//    }
//
//    _, err := uuid.Parse(value)
//    if err != nil {
//        errors = append(errors, schema.ValidateErrorf("%q is not a valid UUID: %s", k, err))
//    }
//
//    return
// }
