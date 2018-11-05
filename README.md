# cryptostudy

[![CircleCI](https://circleci.com/gh/mas9612/cryptostudy.svg?style=svg)](https://circleci.com/gh/mas9612/cryptostudy)

AES implementation with Go.
Source codes in this repository are written for personal use.

```
$ go build ./cmd/aestest
$ ./aestest -help
Usage of AES:
  -K string
        Encrypt key (hexadecimal notation)
  -d    Decrypt (Default Encrypt)
  -help
        Print help and exit
  -iv string
        IV
  -mode string
        Cipher mode. Valid mode is one of [ECB, CBC, CFB, OFB, CTR]
  -r int
        Print round N result (default -1)

$ go build ./cmd/extgcd
$ ./extgcd 5 13
8
```
