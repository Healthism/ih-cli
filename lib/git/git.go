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
	log.Printf("[GIT] Loading release - %s", target)
	err := gitExec(CHECKOUT, target)
	if err != nil {
		return log.Errorf("[GIT] Failed to 'CHECKOUT': %v", err)
	}

	log.Print("[GIT] Pulling latest configuration file")
	err = gitExec(PULL)
	if err != nil {
		return log.Errorf("[GIT] Failed to 'PULL': %v", err)
	}

	return nil
}

func Set(target string) error {
	log.Print("[GIT] Reflecting changes")
	err := gitExec(ADD, ".")
	if err != nil {
		err = log.Errorf("[GIT] Failed to 'ADD': %v", err)
		recover(target)
		return err
	}

	log.Print("[GIT] Commiting")
	err = gitExec(COMMIT, "-m", time.Now().Format(time.RFC850))
	if err != nil {
		err = log.Errorf("[GIT] Failed to 'COMMIT': %v", err)
		recover(target)
		return err
	}

	log.Print("[GIT] Pushing changes to the server")
	err = gitExec(PUSH)
	if err != nil {
		err = log.Errorf("[GIT] Failed to 'PUSH': %v", err)
		recover(target)
		return err
	}

	log.Print("[GIT] Clean up")
	err = gitExec(CHECKOUT, "master")
	if err != nil {
		err = log.Errorf("[GIT] Failed to 'CLEAN UP': %v", err)
		recover(target)
		return err
	}

	return nil
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

	err = gitExec(CHECKOUT, "master")
	if err != nil {
		log.Errorf("[GIT] Failed to 'RECOVER': %v", err)
		return
	}

	log.Print("[GIT] Recoverd from Error")
}

func gitExec(options ...string) error {
	var gitCmd = "git"
	var gitOption = []string{"-C", "config"}

	git := exec.Command(gitCmd, append(gitOption, options...)...)

	stdoutStderr, err := git.CombinedOutput()
	log.Debugf("%s", stdoutStderr)

	return err
}
