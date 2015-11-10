package main

import (
  // "github.com/anuaimi/rage4"
  "github.com/hashicorp/terraform/helper/schema"
)

func resourceRage4Record() *schema.Resource {
  return &schema.Resource{
    Create: resourceRage4RecordCreate,
    Read:   resourceRage4RecordRead,
    Update: resourceRage4RecordUpdate,
    Delete: resourceRage4RecordDelete,

    Schema: map[string]*schema.Schema{
      "id": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },


      "domain": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
        ForceNew: true,
      },


      "name": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },

      "hostname": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },

      "type": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
        ForceNew: true,
      },

      "value": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },

      "ttl": &schema.Schema{
        Type:     schema.TypeString,
        Optional: true,
        Default:  "3600",
      },

      "priority": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
      },
    },
  }
}

func resourceRage4RecordCreate(d *schema.ResourceData, meta interface{}) error {
  // client := meta.(*rage4.Client)
  return nil
}

func resourceRage4RecordRead(d *schema.ResourceData, meta interface{}) error {
  // client := meta.(*rage4.Client)
  return nil
}

func resourceRage4RecordUpdate(d *schema.ResourceData, meta interface{}) error {
  // client := meta.(*rage4.Client)
  return nil
}

func resourceRage4RecordDelete(d *schema.ResourceData, meta interface{}) error {
  // client := meta.(*rage4.Client)
  return nil
}
