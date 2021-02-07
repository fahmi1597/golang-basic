package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)
func encryptSalt(str string) ([]byte, string){
    salt := fmt.Sprintf("%d", time.Now().UnixNano())
    saltStr := fmt.Sprintf("%s%s", salt, str)
    
    var hash = sha1.New()
    hash.Write([]byte(saltStr))
    sha1Hash := hash.Sum(nil)
    return sha1Hash, salt
} 

func main() {
    var securePassword = "secretpassword"
    var hash = sha1.New()
    // encrypt tanpa salt
    hash.Write([]byte(securePassword))
    sha1Hash := hash.Sum(nil)  // "sha" is uint8 type, encoded in base16
    fmt.Printf("%x\n", sha1Hash)

    // dengan salt
    salt, hashWithSalt := encryptSalt(securePassword)
    fmt.Printf("Salt : %x \nHash : %s\n", salt, hashWithSalt)
}