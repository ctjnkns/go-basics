package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type xkcd struct {
	Num        int    `json:"num"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "no file given")
		os.Exit(-1)
	}

	fn := os.Args[1]

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "no search term")
		os.Exit(-1)
	}

	var (
		items []xkcd
		terms []string
		input io.ReadCloser
		cnt   int
		err   error
	)

	if input, err = os.Open(fn); err != nil {
		fmt.Fprintf(os.Stderr, "file does not exist: %s\n", err)
		os.Exit(-1)
	}

	//decode the file
	if err = json.NewDecoder(input).Decode(&items); err != nil {
		fmt.Fprintf(os.Stderr, "unable to decode json: %s\n", err)
		os.Exit(-1)
	}

	fmt.Fprintf(os.Stderr, "read %d comics\n", len(items))

	//get search terms
	for _, t := range os.Args[2:] {
		terms = append(terms, strings.ToLower(t)) //conver everything to lowercase for searching
	}

	//search
outer:
	for _, item := range items {
		title := strings.ToLower(item.Title)
		transcript := strings.ToLower(item.Transcript)

		for _, term := range terms {
			if !strings.Contains(title, term) && !strings.Contains(transcript, term) {
				continue outer //use the outer key work to skip to the next interation of the outer loop instead of continuing the inner loop
			}
		}
		fmt.Printf("https://xkcd.com/%d/ %s/%s/%s %s\n",
			item.Num, item.Month, item.Day, item.Year, item.Title) //allows you to link directly to the comic
		cnt++
	}

	fmt.Fprintf(os.Stderr, "found %d comics\n", cnt)

}
