package themagicpipe

import (
	"encoding/base64"
	"errors"
	"gopkg.in/gographics/imagick.v3/imagick"
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
	// doing conversion inline soothes the soul
	imagick.Initialize()
	defer imagick.Terminate()
	magick := imagick.NewMagickWand()
	defer magick.Destroy()
	// read in from base64
	err = magick.ReadImageBlob(imageRaw)
	if err != nil {
		return "", err
	}
	// swap all transparent pixels with white
	transparent := imagick.NewPixelWand()
	defer transparent.Destroy()
	transparent.SetColor("none")
	whiteFill := imagick.NewPixelWand()
	defer whiteFill.Destroy()
	whiteFill.SetColor("#ffffff")
	err = magick.OpaquePaintImage(transparent, whiteFill, 5, false)
	if err != nil {
		return "", err
	}
	err = magick.ResizeImage(400, 200, imagick.FILTER_POINT)
	if err != nil {
		return "", err
	}
	err = magick.SetCompression(imagick.COMPRESSION_LZMA)
	if err != nil {
		return "", err
	}
	err = magick.StripImage()
	if err != nil {
		return "", err
	}
	err = magick.SetDepth(8)
	if err != nil {
		return "", err
	}
	err = magick.QuantizeImage(4, imagick.COLORSPACE_UNDEFINED, 0, imagick.DITHER_METHOD_RIEMERSMA, false)
	if err != nil {
		return "", err
	}
	imageBytesProcessed, err := magick.GetImageBlob()
	if err != nil {
		return "", err
	}
	imageBase64Processed := base64.StdEncoding.EncodeToString(imageBytesProcessed)
	return "data:image/png;base64," + imageBase64Processed, nil
}
