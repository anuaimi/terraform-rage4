# configure provider
provider "rage4" {
  "email" = "whatever_used_to_login_to_rage4"
  "api_key" = "get_account_page_on_webui"
}

# add www record to a domain
resource "rage4_record" "www" {
  "zone_id" = ""
  "name" = ""
  "content" = ""
  "type" = ""
  "priority" = 1
  "active" = "true"

}
