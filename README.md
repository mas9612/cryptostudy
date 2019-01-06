# cryptostudy

[![CircleCI](https://circleci.com/gh/mas9612/cryptostudy.svg?style=svg)](https://circleci.com/gh/mas9612/cryptostudy)

Some crypto algorithm implementation with Go.
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

$ go build ./cmd/hmac
$ ./hmac
Usage of ./hmac:
  -algorithm string
        Hash algorithm used to calculate HMAC. Default is "MD5". Valid algorithm is one of [MD5, SHA-1, SHA-224, SHA-256, SHA-384, SHA-512] (default "MD5")
  -key string
        Secret key to calculate HMAC. Specify as hex notation without preceding "0x".
```
