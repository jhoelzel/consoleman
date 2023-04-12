package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	// ANSI escape codes
	clearScreen = "\033[2J"
	moveCursor  = "\033[%d;%dH"
)

type requestData struct {
	Protocol    string
	URL         string
	RequestType string
	Auth        string
	Headers     string
	Body        string
}

func main() {
	protocolPtr := flag.String("protocol", "", "Protocol (http or https)")
	urlPtr := flag.String("url", "", "URL")
	requestTypePtr := flag.String("requestType", "", "Request Type (GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS, CONNECT)")
	authPtr := flag.String("auth", "", "Basic Auth credentials (username:password)")
	headersPtr := flag.String("headers", "", "Headers (Key:Value, separated by ';')")
	bodyPtr := flag.String("body", "", "Body")
	noUIPtr := flag.Bool("noUI", false, "Skip empty flags when any flags are provided")

	flag.Parse()
	// if all flags are provided, skip the UI
	// Logo and Response will still be displayed
	noUI := *noUIPtr || (*protocolPtr != "" && *urlPtr != "" && *requestTypePtr != "" && *authPtr != "" && *headersPtr != "" && *bodyPtr != "")

	data := requestData{
		Protocol:    *protocolPtr,
		URL:         *urlPtr,
		RequestType: *requestTypePtr,
		Auth:        *authPtr,
		Headers:     *headersPtr,
		Body:        *bodyPtr,
	}
	reader := bufio.NewReader(os.Stdin)
	if !(*noUIPtr) {
		displayConsoleman()
		time.Sleep(time.Second * 1) // wait 1 seconds before restarting clearing the screen for dramatic effect
		fmt.Print(clearScreen)
	}
	// Collect protocol
	if data.Protocol == "" {
		fmt.Print(clearScreen)
		displayPreviousData(&data)
		data.Protocol = selectProtocol(reader, &data)
	}

	// Collect URL
	if data.URL == "" {
		fmt.Print(clearScreen)
		displayPreviousData(&data)
		data.URL = inputURL(reader, &data)
	}

	// Collect request type
	if data.RequestType == "" {
		fmt.Print(clearScreen)
		displayPreviousData(&data)
		data.RequestType = selectRequestType(reader, &data)
	}

	// Collect authentication
	if data.Auth == "" && !(*noUIPtr) {
		fmt.Print(clearScreen)
		displayPreviousData(&data)
		data.Auth = inputAuth(reader, &data)
	}

	// Collect headers
	if data.Headers == "" && !(*noUIPtr) {
		fmt.Print(clearScreen)
		displayPreviousData(&data)
		data.Headers = inputHeaders(reader, &data)
	}

	// Collect body
	if data.Body == "" && !(*noUIPtr) {
		fmt.Print(clearScreen)
		displayPreviousData(&data)
		data.Body = inputBody(reader, &data)
	}

	if !(noUI) {
		fmt.Print(clearScreen)
	}
	// Send request
	response, err := sendRequest(data.Protocol, data.URL, data.RequestType, data.Auth, data.Headers, data.Body)
	if err != nil {
		fmt.Println("----------------------------------------")
		fmt.Println("Error:")
		fmt.Println("----------------------------------------")
		fmt.Println(err)
	} else {
		if !(*noUIPtr) {
			fmt.Printf(moveCursor, 1, 1)
			fmt.Println("----------------------------------------")
			fmt.Println("Response:")
			fmt.Println("----------------------------------------")
		}
		fmt.Println(response)
	}
}

func displayConsoleman() {
	fmt.Println(`
------------------------------------------------------------------------	
_________                            .__                                
\_   ___ \  ____   ____   __________ |  |   ____   _____ _____    ____  
/    \  \/ /  _ \ /    \ /  ___/  _ \|  | _/ __ \ /     \\__  \  /    \ 
\     \___(  <_> )   |  \\___ (  <_> )  |_\  ___/|  Y Y  \/ __ \|   |  \
 \______  /\____/|___|  /____  >____/|____/\___  >__|_|  (____  /___|  /
        \/            \/     \/                \/      \/     \/     \/ 
			Like Postman, but in the console!	
------------------------------------------------------------------------			
		`)
}

