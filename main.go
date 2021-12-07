package main

import (
	"fmt"
	"nextcloudUploader/run"
	"nextcloudUploader/utils"
	"os"
)

func main() {
	var bs64 run.Base64
	var local run.Local
	var http run.Http
	var r struct {
		url string
		req string
	}

	bs64.ConfigPath = utils.ConfigPath("./config.json")
	local.ConfigPath = utils.ConfigPath("./config.json")
	http.ConfigPath = utils.ConfigPath("./config.json")

	for idx, args := range os.Args {
		if idx == 0 {
			continue
		}
		r.req = utils.FileType(&args)
		if r.req == "base64" {
			r.url = run.Run(&bs64, &args)
		} else if r.req == "url" {
			r.url = run.Run(&http, &args)
		} else if r.req == "local" {
			r.url = run.Run(&local, &args)
		}
		if r.url != "" {
			fmt.Printf("Upload Success:\n")
			fmt.Printf(r.url + "\n")
		}
	}
}
