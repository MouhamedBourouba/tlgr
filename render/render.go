package render

import (
	"io"
	"os"
)

func RenderPage(writer io.Writer, pageFilePath string) error {
	data, err := os.ReadFile(pageFilePath)
	if err != nil {
		return err
	}

	_, err = writer.Write(data)
	if err != nil {
		return err
	}

	return nil
}
