package main

import (
	"git.gastrodon.io/imonke/monkebase"
	"git.gastrodon.io/imonke/monkelib"

	"fmt"
	"net/http"
)

func exists(key string, request *http.Request) (code int, r_map map[string]interface{}, err error) {
	var parts []string = monkelib.SplitPath(request.URL.Path)
	var part string = parts[len(parts)-1]
	var exists bool
	switch key {
	case "nick":
		_, exists, err = monkebase.ReadSingleUserNick(part)
	case "email":
		_, exists, err = monkebase.ReadSingleUserEmail(part)
	default:
		err = fmt.Errorf("I have no idea how I got key %s", key)
	}

	code = 200
	r_map = map[string]interface{}{"exists": exists}
	return
}

func checkNick(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	code, r_map, err = exists("nick", request)
	return
}

func checkEmail(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	code, r_map, err = exists("email", request)
	return
}
