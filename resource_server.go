// resource_server.go
package main

import (
	"fmt"
	"log"

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

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	environment_id := d.Get("environment_id").(string)

	client := resty.New()
	resp, err := client.R().SetHeader("Accept", "application/json").SetHeader("Content-Type", "application/json").SetBody("{'repository': 'git@github.com:Deepak-Vohra/greetings.git@master','profile': 'Greetings Project','files': [{'mode': 0600,'path': 'config.json','contents': '{}'}]}").SetAuthToken("eyJhbGciOi").Post("https://ca-1.platform.sh/api/projects/zis3mqzwuqnu4/environments/" + environment_id + "/initialize")
	if err != nil {
		log.Fatal(err)
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

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	environment_id := d.Id()
	client := resty.New()
	resp, err := client.R().SetHeader("Accept", "application/json").SetAuthToken("eyJhbGciOiJFUzI").Get("https://ca-1.platform.sh/api/projects/zis3mqzwuqnu4/environments/" + environment_id)
	if err != nil {
		log.Fatal(err)
	}
	//defer resp.Body.Close()
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

	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {

	environment_id := d.Get("environment_id").(string)

	d.SetId(environment_id)
	client := resty.New()
	resp, err := client.R().SetHeader("Accept", "application/json").SetHeader("Content-Type", "application/json").SetBody("{'name': 'updated env','title': 'updated env'}").SetAuthToken("eyJhbGciOiJFUzI1NiIsImtpZCI6ImU2ODAyNTc5NTVlOWFmYTUzZmRlY2U5YWNmYTMzMDYzZjEzYTYxZmI2OWRmMTg3YTI0NDdkYjFjOGQyMmRjMzAiLCJ0eXAiOiJKV1QifQ.eyJhY2Nlc3NfaWQiOiJMT01MTlZRNVZPN1JFU0xERUkyMzZEVUJZTFFTUlpJVFJPNDdCRjNVVTJXMzIiLCJhY3QiOnt9LCJjaWQiOiJwbGF0Zm9ybS1hcGktdXNlciIsImV4cCI6MTcxNjY5NzI4NiwiZ3JhbnQiOiJhcGlfdG9rZW4iLCJpYXQiOjE3MTY2OTYzODYsImlzcyI6Imh0dHBzOi8vYXV0aC5hcGkucGxhdGZvcm0uc2giLCJqdGkiOiIwMUhZU0g4MzA4OERLWlRLQ1g1UlpFWU1BSCIsIm5iZiI6MTcxNjY5NjM4NiwibnMiOiJwbGF0Zm9ybXNoIiwic3ViIjoiNzZjNjBlNDMtMDkzOC00NjI5LWFlMDAtY2IwYTU5YTQ0YzIwIn0.wXaw8BbQ3ufJoThpW-xKYy0a3P4ds8SdWAOCFBBb-LuZAR9GcUBa5HTcrWhkV13-j6jetnx8bLeYRzgBvrygFg").Patch("https://ca-1.platform.sh/api/projects/zis3mqzwuqnu4/environments/" + environment_id)
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

	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	environment_id := d.Id()

	d.SetId("")
	client := resty.New()
	resp, err := client.R().SetHeader("Accept", "application/json").SetAuthToken("eyJhbGciOiJFUzI1NiIsImtpZCI6ImU2ODAyNTc5NTVlOWFmYTUzZmRlY2U5YWNmYTMzMDYzZjEzYTYxZmI2OWRmMTg3YTI0NDdkYjFjOGQyMmRjMzAiLCJ0eXAiOiJKV1QifQ.eyJhY2Nlc3NfaWQiOiJMT01MTlZRNVZPN1JFU0xERUkyMzZEVUJZTFFTUlpJVFJPNDdCRjNVVTJXMzIiLCJhY3QiOnt9LCJjaWQiOiJwbGF0Zm9ybS1hcGktdXNlciIsImV4cCI6MTcxNjY5NzI4NiwiZ3JhbnQiOiJhcGlfdG9rZW4iLCJpYXQiOjE3MTY2OTYzODYsImlzcyI6Imh0dHBzOi8vYXV0aC5hcGkucGxhdGZvcm0uc2giLCJqdGkiOiIwMUhZU0g4MzA4OERLWlRLQ1g1UlpFWU1BSCIsIm5iZiI6MTcxNjY5NjM4NiwibnMiOiJwbGF0Zm9ybXNoIiwic3ViIjoiNzZjNjBlNDMtMDkzOC00NjI5LWFlMDAtY2IwYTU5YTQ0YzIwIn0.wXaw8BbQ3ufJoThpW-xKYy0a3P4ds8SdWAOCFBBb-LuZAR9GcUBa5HTcrWhkV13-j6jetnx8bLeYRzgBvrygFg").Delete("https://ca-1.platform.sh/api/projects/zis3mqzwuqnu4/environments/" + environment_id)
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
