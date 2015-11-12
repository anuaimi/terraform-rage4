package main

import (
  "fmt"
  "log"
  "strconv"
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

  // get id of domain we are going to add server to
  domainId, err := strconv.Atoi(d.Get("domainId").(string))

  // create new A record
  newRecord := new(rage4.Record)
  newRecord.Name = d.Get("name").(string)
  newRecord.Content = d.Get("content").(string)
  newRecord.Type = "A"
  // newRecord.TTL = d.Get("ttl")

  log.Printf("[DEBUG] Rage4 Record create configuration: %#v", newRecord)

  // ask Rage4 to add A record to domain
  client := meta.(*rage4.Client)
  status, err := client.CreateRecord( domainId, *newRecord)
  if (err != nil) {
    return fmt.Errorf("Failed to create Rage4 Record: %s", err)
  }

  // need to extract record id & save
  recId := status.Id
  d.SetId(string(recId))
  log.Printf("[INFO] created record ID: %s", d.Id())

  return nil
}

func resourceRage4RecordRead(d *schema.ResourceData, meta interface{}) error {
  // get which record we want to get details on
  recordId, err := strconv.Atoi(d.Id())
  domainId, err := strconv.Atoi(d.Get("domainId").(string))

  log.Printf("[INFO] Reading Rage4 Record %d in domain %s", recordId, domainId)

  // get details of records from Rage4
  client := meta.(*rage4.Client)
  _, err = client.GetRecords( domainId)

  return err
}

func resourceRage4RecordUpdate(d *schema.ResourceData, meta interface{}) error {
  // get which record we want to update
  recordId, err := strconv.Atoi(d.Id())
  domainId, err := strconv.Atoi(d.Get("domainId").(string))

  log.Printf("[INFO] Updating Rage4 Record %d in domain %s", recordId, domainId)

  record := new(rage4.Record)
  record.Name = d.Get("name").(string)
  record.Content = d.Get("content").(string)
  record.Type = "A"

  client := meta.(*rage4.Client)
  _, err = client.UpdateRecord( recordId, *record)

  return err
}

// delete given resource record for server
func resourceRage4RecordDelete(d *schema.ResourceData, meta interface{}) error {
  // get which domain we're working on
  recordId, err := strconv.Atoi(d.Id())
  domainId := d.Get("domainId")

  log.Printf("[INFO] Deleting Rage4 Record %d in domain %s", recordId, domainId)

  // ask rage4 to delete record
  client := meta.(*rage4.Client)
  _, err = client.DeleteRecord( recordId)

  return err
}
