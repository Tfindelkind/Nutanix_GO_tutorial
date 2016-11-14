package main

import (
	"encoding/base64"
)

// EncodeCredentials this func is encoding the Username and Password with base64 encoding which is
// required for Nutanix
func EncodeCredentials(n *NTNXConnection) {
	n.SEnc = base64.StdEncoding.EncodeToString([]byte(n.Username + ":" + n.Password))
}

// NutanixAHVurl ...
func v0_8(n *NTNXConnection) string {

	return "https://" + n.NutanixHost + ":9440/api/nutanix/v0.8/"

}

// NutanixRestURL ...
func v1_0(n *NTNXConnection) string {

	return "https://" + n.NutanixHost + ":9440/PrismGateway/services/rest/v1/"

}



func func main() {




}
