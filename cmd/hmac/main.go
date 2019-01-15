package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/sha3"

	"github.com/mas9612/cryptostudy/pkg/hmac"
	"github.com/mas9612/cryptostudy/pkg/util"
)

func main() {
	key := flag.String("key", "", "Secret key to calculate HMAC. Specify as hex notation without preceding \"0x\".")
	hashAlgo := flag.String("algorithm", "MD5", "Hash algorithm used to calculate HMAC. Default is \"MD5\". Valid algorithm is one of [MD5, SHA-1, SHA-224, SHA-256, SHA-384, SHA-512]")
	flag.Parse()

	var h hash.Hash
	switch *hashAlgo {
	case "MD5":
		h = md5.New()
	case "SHA-1":
		h = sha1.New()
	case "SHA-224":
		h = sha256.New224()
	case "SHA-256":
		h = sha256.New()
	case "SHA-384":
		h = sha512.New384()
	case "SHA-512":
		h = sha512.New()
	case "SHA3-224":
		h = sha3.New224()
	case "SHA3-256":
		h = sha3.New256()
	case "SHA3-384":
		h = sha3.New384()
	case "SHA3-512":
		h = sha3.New512()
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	keyBytes := util.HexStringToBytes(*key)
	fmt.Printf("%x\n", hmac.Hmac(h, keyBytes, len(keyBytes), data))
}
