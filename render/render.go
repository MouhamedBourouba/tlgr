package render

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

func renderPageToString(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}

	stringBuilder := strings.Builder{}
	scanner := bufio.NewScanner(file)

	expression, err := regexp.Compile("{{.+?}}")
	if err != nil {
		panic("Bad regular expression")
	}

	for i := 1; scanner.Scan(); i++ {
		if i == 1 {
			continue
		}

		if strings.HasPrefix(scanner.Text(), "> ") {
			res := strings.Replace(scanner.Text(), "> ", "", 1)
			res = color.HiRedString(res)
			res = fmt.Sprintf("  %s\n", res)

			stringBuilder.WriteString(res)
			continue
		}

		if strings.HasPrefix(scanner.Text(), "- ") {
			res := strings.Replace(scanner.Text(), "- ", "", 1)
			res = color.HiBlueString(res)
			res = fmt.Sprintf("  %s\n", res)

			stringBuilder.WriteString(res)
			continue
		}

		if strings.HasPrefix(scanner.Text(), "`") {
			res := strings.Trim(scanner.Text(), "`")

			if expression.Match([]byte(res)) {
				res = expression.ReplaceAllStringFunc(res, func(s string) string {
					trimmed := strings.TrimPrefix(strings.TrimSuffix(s, "}}"), "{{")
					underliner := color.New(color.Underline)
					return underliner.Sprint(trimmed)
				})
			}

			codeColor := color.New(color.FgHiMagenta, color.Bold)
			res = codeColor.Sprint(res)

			res = fmt.Sprintf("    %s\n", res)

			stringBuilder.WriteString(res)
			continue
		}

		stringBuilder.Write(scanner.Bytes())
		stringBuilder.WriteString("\n")
	}

	stringBuilder.WriteRune('\n')

	return stringBuilder.String(), nil
}

func RenderPage(writer io.Writer, pageFilePath string) error {
	page, err := renderPageToString(pageFilePath)
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(page))
	if err != nil {
		return err
	}

	return nil
}
