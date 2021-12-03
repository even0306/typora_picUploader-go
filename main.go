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

	bs64.ConfigPath = utils.ConfigPath(utils.GetLocalPath() + "/config.json")
	local.ConfigPath = utils.ConfigPath(utils.GetLocalPath() + "/config.json")
	http.ConfigPath = utils.ConfigPath(utils.GetLocalPath() + "/config.json")

	// test := []string{"data:image/jpeg;base64,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"}
	// test := []string{"E:\\abc.jpg"}
	// test := []string{"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSVkhuBwzZrFrATlA7L5clrFsFjttL8HNcs8Q&usqp=CAU"}

	// for idx, args := range test {
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
