package main

import (
	"github.com/Healthism/ih-cli/cmd"
	// "encoding/json"
	// "fmt"
	// "net/http"
	// "strings"
	// "github.com/Healthism/ih-cli/config"
	// "github.com/Healthism/ih-cli/util/console"
)

func main() {
	cmd.Execute()
	// versionInfo, err := http.Get(config.CLI_GIT_URL)
	// if err != nil {
	// 	console.Errorf("[UPDATE] Failed to load latest version information: %v", err)
	// 	return
	// }

	// defer versionInfo.Body.Close()
	// version := make(map[string]interface{})
	// json.NewDecoder(versionInfo.Body).Decode(&version)
	// latestVersion := fmt.Sprintf("%v", version["name"])
	// if latestVersion == "dfdf" {
	// 	return
	// }

	// console.Print(console.SprintBlue("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓"))
	// console.Printf("%s  %-83v%s", console.SprintBlue("┃"), console.SprintYellow(fmt.Sprintf("Version %s -> %s", latestVersion, latestVersion)), console.SprintBlue("┃"))
	// console.Printf("%s  %-83v%s", console.SprintBlue("┃"), console.SprintYellow(" "), console.SprintBlue("┃"))

	// whatsNew := fmt.Sprintf("%v", version["body"])
	// if whatsNew != "" {
	// 	for _, infoLine := range strings.Split(whatsNew, "\n") {
	// 		wrappedString := wrapString(infoLine, 75)
	// 		for _, xx := range strings.Split(wrappedString, "\n") {
	// 			console.Printf("%s  %-74v%s", console.SprintBlue("┃"), xx, console.SprintBlue("┃"))
	// 		}
	// 	}
	// }
	// console.Print(console.SprintBlue("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛"))
}

// func wrapString(text string, lineWidth int) string {
// 	words := strings.Fields(strings.TrimSpace(text))
// 	if len(words) == 0 {
// 		return text
// 	}
// 	wrapped := words[0]
// 	spaceLeft := lineWidth - len(wrapped)
// 	for _, word := range words[1:] {
// 		if len(word)+1 > spaceLeft {
// 			wrapped += "\n" + word
// 			spaceLeft = lineWidth - len(word)
// 		} else {
// 			wrapped += " " + word
// 			spaceLeft -= 1 + len(word)
// 		}
// 	}

// 	return wrapped

// }
