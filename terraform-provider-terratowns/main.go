// `package main`  This line indicates that this Go file belongs to the "main" package. In Go, the main package is special and is used for creating executable programs.
package main

import (
  "log"
  "fmt"
  "context"
  "github.com/google/uuid"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
  "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)
// `import "fmt"` This line indicates that this Go file belongs to the "main" package. In Go, the main package is special and is used for creating executable programs.

func main() {
       plugin.Serve(&plugin.ServeOpts{
	          ProviderFunc: Provider,

})
}	

type Config struct {
	Endpoint string
	Token string
	UserUuid string
}

func Provider() *schema.Provider {
var p *schema.Provider
p = &schema.Provider{
  ResourcesMap: map[string]*schema.Resource{
	"terratowns_home": Resource(),
	

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

p.ConfigureContextFunc = providerConfigure(p)
return p
}



func validateUUID(v interface{}, k string) (ws []string, errors []error) {
	log.Print("validateUUID:start")
	value := v.(string)
	if _, err := uuid.Parse(value); err != nil {
		errors = append(errors, fmt.Errorf("invalid UUID format"))
	}
	log.Print("validateUUID:end")
	return

}

func providerConfigure(p *schema.Provider) schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics ) {
		log.Print("providerConfigure:start")
		config := Config{
			Endpoint: d.Get("endpoint").(string),
			Token: d.Get("token").(string),
			UserUuid: d.Get("user_uuid").(string),
		}
		log.Print("providerConfigure:end")
		return &config, nil
	}
}

func Resource() *schema.Resource {
	log.Print("Resource:start")
	resource := &schema.Resource{
		CreateContext: resourceHouseCreate,
		ReadContext: resourceHouseRead,
		UpdateContext: resourceHouseUpdate,
		DeleteContext: resourceHouseDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: true,
				Description: "Name of home",
			},
			"description": {
				Type: schema.TypeString,
				Required: true,
				Description: "Description of home",
			},
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
				Description: "Domain name of home eg. *.cloudfront.net",
			},
			"town": {
				Type: schema.TypeString,
				Required: true,
				Description: "The town to which the home will belong to",
			},
			"content_version": {
				Type: schema.TypeInt,
				Required: true,
				Description: "The content version of the home",
			},
		},
	}
	log.Print("Resource:start")
	return resource

	func resourceHouseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		log.Print("resourceHouseCreate:start")
		var diags diag.Diagnostics
	
		config := m.(*Config)
	
		payload := map[string]interface{}{
			"name":           d.Get("name").(string),
			"description":    d.Get("description").(string),
			"domain_name":    d.Get("domain_name").(string),
			"town":           d.Get("town").(string),
			"content_version": d.Get("content_version").(int),
		}
	
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return diag.FromErr(err)
		}
	
		// You should return diags here or handle the creation logic.
		// If everything went well, you can return diags with no issues.
		return diags
	}
	