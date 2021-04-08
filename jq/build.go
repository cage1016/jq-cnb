package jq

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/cloudfoundry/packit"
)

func Build() packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		file, err := os.Open(filepath.Join(context.CNBPath, "buildpack.toml"))
		if err != nil {
			return packit.BuildResult{}, err
		}

		var m struct {
			Metadata struct {
				Dependencies []struct {
					URI string `toml:"uri"`
				} `toml:"dependencies"`
			} `toml:"metadata"`
		}
		_, err = toml.DecodeReader(file, &m)
		if err != nil {
			return packit.BuildResult{}, err
		}

		uri := m.Metadata.Dependencies[0].URI
		fmt.Printf("URI -> %s\n", uri)

		jq, err := context.Layers.Get("jq", packit.LaunchLayer)
		if err != nil {
			return packit.BuildResult{}, err
		}

		err = jq.Reset()
		if err != nil {
			return packit.BuildResult{}, err
		}

		downloadDir, err := ioutil.TempDir("", "downloadDir")
		if err != nil {
			return packit.BuildResult{}, err
		}
		defer os.RemoveAll(downloadDir)

		fmt.Println("Downloading dependency...")
		err = exec.Command("curl",
			"-Lo", filepath.Join(downloadDir, "jq"),
			uri,
		).Run()
		if err != nil {
			return packit.BuildResult{}, err
		}

		fmt.Println("Moving dependency...")
		data, err := ioutil.ReadFile(filepath.Join(downloadDir, "jq"))
		if err != nil {
			return packit.BuildResult{}, err
		}
		dir := filepath.Join(jq.Path, "bin")
		os.Mkdir(dir, 0777)
		err = ioutil.WriteFile(filepath.Join(dir, "jq"), data, 0777)
		if err != nil {
			return packit.BuildResult{}, err
		}

		return packit.BuildResult{
			Plan: context.Plan,
			Layers: []packit.Layer{
				jq,
			},
		}, nil
	}
}
