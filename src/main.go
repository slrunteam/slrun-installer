package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"strings"
  "path/filepath"
  "runtime"
)

func downloadFile(url string, dest string) (file string) {
	tokens := strings.Split(url, "/")
	filePath := filepath.Join(dest, tokens[len(tokens)-1])

	output, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error while creating", filePath, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	if _, err := io.Copy(output, response.Body); err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}  
  return filePath
}

func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

func getTargetOS() (targetOS string){
  var re = "linux"
  if (runtime.GOOS == "windows") {
    re = "win"
  } else if (runtime.GOOS == "darwin") {
    re = "macos"
  }
  if (runtime.GOARCH == "amd64") {
    re += "-x64"
  } else if (runtime.GOARCH == "amd64") {
    re += "-x86"
  }
  return re
}

func main() {
  // Initialize
	user, _ := user.Current()
	var homeDir = user.HomeDir + "/." + AppName
	fmt.Print("Installation path [", homeDir, "]: ")
	var homeDirInput string
	fmt.Scanln(&homeDirInput)
	if homeDirInput != "" {
		homeDir = homeDirInput
  }

  var version = "latest"
  fmt.Print("Version [", version, "]: ")
	var versionInput string
	fmt.Scanln(&versionInput)
	if versionInput != "" {
		version = versionInput
	}
  
  var downloadUrl = strings.Replace(DownloadUrl, "{version}", version, -1)
  downloadUrl = strings.Replace(downloadUrl, "{env}", getTargetOS(), -1)

	// Install
	fmt.Println("Installing...")
  var file = downloadFile(downloadUrl, os.TempDir())
	unzip(file, homeDir)
	fmt.Println(file)
  os.Remove(file)
  fmt.Println("Finished")
}
