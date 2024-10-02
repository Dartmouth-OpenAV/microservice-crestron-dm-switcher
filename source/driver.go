package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"strconv"
	"github.com/Dartmouth-OpenAV/microservice-framework/framework"
)

// GLOBAL VARIABLES
var cmdPrompt = ""

// Reads a response from the Crestron DM.
func readAndConvert(socketKey string) (string, string, error) {
	function := "readAndConvert"
	resp := framework.ReadLineFromSocket(socketKey)

	// Normally, the Crestron DM may return a blank line, so just log it. No error.
	if resp == "" {
		framework.Log(function + " - k3kxlpo - Response was blank.")
	}

	return resp, "", nil
}
// Handles the Telnet negotiation for the Crestron DM connection
func loginNegotiation(socketKey string) bool {
	function := "loginNegotiation"
	count := 0
	// Breaks if the negotiations go over 7 rounds to avoid an infinite loop.
	// Normal negotiations so far are 3-4 rounds.
	for count < 7 {
		count += 1
		negotiationResp, errMsg, err := readAndConvert(socketKey)
		if err != nil {
			framework.AddToErrors(socketKey, errMsg)
		}
		framework.Log("Printing Negotiation from Crestron DM switcher: " + negotiationResp)
		if strings.HasPrefix(negotiationResp, "DM") && strings.HasSuffix(negotiationResp, ">") {
			cmdPrompt = negotiationResp
            framework.Log("Negotiations are over. Command line prompt is "+cmdPrompt)
            return true
		}
	}
	errMsg := function + " - mrk42 - Stopped negotiation loop after 7 rounds to avoid infinite loop."
	framework.AddToErrors(socketKey, errMsg)

	return false
}

//MAIN FUNCTIONS

// GET Functions

func getAVRoute(socketKey string, channel string) (string, error) {
	function := "getAVRoute"

	value := `"unknown"`
	err := error(nil)
	maxRetries := 2
	for maxRetries > 0 {
		value, err = getAVRouteDo(socketKey, channel)
		if value == `"unknown"` { // Something went wrong - perhaps try again
			framework.Log(function + " - fq3sdvc - retrying getAVRoute operation")
			maxRetries--
			time.Sleep(1 * time.Second)
			if maxRetries == 0 {
				errMsg := fmt.Sprintf(function + "f839dk4 - max retries reached")
				framework.AddToErrors(socketKey, errMsg)
			}
		} else { // Succeeded
			maxRetries = 0
		}
	}

	return value, err
}
// Gets the routing information of the Crestron DM switcher. Returns several lines.
func getAVRouteDo(socketKey string, channel string) (string, error) {
	function := "getAVRouteDo"

	connected := framework.ConnectionsMapExists(socketKey)
	if connected == false{
		negotiation := loginNegotiation(socketKey)
		if negotiation == false {
			errMsg := fmt.Sprintf(function + " - h3okxu3 - error connecting")
			framework.AddToErrors(socketKey, errMsg)
			return errMsg, errors.New(errMsg)
		}
	}

	// dump the routing info and parse for the specific channel
	cmdString := "dumpdmrouteinfo\r"

	sent := framework.WriteLineToSocket(socketKey, cmdString)

	if sent != true {
		errMsg := fmt.Sprintf(function + " - dj3dke - error sending command")
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}

	// parse line by line looking for Out{channel}->In{value}
	respArr, errMsg, err := readAndConvert(socketKey)
	for respArr != cmdPrompt{
		if err != nil{
			return errMsg, err
		}
		if strings.Contains(respArr, "Out"+channel+"->In") {
			framework.Log("Found Output Channel "+channel)
			break
		}
		if strings.HasPrefix(respArr, "dumprouteinfo\r"){
			framework.Log("GOT AN ECHO")
		}
		respArr, errMsg, err = readAndConvert(socketKey)
	}
	_, value, found := strings.Cut(respArr, "Out"+channel+"->In")
	// Read remaining lines to empty the buffer
	respLine := respArr
	for respLine != cmdPrompt{
		respLine, errMsg, err = readAndConvert(socketKey)
	}
	framework.Log("Read Buffer Empty")
	if found != true{
		errMsg := fmt.Sprintf(function + " - kcj3jj - error reading response")
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}
	framework.Log(function + " - Response: "+ value)

	// If we got here, the response was good, so successful return with the state indication
	return `"` + value + `"`, nil
}


//SET Functions

func setAVRoute(socketKey string, output string, input string) (string, error) {
	function := "setAVRoute"

	value := "notok"
	err := error(nil)
	maxRetries := 2
	for maxRetries > 0 {
		value, err = setAVRouteDo(socketKey, output, input)
		if value != "ok" { // Something went wrong - perhaps try again
			framework.Log(function + " - fq3svucc - retrying setAVRoute operation")
			maxRetries--
			time.Sleep(1 * time.Second)
			if maxRetries == 0 {
				errMsg := fmt.Sprintf(function + " - fds3nf3 - max retries reached")
				framework.AddToErrors(socketKey, errMsg)
			}
		} else { // Succeeded
			maxRetries = 0
		}
	}

	return value, err
}
// Sets the input to the given output channel.
func setAVRouteDo(socketKey string, output string, input string) (string, error) {
	function := "setAVRouteDo"
	input = strings.Trim(input, "\"")

	connected := framework.ConnectionsMapExists(socketKey)
	if connected == false{
		negotiation := loginNegotiation(socketKey)
		if negotiation == false {
			errMsg := fmt.Sprintf(function + " - h3okxu3 - error connecting")
			framework.AddToErrors(socketKey, errMsg)
			return errMsg, errors.New(errMsg)
		}
	}

	outnum , err := strconv.Atoi(output)
	//check if error occured
	if err != nil{
	//executes if there is any error
		errMsg := fmt.Sprintf(function + " - i5kcfof - error converting output channel to integer")
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}
	slot := strconv.Itoa(outnum+16)
	framework.Log("Value of slot is "+slot)
	cmdString := "setavroute " + input + " " + slot + "\r"

	sent := framework.WriteLineToSocket(socketKey, cmdString)

	if sent != true {
		errMsg := fmt.Sprintf(function + " - i5kcfoe - error sending command")
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}

	respArr, errMsg, err := readAndConvert(socketKey)

	if err != nil{
		return errMsg, err
	}

	if strings.HasPrefix(respArr, cmdPrompt){
		framework.Log("No error")
	} else if strings.HasPrefix(respArr, "setavroute " + input + " " + slot) {
		framework.Log("GOT AN ECHO")
		readAndConvert(socketKey)
	} else {
		errMsg := fmt.Sprintf(function + " - uygv8 - Error: " + respArr)
		framework.AddToErrors(socketKey, errMsg)
		return errMsg, errors.New(errMsg)
	}
	
	framework.Log(function + " - Response: "+ respArr)

	// If we got here, the response was good, so successful return with the state indication
	return "ok", nil
}
