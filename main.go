package main

import (
	"bufio"
	"cowsay-cli/main/utils"
	"fmt"
	"io"
	"os"
)

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

	lines = utils.TabsToSpaces(lines);
	maxwidth := utils.CalculateMaxWidth(lines);
	messages := utils.NormalizeStingsLength(lines, maxwidth);
	balloon := utils.BuildBallon(messages, maxwidth);

	fmt.Println(balloon);
	fmt.Println(cow);
	fmt.Println();
}