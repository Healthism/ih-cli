package yaml

import (
	"ih/lib/log"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func UpdateValue(keyValues []string) error {
	/** Open YAML file **/
	file, err := os.Open("/usr/local/lib/ih/values.yaml")
	if err != nil {
		return log.Error("[YAML] Unable to load values.yaml")
	}

	/** Decode YAML file **/
	root := make(map[interface{}]interface{})
	err = yaml.NewDecoder(file).Decode(&root)
	if err != nil {
		return log.Error("[YAML] Unable to decode values.yaml")
	}

	/** Validate Arguements **/
	keyValue, err := validateArg(keyValues)
	if err != nil {
		return log.Error("[YAML] Invalid arguments")
	}

	/** Update YAML with New Arguements **/
	for key, value := range keyValue {
		inner := root
		splitedKeys := strings.Split(key, ".")
		for i, k := range splitedKeys {
			if i == len(splitedKeys)-1 {
				break
			}
			if _, hasKey := inner[k]; !hasKey {
				return log.Error("[YAML] Requested key does not exist")
			}
			inner = inner[k].(map[interface{}]interface{})
		}

		if _, hasKey := inner[splitedKeys[len(splitedKeys)-1]]; !hasKey {
			return log.Error("[YAML] Requested key does not exist")
		}

		log.Printf("[YAML] Update [%s]: '%s' -> '%s'", splitedKeys[len(splitedKeys)-1], inner[splitedKeys[len(splitedKeys)-1]], value)
		inner[splitedKeys[len(splitedKeys)-1]] = value
	}

	/** Marshal Updated YAML File **/
	d, err := yaml.Marshal(&root)
	if err != nil {
		return log.Error("[YAML] Unable to marshal new yaml file")
	}

	/** Write Updated YAML File **/
	err = ioutil.WriteFile("/usr/local/lib/ih/values.yaml", d, 0644)
	if err != nil {
		return log.Error("[YAML] Failed to write new yaml file")
	}

	log.Print("[YAML] Template Generated")
	return nil
}

func validateArg(args []string) (map[string]string, error) {
	updateTable := make(map[string]string)
	for _, arg := range args {
		argPair := strings.Split(arg, "=")
		if len(argPair) < 2 {
			return nil, log.Errorf("Update Value Must Be Form of 'Key=Value'")
		}

		key := argPair[0]
		value := strings.Join(argPair[1:], "=")

		updateTable[key] = value
	}

	return updateTable, nil
}
