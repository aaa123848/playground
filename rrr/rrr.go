package rrr

import "log"

type WaletLog string

const (
	Settle WaletLog = "settle the wallet"
	Cancel WaletLog = "cancel"
)

func StroeWalletLog(s WaletLog) {
	log.Printf("do something to wallet log %v \n", s)
}
