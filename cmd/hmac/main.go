package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"hash"
	"io/ioutil"
	"log"
	"os"

	"github.com/mas9612/cryptostudy/pkg/hmac"
	"github.com/mas9612/cryptostudy/pkg/util"
)

func main() {
	key := flag.String("key", "", "Secret key to calculate HMAC. Specify as hex notation without preceding \"0x\".")
	hashAlgo := flag.String("algorithm", "MD5", "Hash algorithm used to calculate HMAC. Default is \"MD5\".")
	flag.Parse()
	if *key == "" {
		log.Fatalln("-key is required")
	}

	var h hash.Hash
	switch *hashAlgo {
	case "MD5":
		h = md5.New()
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	keyBytes := util.HexStringToBytes(*key)
	fmt.Printf("%x\n", hmac.Hmac(h, keyBytes, len(keyBytes), data))
}
