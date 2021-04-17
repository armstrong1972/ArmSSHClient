package main

import (
	cpt "armstrong/arm_crypt"
	"fmt"
	"os"
)

func main() {

	//fmt.Println(len(os.Args), os.Args)

	if len(os.Args) < 2 {
		fmt.Println("> shpwd yourpassword")
		os.Exit(1)
	}
	sPubKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwIc4VUrADMGHk7ywstSENsswvAsau0q8Rui+zbPC6OsyAcXAo/MFBlZv+G5YSz1yQqC/b9asMs2Y3QO9grhdPOdfhgkw14Ok/IGTewFKEEyO4evdeRtKlPcjiRP/RVwiAa6hoFlfEDkQlUu5h0taZz7yKkkAL33CC32vPTJCj/U+7Y1YPpGUnbzkZVE0Q2wpu9i73niM9nwhcrm2XR7K6Hn3uXV5UcmuBEELw7ZfXl2RsQmFHlNTbO8n7aTXLQAReOi/JDEcRoWkfazdqsEA8YppVbWue4wqth42SzEcgR7Otv8aVvxNDl4fO2kzZjfx/jiL/EaZdFfpHwLZOEFb7wIDAQAB"
	PubKey := cpt.Str2Buff(sPubKey)

	sText := os.Args[1]
	encBuff := cpt.RSA_Encrypt([]byte(sText), PubKey)
	sEnc := cpt.Bytes2Str(encBuff)

	fmt.Println("[Cipher]:")
	fmt.Println(sEnc)
	fmt.Println()
}
