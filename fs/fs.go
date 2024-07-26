package fspackage

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var banners = map[string]bool{
	"standard":   true,
	"shadow":     true,
	"thinkertoy": true,
}

var (
	bad_request           = "status 400 - Bad Request"
	internal_server_error = "500 - Internal Server Errors"
)

func Ascii_Art(text, banner string) (string, int, error) {
	if !banners[banner] {
		return "", 400, errors.New(bad_request)
	}
	if text == "" {
		return "", 200, nil
	}

	chars_indexes, err := getIndexes(text)
	if err != nil {
		return "", 500, errors.New(internal_server_error)
	}

	path_name := "static/banners/" + banner + ".txt"
	chars_map, err := GetCharacters(chars_indexes, path_name)
	if err != nil {
		return "", 500, errors.New(internal_server_error)
	}

	words := strings.Split(text, "\n")
	// this variable check if there is a newline at the end of the argument
	return Writer(words, chars_map), 200, nil
}

func getIndexes(text string) (map[int]rune, error) {
	index_map := make(map[int]rune, len(text))
	/* this loop returns the  starting  line  of each character at the ascii art file */
	for _, char := range text {
		if char < ' ' || char > '~' {
			continue
		}
		index_map[((int(char)-31)*9 - 7)] = char
	}
	return index_map, nil
}

func GetCharacters(index_map map[int]rune, banner string) (map[rune][]string, error) {
	// Open file,  intialize a bufio scanner along side some variables
	var (
		chars_map  = map[rune][]string{}
		line_num   = 1
		char_slice = []string{}
	)
	file, err := openfile(banner)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		// in eather the first line of a characterin or  the middle of reading a character
		if _, ok := index_map[line_num]; ok || len(char_slice) >= 1 {
			line := scanner.Text()
			char_slice = append(char_slice, line)
		}
		if len(char_slice) == 8 { // when we are done reading a character
			chars_map[index_map[line_num-7]] = char_slice
			char_slice = []string{}
		}
		if len(chars_map) == len(index_map) { // when we read the whole characters
			break
		}
		line_num++
	}
	return chars_map, scanner.Err()
}

func Writer(words []string, chars_map map[rune][]string) string {
	text := ""
	// this variable check if there is a newline at the end of the argument
	last_new_line := false
	for _, str := range words {
		// check if the word is a newline, print newline, and keep going
		if str == "" {
			text += "\n"
		} else {
			last_new_line = true
			text_slice := [][]string{}
			for _, char := range str {
				// collect chars from chars_map by thier indexes
				text_slice = append(text_slice, chars_map[char])
			}
			// Print ascii text
			text += writeChars(text_slice)
		}
	}
	if last_new_line {
		// this condition is for printing the last new line
		text += "\n"
	}
	return text
}

func writeChars(slice [][]string) string { // this function print banner text in the terminal
	text := ""
	for in := 0; in < 8; in++ {
		/* this loop for print banner line by line */
		for index := range slice {
			/* this loop for print all character in the line */
			if len(slice[index]) == 0 {
				continue
			}
			line := slice[index][in]
			// check if the index of the current letter in the original input is a colored letter and colored
			text += line
		}
		if in < 7 {
			/* this condition for print new line if not last line in character */
			text += "\n"
		}
	}
	return "\n" + text
}

func openfile(banner string) (*os.File, error) {
	file, err := os.Open(banner)
	if err != nil {
		/* if file has err in open return error message */
		return nil, err
	}
	return file, nil
}
