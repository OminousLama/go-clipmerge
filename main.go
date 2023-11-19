package main

import (
	"fmt"
	"os"
	"os/exec"
)

var version = "undefined"
var metaBuildTime = "undefined"
var metaBuilderOS = "undefined"
var metaBuilderArch = "undefined"

func main() {

	if len(os.Args) < 4 {
		fmt.Println("Usage: clipmerge output.mp4 input1.mp4 input2.mp4 ...")
		return
	}

	fmt.Println("fvidl version info:")
	fmt.Println("- Version:", version)
	fmt.Println("- Build time:", metaBuildTime)
	fmt.Println("- Builder OS:", metaBuilderOS)
	fmt.Println("- Builder Arch:", metaBuilderArch)

	outputFileName := os.Args[1]
	inputFileNames := os.Args[2:]

	if err := mergeVideos(outputFileName, inputFileNames); err != nil {
		fmt.Println("Error:", err)
	}
}

func mergeVideos(outputFileName string, inputFiles []string) error {

	tempFileName := "ffmpeg_concat_list.txt"
	tempFile, err := os.Create(tempFileName)
	if err != nil {
		return err
	}
	defer os.Remove(tempFileName)
	defer tempFile.Close()

	for _, inputFile := range inputFiles {
		_, err := fmt.Fprintf(tempFile, "file '%s'\n", inputFile)
		if err != nil {
			return err
		}
	}

	ffmpegArgs := []string{"-y", "-f", "concat", "-safe", "0", "-i", tempFileName, "-c", "copy", outputFileName}

	cmd := exec.Command("ffmpeg", ffmpegArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Printf("Merged videos successfully. Output saved to %s\n", outputFileName)
	return nil
}