func displayPreviousData(data *requestData) {
	fmt.Printf(moveCursor, 1, 1)
	fmt.Println("----------------------------------------")
	fmt.Printf("Protocol: %s\n", data.Protocol)
	fmt.Printf("URL: %s\n", data.URL)
	fmt.Printf("Request Type: %s\n", data.RequestType)
	fmt.Printf("Auth: %s\n", data.Auth)
	fmt.Printf("Headers: %s\n", data.Headers)
	fmt.Println("----------------------------------------")
}

func selectProtocol(reader *bufio.Reader, data *requestData) string {

	protocols := []string{"http", "https"}
	fmt.Printf(moveCursor, 8, 1)
	fmt.Println("Select Protocol:")
	for i, p := range protocols {
		fmt.Printf(" %d. %s\n", i+1, p)
	}

	fmt.Printf(moveCursor, len(protocols)+8+1, 1)
	fmt.Print("Enter the number corresponding to the protocol: ")
	var protocolIndex int
	fmt.Scanf("%d", &protocolIndex)
	protocol := protocols[protocolIndex-1]

	return protocol
}

func inputURL(reader *bufio.Reader, data *requestData) string {
	fmt.Printf(moveCursor, 8, 1)
	fmt.Print("Enter URL: ")
	url, _ := reader.ReadString('\n')
	url = strings.TrimSpace(url)

	return url
}

func selectRequestType(reader *bufio.Reader, data *requestData) string {
	requestTypes := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "CONNECT"}
	fmt.Printf(moveCursor, 8, 1)
	fmt.Println("Select Request Type:")
	for i, rt := range requestTypes {
		fmt.Printf(" %d. %s\n", i+1, rt)
	}

	fmt.Printf(moveCursor, len(requestTypes)+8+1, 1)
	fmt.Print("Enter the number corresponding to the request type: ")
	var requestTypeIndex int
	fmt.Scanf("%d", &requestTypeIndex)
	requestType := requestTypes[requestTypeIndex-1]

	return requestType
}

func inputAuth(reader *bufio.Reader, data *requestData) string {
	fmt.Printf(moveCursor, 8, 1)
	fmt.Print("Enter Basic Auth credentials (username:password) or leave blank: ")
	auth, _ := reader.ReadString('\n')
	auth = strings.TrimSpace(auth)

	return auth
}

func inputHeaders(reader *bufio.Reader, data *requestData) string {
	fmt.Printf(moveCursor, 8, 1)
	fmt.Print("Enter Headers (Key:Value, separated by ';' or leave blank): ")
	headers, _ := reader.ReadString('\n')
	headers = strings.TrimSpace(headers)

	return headers
}

func inputBody(reader *bufio.Reader, data *requestData) string {
	fmt.Printf(moveCursor, 8, 1)
	fmt.Print("Enter Body or leave blank: ")
	body, _ := reader.ReadString('\n')
	body = strings.TrimSpace(body)

	return body
}

func sendRequest(protocol, url, requestType, auth, headers, body string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(requestType, fmt.Sprintf("%s://%s", protocol, url), strings.NewReader(body))
	if err != nil {
		return "", err
	}

	// Set authentication
	if auth != "" {
		authEncoded := base64.StdEncoding.EncodeToString([]byte(auth))
		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", authEncoded))
	}

	// Parse and set headers
	headerPairs := strings.Split(headers, ";")
	for _, pair := range headerPairs {
		if pair == "" {
			continue
		}

		kv := strings.Split(pair, ":")
		if len(kv) != 2 {
			return "", fmt.Errorf("invalid header format")
		}
		req.Header.Set(strings.TrimSpace(kv[0]), strings.TrimSpace(kv[1]))
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
