package main

import (
	adx "ad/ad"
	"fmt"
)

func main() {
	cfg := &adx.Config{
		Server:   "172.30.137.212",
		Port:     "636",
		Security: adx.SecurityTLS,
		// User: "",
		// Password: "",
		User:     "pradist.k_ocptest@kbankpocnet.com",
		Password: "P@ssw0rd",
		BaseDN:   "DC=kbankpocnet,DC=com",
	}
	adClient, _, err := adx.NewClient("test", cfg)()
	if err != nil {
		panic(err)
	}

	_, _, err = adClient.Authenticate("boonyurai.h@kbankpocnet.com", "P@ssw9rd1")
	// _, _, err = adClient.Authenticate("phubtest01@kbankpocnet.com", "P@ssw0rd")
	// _, _, err = adClient.AuthenticateWithOutDomain("pradist.k_ocptest", "P@ssw0rd")

	// result, _ := adClient.FindInfo("sAMAccountName", "boonyurai.h")
	// fmt.Printf("%v", result)

	if err != nil {
		panic(err)
	}
	fmt.Println("login is ok.......!!!!!!!!!!")

	//user, err := adClient.FindInfo("sAMAccountName", "boonyurai.h")
	//user, err := adClient.FindInfo("sAMAccountName", "pradist.k_ocptest")
	//user, err := adClient.FindInfo("sAMAccountName", "phubtest01")

	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//
	//for key, value := range user {
	//	fmt.Println("Key:", key, "Value:", value)
	//}
}
