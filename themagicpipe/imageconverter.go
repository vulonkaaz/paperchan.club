package themagicpipe
import (
	"os/exec"
	"strings"
	"errors"
)

// Where the magic happens, the base64 data url image is sent through Imagemagick
// and converted to indexed colors 400x200 png
func DataURLConverter(dataURL string) (string, error) {
	if !strings.HasPrefix(dataURL, "data:image/png;base64,") {
		return "", errors.New("invalid dataURL")
	}
	cmd := "base64 -d | convert - -background white -flatten -resize 400x200! -colors 4 PNG8:- | base64 -w0"
   converter := exec.Command("sh","-c",cmd)
	converter.Stdin = strings.NewReader(strings.TrimPrefix(dataURL, "data:image/png;base64,"))
	output, err := converter.Output()
	if err != nil {
		return "", err
	}
	return "data:image/png;base64,"+string(output), nil
}

