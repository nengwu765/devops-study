package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "http://wg-dl.tech.2345.cn/mobile/artifacts-android/market-channelpack/channelpack/463/zip/output.zip"

	splits := strings.Split(str, "/mobile/")

	fmt.Printf("%v",splits)

}
