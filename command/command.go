package command

import (
	"bytes"
	"fmt"
	"github.com/Maki-Daisuke/go-lines"
	"github.com/atotto/clipboard"
	"github.com/codegangsta/cli"
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
	body, err := clipboard.ReadAll()

	if err != nil {
		return err
	}

	re := regexp.MustCompile(c.String("regex"))
	line_chan, err_chan := lines.LinesWithError(bytes.NewReader([]byte(body)))

	var outputs []string

	for line := range line_chan {
		if re.MatchString(line) {
			outputs = append(outputs, line)
		}
	}
	if err := <-err_chan; err != nil {
		return err
	}
	fmt.Println(strings.Join(outputs, c.String("implode")))
	return nil
}
