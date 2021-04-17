package main

import (
	cpt "armstrong/arm_crypt"
	"fmt"
)

func main() {

	Private_Key, Public_Key := cpt.GenerateRSAKey(2048)
	sPri := cpt.Bytes2Str(Private_Key)
	sPub := cpt.Bytes2Str(Public_Key)

	fmt.Println("Public Key :")
	fmt.Println(sPub)
	fmt.Println("\nPrivate Key :")
	fmt.Println(sPri)
}
