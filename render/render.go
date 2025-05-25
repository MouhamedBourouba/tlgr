package render

import (
	"io"
	"os"
)

func RenderPage(writer io.Writer, pageFilePath string) (error) {
	data, err := os.ReadFile(pageFilePath)
	if err != nil {
		return err
	}

	writer.Write(data)

	return nil
}
