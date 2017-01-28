package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"v2ray.com/core"
)

var (
	flagTargetDir    = flag.String("dir", "", "Directory to put generated files.")
	flagTargetOS     = flag.String("os", runtime.GOOS, "Target OS of this build.")
	flagTargetArch   = flag.String("arch", runtime.GOARCH, "Target CPU arch of this build.")
	flagArchive      = flag.Bool("zip", false, "Whether to make an archive of files or not.")
	flagMetadataFile = flag.String("metadata", "metadata.txt", "File to store metadata info of released packages.")

	binPath string
)

func createTargetDirectory(version string, goOS GoOS, goArch GoArch) (string, error) {
	var targetDir string
	if len(*flagTargetDir) > 0 {
		targetDir = *flagTargetDir
	} else {
		suffix := getSuffix(goOS, goArch)

		targetDir = filepath.Join(binPath, "v2ray-"+version+suffix)
		if version != "custom" {
			os.RemoveAll(targetDir)
		}
	}

	err := os.MkdirAll(targetDir, os.ModeDir|0777)
	return targetDir, err
}

func getTargetFile(goOS GoOS) string {
	suffix := ""
	if goOS == Windows {
		suffix += ".exe"
	}
	return "v2ray" + suffix
}

func getBinPath() string {
	GOPATH := os.Getenv("GOPATH")
	return filepath.Join(GOPATH, "bin")
}

func main() {
	flag.Parse()
	binPath = getBinPath()
	build(*flagTargetOS, *flagTargetArch, *flagArchive, "", *flagMetadataFile)
}

func build(targetOS, targetArch string, archive bool, version string, metadataFile string) {
	v2rayOS := parseOS(targetOS)
	v2rayArch := parseArch(targetArch)

	if len(version) == 0 {
		version = os.Getenv("TRAVIS_TAG")
	}
	if len(version) == 0 {
		version = core.Version()
	}

	fmt.Printf("Building V2Ray (%s) for %s %s\n", version, v2rayOS, v2rayArch)

	targetDir, err := createTargetDirectory(version, v2rayOS, v2rayArch)
	if err != nil {
		fmt.Println("Unable to create directory " + targetDir + ": " + err.Error())
	}

	targetFile := getTargetFile(v2rayOS)
	err = buildV2Ray(filepath.Join(targetDir, targetFile), version, v2rayOS, v2rayArch)
	if err != nil {
		fmt.Println("Unable to build V2Ray: " + err.Error())
	}

	err = copyConfigFiles(targetDir, v2rayOS)
	if err != nil {
		fmt.Println("Unable to copy config files: " + err.Error())
	}

	if archive {
		err := os.Chdir(binPath)
		if err != nil {
			fmt.Printf("Unable to switch to directory (%s): %v\n", binPath, err)
		}
		suffix := getSuffix(v2rayOS, v2rayArch)
		zipFile := "v2ray" + suffix + ".zip"
		root := filepath.Base(targetDir)
		err = zipFolder(root, zipFile)
		if err != nil {
			fmt.Printf("Unable to create archive (%s): %v\n", zipFile, err)
		}

		metadataWriter, err := os.OpenFile(filepath.Join(binPath, metadataFile), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			fmt.Printf("Unable to create metadata file (%s): %v\n", metadataFile, err)
		}
		defer metadataWriter.Close()

		err = CalcMetadata(zipFile, metadataWriter)
		if err != nil {
			fmt.Printf("Failed to calculate metadata for file (%s): %v", zipFile, err)
		}
	}
}
