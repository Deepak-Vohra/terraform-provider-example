// resource_server.go
package main

import (
	"fmt"
	"tflog"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"environment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
                       "name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			        Default:  "a platform env",
			},
		},
	}
}

func resourceServerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics  {
	environment_id := d.Get("environment_id").(string)

	client := resty.New()
	tflog.Info(ctx, "Using Platform.sh API token for authentication")
	resp, err := client.R().SetHeader("Accept", "application/json").SetHeader("Content-Type", "application/json").SetBody("{'repository': 'git@github.com:Deepak-Vohra/greetings.git@master','profile': 'Greetings Project','files': [{'mode': 0600,'path': 'config.json','contents': '{}'}]}").SetAuthToken("eyJhbGciOi").Post("https://ca-1.platform.sh/api/projects/zis3mqzwuqnu4/environments/" + environment_id + "/initialize")
	 if err != nil {
          // Convert a Go error to Diagnostics
          return diag.FromErr(err)
         }

	// Explore response object (optional)
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	//Obtain id of new environment from the response Body and set in the resource's state
	d.SetId("new environment id")
	
	return resourceServerRead(d, m)
}

func resourceServerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics  {
	environment_id := d.Id()
	client := resty.New()
	resp, err := client.R().SetHeader("Accept", "application/json").SetAuthToken("eyJhbGciOiJFUzI").Get("https://ca-1.platform.sh/api/projects/zis3mqzwuqnu4/environments/" + environment_id)
	
	// Collect Errors & Warnings in a slice type
  var diags diag.Diagnostics
	// Return formatted error  
    if err != nil {
        diags = append(diags, diag.Errorf("unexpected: %s", err)...)
	      return diags
    }
     
	
	 
	// Explore response object (optional)
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

      //  Map the response body to resource schema attributes
      // Update state
	return nil
}

func resourceServerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics  {
       //Implementation is optional
	environment_id := d.Get("environment_id").(string) //unique resource identifier
        name := d.Get("name").(string)
	 
	client := resty.New()
	resp, err := client.R().SetHeader("Accept", "application/json").SetHeader("Content-Type", "application/json").SetBody("{'name': 'updated env','title': 'updated env'}").SetAuthToken("eyJhbGciOiJFUzI").Patch("https://ca-1.platform.sh/api/projects/zis3mqzwuqnu4/environments/" + environment_id)
	if err != nil {
          // Convert a Go error to Diagnostics
          return diag.FromErr(err)
         }

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	//Call Read to update state with the updated resource
	return resourceServerRead(d, m)
}

func resourceServerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics  {
	environment_id := d.Id()

	d.SetId("")
	client := resty.New()
	resp, err := client.R().SetHeader("Accept", "application/json").SetAuthToken("eyJhbGciOi").Delete("https://ca-1.platform.sh/api/projects/zis3mqzwuqnu4/environments/" + environment_id)
	if err != nil {
          // Convert a Go error to Diagnostics
          return diag.FromErr(err)
         }

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	return nil
}
