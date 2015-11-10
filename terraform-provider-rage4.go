package main

import (
  "errors"
  "os"
  "github.com/hashicorp/terraform/plugin"
  "github.com/hashicorp/terraform/helper/schema"
  "github.com/hashicorp/terraform/terraform"
  "github.com/anuaimi/rage4"
)

// see if environment variable set for given variable
func envDefaultFunc(k string) schema.SchemaDefaultFunc {
  return func() (interface{}, error) {
    if v := os.Getenv(k); v != "" {
      if v == "true" {
        return true, nil
      } else if v == "false" {
        return false, nil
      }
      return v, nil
    }
    return nil, nil
  }
}

// create client to speak to Rage4
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
  var err error
  
  // get email address to use
  email := d.Get("email").(string)
  if (len(email) == 0) {
    return nil, errors.New("rage4 email address not specified")
  }

  // get api key to use
  apiKey := d.Get("api_key").(string)
  if (len(apiKey) == 0) {
    return nil, errors.New("rage4 api key not specified")
  }

  // create rage4 client to return to terraform
  client, err := rage4.NewClient( email, apiKey)

  return client, err
}

func main() {

  // return details on provider plugin
  plugin.Serve(&plugin.ServeOpts {
    ProviderFunc: func() terraform.ResourceProvider {
    
      return &schema.Provider{
    
        // rage4 api only needs email login & api key
        Schema: map[string]*schema.Schema{
          "email": &schema.Schema{
            Type:        schema.TypeString,
            Required:    true,
            DefaultFunc: envDefaultFunc("RAGE4_EMAIL"),
            Description: "email address associated with Rage4 account",
          },
          "api_key": &schema.Schema{
            Type:        schema.TypeString,
            Required:    true,
            DefaultFunc: envDefaultFunc("RAGE4_API_AKEY"),
            Description: "Rage4 API Key",
          },
        },

        ConfigureFunc: providerConfigure,
      }

    },
  })

}

