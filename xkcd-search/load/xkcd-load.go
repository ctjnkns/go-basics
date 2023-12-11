package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

// returns metadata for one comic by number
func getComic(i int) []byte {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprint(os.Stderr, "Can't read %s\n", err)
		os.Exit(-1) //unexpected error
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "skipping %d: got %d\n", i, resp.StatusCode)
		return nil //return nil if we got a 404, this will cause 'fail' to increment in main.
	}

	body, err := io.ReadAll(resp.Body) //buffering not needed since the body is small

	if err != nil {
		fmt.Fprint(os.Stderr, "invalid body: %s\n", err)
		os.Exit(-1) //unexepcted error
	}

	return body
}

func main() {
	var (
		output io.WriteCloser = os.Stdout //set the default output to stdout, can be changed with cli arg
		err    error
		cnt    int
		fails  int
		data   []byte
	)

	if len(os.Args) > 0 {
		output, err = os.Create(os.Args[1]) //create a file for the output

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		defer output.Close() //make sure the file is closed when the function exits
	}

	//output will be in JSON array
	//add brackets before and after
	fmt.Fprint(output, "[")
	defer fmt.Fprint(output, "]")

	for i := 1; fails < 2; i++ { //loop unti we have 2 consecutive fails
		if data = getComic(i); data == nil {
			fails++
			continue
		}

		if cnt > 0 {
			fmt.Fprint(output, ",") //OB1, don't add the comma for the first one
		}

		_, err = io.Copy(output, bytes.NewBuffer(data))

		if err != nil {
			fmt.Fprintf(os.Stderr, "error copying: %s\n", err)
			os.Exit(-1) //error copying
		}

		fails = 0 //reset fails; we only want to stop if there are 2 back to back
		cnt++
	}

	fmt.Fprintf(os.Stderr, "read %d comics\n", cnt)
}
