package git

import (
	"time"

	"github.com/Healthism/ih-cli/config"
	"github.com/Healthism/ih-cli/util"
)

func Load(target string) (err error) {
	_, err = util.BufferedExec("git", []string{"-C", config.GIT_PATH, "reset", "--hard"}...)
	if err != nil {
		return err
	}

	_, err = util.BufferedExec("git", []string{"-C", config.GIT_PATH, "clean", "-df"}...)
	if err != nil {
		return err
	}

	_, err = util.BufferedExec("git", []string{"-C", config.GIT_PATH, "checkout", target}...)
	if err != nil {
		return err
	}

	_, err = util.BufferedExec("git", []string{"-C", config.GIT_PATH, "pull"}...)
	if err != nil {
		return err
	}

	return
}

func Save(target string) (err error) {
	_, err = util.BufferedExec("git", []string{"-C", config.GIT_PATH, "add", "."}...)
	if err != nil {
		return err
	}

	_, err = util.BufferedExec("git", []string{"-C", config.GIT_PATH, "commit", "-m", time.Now().Format(time.RFC850)}...)
	if err != nil {
		return err
	}

	_, err = util.BufferedExec("git", []string{"-C", config.GIT_PATH, "push"}...)
	if err != nil {
		return err
	}

	return
}
