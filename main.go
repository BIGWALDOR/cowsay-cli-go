package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

// buildBallon takes in an array of strings representing the lines to display in the balloon
// and an integer value of the maximum width of the balloon. It returns a string that represents
// the balloon.
func buildBallon(lines []string, maxwidth int) string {
	var borders []string
	count := len(lines)
	var ret []string

	// Define the different borders for the balloon
	borders = []string{"/", "\\", "\\", "/", "|", "<", ">"}

	// Define the top and bottom border strings
	top := " " + strings.Repeat("_", maxwidth+2)
	bottom := " " + strings.Repeat("-", maxwidth+2)

	// Append the top border string to the return slice
	ret = append(ret, top)

	if count == 1 {
		// If there is only one line, use the left and right borders for the line
		s := fmt.Sprintf("%s %s %s", borders[5], lines[0], borders[6])
		ret = append(ret, s)
	} else {
		// If there are multiple lines, use the top and bottom borders for the first and last line
		// and use the middle border for the rest of the lines
		s := fmt.Sprintf("%s %s %s", borders[0], lines[0], borders[1])
		ret = append(ret, s)
		i := 1

		for ; i < count-1; i++ {
			s = fmt.Sprintf("%s %s %s", borders[4], lines[i], borders[4])
			ret = append(ret, s)
		}

		s = fmt.Sprintf("%s %s %s", borders[2], lines[i], borders[3])
		ret = append(ret, s)
	}

	// Append the bottom border string to the return slice
	ret = append(ret, bottom)

	// Join all the strings in the return slice with a newline separator and return the resulting string
	return strings.Join(ret, "\n")
}


// tabsToSpaces takes a slice of strings and replaces all tabs with four spaces.
// It returns the resulting slice of strings.
func tabsToSpaces(lines []string) []string {
	var ret []string;

	for _, l := range lines {
		// Replace each tab character with four spaces
		l = strings.Replace(l, "\t", "    ", -1);
		ret = append(ret, l);
	}

	return ret;
}

// calculateMaxWidth calculates the maximum width of a slice of strings
// based on the length of the longest string.
//
// It takes a slice of strings as input and returns an integer representing
// the maximum width.
func calculateMaxWidth(lines []string) int {
	w := 0; // Initialize the width variable to 0

	// Loop through each string in the slice
	for _, l := range lines {
			len := utf8.RuneCountInString(l); // Get the length of the string in runes

			// If the length is greater than the current width, set the width to the length
			if len > w {
					w = len;
			}
	}

	return w; // Return the maximum width
}

// normalizeStingsLength normalizes the length of strings in a slice to the specified maximum width
func normalizeStingsLength(lines []string, maxwidth int) []string {
	var ret []string;

	for _, l := range lines {
		// Append spaces to the end of the string to normalize its length
		s := l + strings.Repeat(" ", maxwidth - utf8.RuneCountInString(l));
		ret = append(ret, s);
	}

	return ret;
}

func main() {
	info, _ := os.Stdin.Stat();

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.");
		fmt.Println("Usage: fortune | gocowsay");

		return;
	}

	var lines []string;

	reader := bufio.NewReader(os.Stdin);

	for {
		line, _, err := reader.ReadLine();
		if err != nil && err == io.EOF {
			break;
		}

		lines = append(lines, string(line));
	}

	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

	lines = tabsToSpaces(lines);
	maxwidth := calculateMaxWidth(lines);
	messages := normalizeStingsLength(lines, maxwidth);
	balloon := buildBallon(messages, maxwidth);

	fmt.Println(balloon);
	fmt.Println(cow);
	fmt.Println();
}