package main

import (
	"fmt"
	"git.noxgroup.com/chenzhiyong/com-nox-middleware-go.git/common"
)

func main() {
	url := "http://chenzhiyong:123@10.8.1.211:80/api/test?a=1&b=2&c=3"
	parser := common.UrlParser{Url: url}

	parser.Parser()
	fmt.Println(parser)
	fmt.Println(parser.GetParamByField("a"))
}
