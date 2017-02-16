package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// EncodeCredentials this func is encoding the Username and Password with base64 encoding which is
// required for Nutanix
func EncodeCredentials(username string, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

// v0_8 returns the main entry point for the v0.8 Nutanix API
func v0_8(NutanixHost string) string {

	return "https://" + NutanixHost + ":9440/api/nutanix/v0.8/"

}

// v1_0 returns the main entry point for the v1.0 Nutanix API
func v1_0(NutanixHost string) string {

	return "https://" + NutanixHost + ":9440/PrismGateway/services/rest/v1/"

}

// v2_0 returns the main entry point for the v2.0 Nutanix API
func v2_0(NutanixHost string) string {

	return "https://" + NutanixHost + ":9440/PrismGateway/services/rest/v2.0/"

}

// v3_0 returns the main entry point for the v3.0 Nutanix API -> Not GA with AOS 5.0
func v3_0(NutanixHost string) string {

	return "https://" + NutanixHost + ":9440/PrismGateway/services/rest/v3.0/"

}

func main() {

	// PRISM user
	var username = "admin"
	// PRISM user password
	var password = "nutanix/4u"
	// Nutanix Cluster IP/DNSName CVM IP/DNSName
	var NutanixHost = "192.168.178.130"

	// Ignores certificates which can not be validated
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// create a HTTP client
	var httpClient = http.Client{Transport: tr}

	// create a http Request pointer
	var req *http.Request
	var err error

	// Defines the HTTP Request
	// send a GET to the NUTANIX API and receives the UUID of the VM "docker-mac"

	req, _ = http.NewRequest("GET", v2_0(NutanixHost)+"/vms/?include_vm_nic_config=true", nil)

	// before the request is send set the HTTP Header key "Authorization" with
	// the value of base64 encoded Username and Password
	req.Header.Set("Authorization", "Basic "+EncodeCredentials(username, password))

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// read the data from the resp.body into bodyText
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// create interface
	var f interface{}

	// Unmarshal into interface f
	if err2 := json.Unmarshal(bodyText, &f); err != nil {
		panic(err2)
	}

	// type assertion to access f´s underlying map[string]interface
	m := f.(map[string]interface{})

	// the response will include entities which includes the data of the VMs we searching for
	e := m["entities"].([]interface{})

	// we can iterate through the map and search for the IP of the VM
	for k := range e {
		t := e[k].(map[string]interface{})

		fmt.Println(t["name"])

		if t["vm_nics"] != nil {
			n := t["vm_nics"].([]interface{})

			for k2 := range n {
				t2 := n[k2].(map[string]interface{})

				fmt.Print("AHV managed: ")
				fmt.Println(t2["requested_ip_address"])

			}

		}

	}

	// Defines the HTTP Request
	// send a GET to the NUTANIX API and receives the details of all VMs

	req, _ = http.NewRequest("GET", v1_0(NutanixHost)+"/vms/", nil)

	// before the request is send set the HTTP Header key "Authorization" with
	// the value of base64 encoded Username and Password
	req.Header.Set("Authorization", "Basic "+EncodeCredentials(username, password))

	resp, err = httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// read the data from the resp.body into bodyText
	bodyText, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal into interface f
	if err2 := json.Unmarshal(bodyText, &f); err != nil {
		panic(err2)
	}

	m = f.(map[string]interface{})

	// the response will include entities which includes the data of the VMs we searching for
	e = m["entities"].([]interface{})

	// we can iterate through the map and search for the IP of the VM
	for k := range e {
		t := e[k].(map[string]interface{})

		fmt.Print("V1 Workaround: ")
		fmt.Println(t["ipAddresses"])
	}

}
