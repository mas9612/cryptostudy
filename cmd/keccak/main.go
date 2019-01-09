package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mas9612/cryptostudy/pkg/keccak"
)

func main() {
	len := flag.Int("d", 256, "The length of the digest of a hash function. Default is \"256\". Valid length is one of [224, 256, 384, 512]")
	flag.Parse()

	switch *len {
	case 224, 256, 384, 512:
	default:
		*len = 256
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	hash, err := keccak.Keccak(*len, data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%x\n", hash)
}
