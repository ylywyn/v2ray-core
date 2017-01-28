package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"v2ray.com/core/testing/assert"
)

func fileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func allFilesExists(files ...string) bool {
	for _, file := range files {
		fullPath := filepath.Join(binPath, file)
		if !fileExists(fullPath) {
			fmt.Println(fullPath + " doesn't exist.")
			return false
		}
	}
	return true
}

func TestBuildMacOS(t *testing.T) {
	assert := assert.On(t)
	tmpPath, err := ioutil.TempDir("", "v2ray")
	assert.Error(err).IsNil()

	binPath = tmpPath

	build("macos", "amd64", true, "test", "metadata.txt")
	assert.Bool(allFilesExists(
		"v2ray-macos.zip",
		"v2ray-test-macos",
		filepath.Join("v2ray-test-macos", "config.json"),
		filepath.Join("v2ray-test-macos", "v2ray"))).IsTrue()

	build("windows", "amd64", true, "test", "metadata.txt")
	assert.Bool(allFilesExists(
		"v2ray-windows-64.zip",
		"v2ray-test-windows-64",
		filepath.Join("v2ray-test-windows-64", "config.json"),
		filepath.Join("v2ray-test-windows-64", "v2ray.exe"))).IsTrue()

	build("linux", "amd64", true, "test", "metadata.txt")
	assert.Bool(allFilesExists(
		"v2ray-linux-64.zip",
		"v2ray-test-linux-64",
		filepath.Join("v2ray-test-linux-64", "vpoint_socks_vmess.json"),
		filepath.Join("v2ray-test-linux-64", "vpoint_vmess_freedom.json"),
		filepath.Join("v2ray-test-linux-64", "v2ray"))).IsTrue()
}
