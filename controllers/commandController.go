package controllers

import (
	"Remote/helpers"
	"bufio"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// Defines the SmartContoller type
type RemoteController struct {
}

// Executes a arbitrary command on the hostsystem and remports the result
func (s RemoteController) Exec(c *gin.Context) {
	//user := os.Getenv("REMOTE_USER")
	//dirname, err := os.UserHomeDir()
    //if err != nil {
    //    log.Fatal( err )
    //}
	host := c.Params.ByName("host")
	command := c.DefaultQuery("cmd", "")
	cmd := exec.Command("ssh", host, command)
	stderr, _ := cmd.StderrPipe()       //fetch the stderr in a pipline
	stdout, _ := cmd.StdoutPipe()       //fetch the stdout in a pipline
	if err := cmd.Start(); err != nil { //start the execution of the command and test for imediate errors //TODO what is if there is an error and a stdout?
		helpers.RespondJSON(c, http.StatusInternalServerError, err)
		return
	}

	//Check stderr for errors
	scanner := bufio.NewScanner(stderr)
	errMsg := []string{}
	for scanner.Scan() {
		errMsg = append(errMsg, scanner.Text())
	}
	if len(errMsg) > 0 {
		helpers.RespondJSON(c, http.StatusBadRequest, errMsg)
		return
	}
	//Check stdout for responses
	scanner = bufio.NewScanner(stdout)
	outMsg := []string{}
	//outMsg = append(outMsg, command, "\n")
	for scanner.Scan() {
		outMsg = append(outMsg, scanner.Text())
	}
	helpers.RespondJSON(c, http.StatusOK, outMsg)
}
