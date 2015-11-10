package main

import (
  "fmt"
  "log"
  "github.com/anuaimi/rage4"
  "github.com/hashicorp/terraform/helper/schema"
)

func resourceRage4Record() *schema.Resource {
  return &schema.Resource{
    Create: resourceRage4RecordCreate,
    Read:   resourceRage4RecordRead,
    Update: resourceRage4RecordUpdate,
    Delete: resourceRage4RecordDelete,

    Schema: map[string]*schema.Schema{
      "domainId": &schema.Schema{
        Type:     schema.TypeInt,
        Required: true,
      },
      "name": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "content": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "type": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "priority": &schema.Schema{
        Type:     schema.TypeInt,
        Optional: true,
        Default:  1,
      },
      "active": &schema.Schema{
        Type:     schema.TypeBool,
        Optional: true,
        Default:  true,
      },
      "failover": &schema.Schema{
        Type:     schema.TypeBool,
        Optional: true,
        Default:  false,
      },
      "failovercontent": &schema.Schema{
        Type:     schema.TypeString,
        Optional: true,
      },
      "failoverwithdraw": &schema.Schema{
        Type:     schema.TypeString,
        Optional: true,
      },
      "ttl": &schema.Schema{
        Type:     schema.TypeString,
        Optional: true,
        Default:  "3600",
      },

    },
  }
}

func resourceRage4RecordCreate(d *schema.ResourceData, meta interface{}) error {

  domainId := d.Get("id").(int)
  newRecord := new(rage4.Record)
  // fill fields in record

  log.Printf("[DEBUG] Rage4 Record create configuration: %#v", newRecord)

  client := meta.(*rage4.Client)
  _, err := client.CreateRecord( domainId, *newRecord)
  if (err != nil) {
    return fmt.Errorf("Failed to create Rage4 Record: %s", err)
  }

  // need to extract record id & save
  recId := 1
  d.SetId(string(recId))
  log.Printf("[INFO] record ID: %s", d.Id())

  return nil
}

func resourceRage4RecordRead(d *schema.ResourceData, meta interface{}) error {
  // client := meta.(*rage4.Client)
  return nil
}

func resourceRage4RecordUpdate(d *schema.ResourceData, meta interface{}) error {
  // client := meta.(*rage4.Client)
  // get id of record to change, 
  // UpdateRecord( recordId int, record Record) (status Status, err error) {

  return nil
}

func resourceRage4RecordDelete(d *schema.ResourceData, meta interface{}) error {
  // get which domain we're working on
  domainId := d.Get("id").(int)
  recordId := 1

  log.Printf("[INFO] Deleting Rage4 Record for %d in domain %s", d.Get(""), domainId)

  // ask rage4 to delete record
  client := meta.(*rage4.Client)
  _,  err := client.DeleteRecord( recordId)

  return err
}
