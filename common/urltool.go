package common

import (
	"strconv"
	"strings"
)

type UrlParser struct {
	Url      string
	Protocol string
	UserName string
	Password string
	Host     string
	Port     int
	Path     string
	Params   map[string]string
}

func (parser *UrlParser) Init(url string) {
	parser.Url = url
	parser.Parser()
}
func (parser *UrlParser) Parser() {
	if len(parser.Url) == 0 {
		return
	}
	url := parser.Url
	index := strings.Index(url, "://")
	if index != -1 {
		parser.Protocol = url[:index]
		url = url[index+3:]
	}
	index = strings.Index(url, "?")
	if index != -1 {
		params := url[index+1:]
		url = url[0:index]
		if len(params) != 0 {
			splits := strings.Split(params, "&")
			parser.Params = make(map[string]string)
			for _, split := range splits {
				tempIndex := strings.Index(split, "=")
				if tempIndex != -1 {
					parser.Params[split[0:tempIndex]] = split[tempIndex+1:]
				} else {
					parser.Params[split] = ""
				}
			}
		}
	}
	index = strings.Index(url, "/")
	if index != -1 {
		parser.Path = url[index+1:]
		url = url[0:index]
	}
	index = strings.Index(url, "@")
	if index != -1 {
		account := url[0:index]
		url = url[index+1:]
		tempIndex := strings.Index(account, ":")
		if tempIndex != -1 {
			parser.UserName = account[0:tempIndex]
			parser.Password = account[tempIndex+1:]
		} else {
			parser.UserName = account
		}
	}
	index = strings.Index(url, ":")
	if index != -1 {
		parser.Host = url[0:index]
		parser.Port, _ = strconv.Atoi(url[index+1:])
	} else {
		parser.Host = url
	}
}
func (parser *UrlParser) GetParamByField(field string) string {
	if parser.Params == nil {
		return ""
	}
	return parser.Params[field]
}
