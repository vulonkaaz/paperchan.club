package themagicpipe

import (
	"os/exec"
	"encoding/base64"
	"errors"
	"bytes"
	"strings"
)

// Where the magic happens, the base64 data url image is sent through Imagemagick
// and converted to indexed colors 400x200 png
func DataURLConverter(dataURL string) (string, error) {
	if !strings.HasPrefix(dataURL, "data:image/png;base64,") {
		return "", errors.New("invalid dataURL")
	}
	imageBase64 := strings.TrimPrefix(dataURL, "data:image/png;base64,")
	imageRaw, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		return "", err
	}

	converter := exec.Command("convert", "-", "-background", "white", "-flatten", "-resize", "400x200!", "-remap", "palette.png", "PNG8:-")
	converter.Stdin = bytes.NewBuffer(imageRaw)

	var output []byte
	output, err = converter.Output()
	if err != nil {
		return "", err
	}

	imageBase64Processed := base64.StdEncoding.EncodeToString(output)
	return "data:image/png;base64," + imageBase64Processed, nil
}
