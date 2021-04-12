package jq

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/cloudfoundry/packit"
)

func Detect() packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {

		fn := func() (packit.DetectResult, error) {
			return packit.DetectResult{
				Plan: packit.BuildPlan{
					Provides: []packit.BuildPlanProvision{
						{Name: "jq"},
					},
					Requires: []packit.BuildPlanRequirement{
						{Name: "jq"},
					},
				},
			}, nil
		}

		if _, err := os.Stat(filepath.Join(context.WorkingDir, ".jq-version")); os.IsNotExist(err) {
			return fn()
		}

		file, _ := os.Open(filepath.Join(context.WorkingDir, ".jq-version"))
		jqVersion, err := ioutil.ReadAll(file)
		if err != nil {
			return fn()
		}

		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{
					{Name: "jq"},
				},
				Requires: []packit.BuildPlanRequirement{
					{
						Name: "jq",
						Metadata: map[string]string{
							"jq-version": string(jqVersion),
						},
					},
				},
			},
		}, nil
	}
}
