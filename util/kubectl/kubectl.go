package kubectl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"text/template"
	"time"

	"github.com/Healthism/ih-cli/config"
	"github.com/Healthism/ih-cli/util"
	"github.com/Healthism/ih-cli/util/console"
)

type Job struct {
	Name      string
	Release   string
	Cluster   string
	NameSpace string
	Command   string
	Image     string
	Pod       string
}

func New(nameSpace string, cluster string, release string, command string) Job {
	image, _ := ioutil.ReadFile(config.IMAGE_PATH)
	return Job{
		Name:      fmt.Sprintf("%s-console-%s", release, time.Now().Format("20060102150405")),
		Release:   release,
		Cluster:   cluster,
		NameSpace: nameSpace,
		Command:   command,
		Image:     string(image),
	}
}

func (job *Job) Create() error {
	console.Verbose("[TEMPLATE] Parsing template\n")
	template, err := template.New("template").Parse(JOB_TEMPLATE)
	if err != nil {
		return err
	}

	console.Verbose("[TEMPLATE] Executing template\n")
	var output = bytes.NewBuffer(nil)
	err = template.Execute(output, job)
	if err != nil {
		return err
	}

	console.Verbose("[TEMPLATE] Writing template\n")
	err = ioutil.WriteFile(config.JOB_PATH, output.Bytes(), 0644)
	if err != nil {
		return err
	}

	_, err = util.BufferedExec("kubectl", []string{"--cluster", job.Cluster, "-n", job.NameSpace, "apply", "-f", config.JOB_PATH}...)
	if err != nil {
		return err
	}

	pod, err := util.BufferedExec("kubectl", []string{"--cluster", job.Cluster, "-n", job.NameSpace, "get", "pods", fmt.Sprintf("--selector=job-name=%s", job.Name), "-o=jsonpath='{.items[0].metadata.name}'"}...)
	if err != nil {
		return err
	}

	job.Pod = pod
	return nil
}

func (job *Job) Wait() error {
	_, err := util.BufferedExec("kubectl", []string{"--cluster", job.Cluster, "-n", job.NameSpace, "wait", "--for=condition=ContainersReady", "--timeout=3m", "pod", job.Pod}...)
	return err
}

func (job *Job) Attach() error {
	err := util.InteractiveExec("kubectl", []string{"--cluster", job.Cluster, "-n", job.NameSpace, "attach", "-it", "jobs", job.Name}...)
	return err
}

func (job *Job) Delete() error {
	_, err := util.BufferedExec("kubectl", []string{"--cluster", job.Cluster, "-n", job.NameSpace, "delete", "jobs", job.Name}...)
	return err
}
