package challenge_13_alt

import (
	"fmt"
	chal10 "github.com/noliverio/friendly-doodle/set2/challenge_10"
	chal9 "github.com/noliverio/friendly-doodle/set2/challenge_9"
	utils "github.com/noliverio/friendly-doodle/utils"
)

var key = []byte("YELLOW SUBMARINE")

func Main() {
	enc := createAndEncryptProfile("abcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAroleAadminAAAAAAAAAAAAA@123.com")

	//Flip bits around the role to create the appropriate metacharacters.
	enc[32] = 247
	enc[43] = 53
	enc[37] = 124

	// Garble the the "role" field provided by the previous function
	enc[64] = 1
	dec := decryptAndParseProfile(enc)

	// My new role field is intact but old role field is garbled to junk.
	// The function recognizes me as role: admin.
	for key, value := range dec {
		fmt.Println()
		fmt.Println("key:")
		fmt.Println(key)
		fmt.Println("value:")
		fmt.Println(value)
	}
}

func createAndEncryptProfile(email string) []byte {
	profile, err := utils.ProfileFor(email)
	utils.Check(err)
	encryptedString := chal10.CbcEncrypt(chal9.PadMessage(profile, 16), initVector, key)
	return encryptedString
}

func decryptAndParseProfile(profileCipherText []byte) map[string]string {

	profile := chal10.CbcDecrypt(profileCipherText, initVector, key)

	return utils.ParseKV(profile)
}
