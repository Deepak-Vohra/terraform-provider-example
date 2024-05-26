// resource_server.go
package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	environment_id := d.Get("environment_id").(string)

	d.SetId(environment_id)
	client := resty.New()
	resp, err := client.R().SetHeader("Accept", "application/json").SetHeader("Content-Type", "application/json").SetBody("{'repository': 'git@github.com:Deepak-Vohra/greetings.git@master','profile': 'Greetings Project','files': [{'mode': 0600,'path': 'config.json','contents': '{}'}]}").SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").Post("https://ca-1.platform.sh/api/projects/zis3mqzwuqnu4/environments/" + environment_id + "/initialize")
	if err != nil {
		log.Fatal(err)
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

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	environment_id := d.Get("environment_id").(string)

	d.SetId(environment_id)
	client := resty.New()
	resp, err := client.R().SetHeader("Accept", "application/json").SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").Get("https://ca-1.platform.sh/api/projects/zis3mqzwuqnu4/environments/" + environment_id)
	if err != nil {
		log.Fatal(err)
	}
	//defer resp.Body.Close()
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

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {

	environment_id := d.Get("environment_id").(string)

	d.SetId(environment_id)
	client := resty.New()
	resp, err := client.R().SetHeader("Accept", "application/json").SetHeader("Content-Type", "application/json").SetBody("{'name': 'updated env','title': 'updated env'}").SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").Patch("https://ca-1.platform.sh/api/projects/zis3mqzwuqnu4/environments/" + environment_id)
	if err != nil {
		log.Fatal(err)
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

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	environment_id := d.Get("environment_id").(string)

	d.SetId(environment_id)
	client := resty.New()
	resp, err := client.R().SetHeader("Accept", "application/json").SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").Delete("https://ca-1.platform.sh/api/projects/zis3mqzwuqnu4/environments/" + environment_id)
	if err != nil {
		log.Fatal(err)
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
