package git

import (
	"os/exec"
	"time"

	"ih/lib/log"
)

const (
	ADD      = "add"
	PUSH     = "push"
	PULL     = "pull"
	COMMIT   = "commit"
	CHECKOUT = "checkout"
)

/**
 * Public
**/

func Get(target string) error {
	err := gitExec(CHECKOUT, target)
	if err != nil {
		return log.Errorf("[GIT] Failed to 'CHECKOUT': %v", err)
	}

	err = gitExec(PULL)
	if err != nil {
		return log.Errorf("[GIT] Failed to 'PULL': %v", err)
	}

	log.Printf("[GIT] Loaded release resource - %s", target)
	return nil
}

func Set(target string) error {
	err := gitExec(ADD, ".")
	if err != nil {
		err = log.Errorf("[GIT] Failed to 'ADD': %v", err)
		recover(target)
		return err
	}

	err = gitExec(COMMIT, "-m", time.Now().Format(time.RFC850))
	if err != nil {
		err = log.Errorf("[GIT] Failed to 'COMMIT': %v", err)
		recover(target)
		return err
	}

	err = gitExec(PUSH)
	if err != nil {
		err = log.Errorf("[GIT] Failed to 'PUSH': %v", err)
		recover(target)
		return err
	}

	log.Print("[GIT] Updated changes to the server")
	return nil
}

func Reset() {
	err := gitExec(CHECKOUT, "master")
	if err != nil {
		log.Errorf("[GIT] Failed to 'Reset': %v", err)
		return
	}
}

/**
 * Private
**/

func recover(target string) {
	origin := "origin/" + target
	err := gitExec("reset", "--hard", origin)
	if err != nil {
		log.Errorf("[GIT] Failed to 'RECOVER': %v", err)
		return
	}

	Reset()
	log.Print("[GIT] Recoverd from Error")
}

func gitExec(options ...string) error {
	var gitCmd = "git"
	var gitOption = []string{"-C", "/usr/local/lib/ih"}

	git := exec.Command(gitCmd, append(gitOption, options...)...)

	stdoutStderr, err := git.CombinedOutput()
	log.CombinedStd("%s", stdoutStderr)

	return err
}
