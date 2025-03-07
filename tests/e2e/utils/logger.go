/*
Copyright 2021 The KubeEdge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	log "k8s.io/klog"
	"os/exec"
	"strconv"
	"strings"
)

func PrintCmdOutput(cmd *exec.Cmd) error {
	log.Info("Executing command: ", strings.Join(cmd.Args, " "))
	stdOutStdErr, err := cmd.CombinedOutput()
	if err != nil {
		log.Error("fail to executing command: ", err)
		return err
	}
	log.Infof("executes command result: %s", stdOutStdErr)
	return nil
}

// read n line container log
func ReadDockerLog(getContainerID string, n int) (string, error) {
	cmd := exec.Command("sh", "-c", getContainerID)
	result, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	containerID := string(result[:12])
	cmd = exec.Command("sh", "-c", "docker logs --tail="+strconv.Itoa(n)+" "+containerID)
	result, err = cmd.CombinedOutput()

	return string(result), nil
}
