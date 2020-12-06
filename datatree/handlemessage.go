package datatree

import (
	"strconv"
	"strings"
	"zookeepergo/network"
)

//DealWithMessage deal with message about create or delete
func DealWithMessage(Message network.NetMessage, root *Datatree) {
	str1 := ""
	if Message.Type == 5 {
		CreateZnode(Message.Str, root, &str1)
	}
	if Message.Type == 6 {
		DeleteZnode(Message.Str, root, &str1)
	}
	if Message.Type == 8 {
		CreateWatcher(Message.Str, root, Message.Info)
	}
	if str1 != "" {
		strset := strings.Split(str1, ":")
		n, err := strconv.Atoi(strset[0])
		if err == nil {
			network.SendOnetimeMessage(n, 4, 7, n, strset[1]+" "+strset[2])
		}

	}
}
