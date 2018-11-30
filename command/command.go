package command

import (
	"fmt"
	"github.com/Maki-Daisuke/go-lines"
	"github.com/codegangsta/cli"
	"os"
	"regexp"
	"strings"
)

var Flags = []cli.Flag{
	cli.StringFlag{
		Name:  "regex",
		Value: "^.*$",
		Usage: "regex filter",
	},
	cli.StringFlag{
		Name:  "implode",
		Value: "\n",
		Usage: "implode multiline charactor",
	},
}

func Command(c *cli.Context) error {
	re := regexp.MustCompile(c.String("regex"))
	lineChan, errChan := lines.LinesWithError(os.Stdin)

	var outputs []string

	for line := range lineChan {
		if re.MatchString(line) {
			outputs = append(outputs, line)
		}
	}
	if err := <-errChan; err != nil {
		return err
	}
	fmt.Println(strings.Join(outputs, c.String("implode")))
	return nil
}
