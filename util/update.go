package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"

	"github.com/Healthism/ih-cli/config"
	"github.com/Healthism/ih-cli/util/console"
	"github.com/inconshreveable/go-update"
)

type MetaData struct {
	Version string
	Changes []string
}

func UpdateCLI(currVersion string) (bool, error) {
	versionInfo, err := http.Get(config.CLI_GIT_URL)
	if err != nil {
		console.Error("⚠️  Error occured while checking update ...")
		return false, err
	}
	defer versionInfo.Body.Close()

	var metaData MetaData
	jsonBytes, err := ioutil.ReadAll(versionInfo.Body)
	if err != nil {
		console.Error("⚠️  Error occured while checking update ...")
		return false, err
	}

	err = json.Unmarshal(jsonBytes, &metaData)
	if err != nil {
		console.Error("⚠️  Error occured while checking update ...")
		return false, err
	}
	if metaData.Version == currVersion {
		return false, nil
	}

	console.AddTable([]string{
		fmt.Sprintf("%s", console.SprintYellow("IH CLI needs update")),
		fmt.Sprintf("%s", console.SprintYellow("Updating Input Health Command Line Interface...")),
		fmt.Sprintf("%s", console.SprintYellow("Version "+currVersion+"  ➜  "+metaData.Version)),
	})

	latestVersionUrl := fmt.Sprintf(config.CLI_DOWNLOAD_URL, metaData.Version)
	if runtime.GOOS == "linux" {
		latestVersionUrl += "-linux"
	}

	getLoading := console.ShowLoading("Downloading latest CLI", "[1/2]")
	latestRelease, err := http.Get(latestVersionUrl)
	getLoading.HideLoading(err)
	if err != nil {
		showUpdateError(err)
		return false, err
	}
	defer latestRelease.Body.Close()

	updateLoading := console.ShowLoading("Updating CLI", "[2/2]")
	err = update.Apply(latestRelease.Body, update.Options{})
	updateLoading.HideLoading(err)
	if err != nil {
		err = update.RollbackError(err)
		showUpdateError(err)
		return false, err
	}

	console.AddLine()
	console.Print("Changes: ")
	for _, change := range metaData.Changes {
		console.Print(change)
	}

	console.AddLine()
	console.Print(console.SprintYellow("🚀 Update Complete"))
	console.Print(console.SprintYellow("Please enter your command again"))

	return true, nil
}

func showUpdateError(err error) {
	console.Errorf("%s", err)
	console.AddLine()
	console.Error("⚠️  Error occured while updating cli ...")
	console.Error("Please try to again few moment later.")
}

func wrapString(text string, lineWidth int) string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return text
	}
	wrapped := words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}

	return wrapped
}
