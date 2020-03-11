package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//FreePlanLimit limit of chrackter in free plan
const FreePlanLimit = 50000

type stop struct {
	error
}

func main() {

	LANGAUGE := "en-US"
	URL := "http://api.grammarbot.io/v2/check"

	botToken := flag.String("token", "XYZ", "Grammarbot token")
	pathToFile := flag.String("path", "", "Path to file")
	flag.Parse()

	text, err := LoadFile(*pathToFile)
	if err != nil {
		fmt.Println(err)
	}
	err = retry(3, time.Second*3, func() error {
		return CheckText(LANGAUGE, URL, *botToken, text)
	})
	if err != nil {
		fmt.Printf("checkText error %v", err)
	}

}

//LoadFile load file and check against planlimit
func LoadFile(path string) (string, error) {

	pwd, err := os.Getwd()
	defer func() {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal panic: %v", err)
			os.Exit(1)
		}
	}()

	content, err := ioutil.ReadFile(pwd + "/" + path)
	defer func() {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal panic: %v", err)
			os.Exit(1)
		}
	}()

	text := string(content)
	defer func() {
		if len(text) > FreePlanLimit {
			fmt.Fprintf(os.Stderr, "Test is to long: %d", len(text))
			os.Exit(1)
		}
	}()

	return text, nil
}

//CheckText send text to grammary
func CheckText(lang, url, token, text string) error {

	var client http.Client
	var data ResponseStruct
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("api_key", token)
	q.Add("language", lang)
	q.Add("text", text)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("Internal GrammaryBot Error")
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return err
	}

	x, err := json.MarshalIndent(data.Matches, "", "\t")
	if err != nil {
		return err
	}

	// empty len((string(x)) == 2
	if len(string(x)) <= 2 {
		fmt.Println("Text is OK")
	} else {
		fmt.Println(string(x))
	}

	return nil
}

// to avoid Internal Server Error from GrammaryBot side
func retry(attempts int, sleep time.Duration, fn func() error) error {

	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			return s.error
		}

		if attempts--; attempts > 0 {
			fmt.Printf("Take a try: %d", attempts)
			time.Sleep(sleep)
			return retry(attempts, 2*sleep, fn)
		}
		return err
	}
	return nil
}
