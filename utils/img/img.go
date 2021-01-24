package img

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
	"time"
)

func FromBase64ToImage(base64img string, nombre string, pathLocal string) (string, error) {
	if !strings.Contains(base64img, ",") {
		return "", errors.New("Error en formato de imagen")
	}
	listBase64 := strings.Split(base64img, ",")
	formato := listBase64[0]
	img := listBase64[1]
	unbased, err := base64.StdEncoding.DecodeString(img)
	if err != nil {
		return "", err
	}

	r := bytes.NewReader(unbased)
	var im image.Image
	extension := ""
	if formato == "data:image/png;base64" {
		im, err = png.Decode(r)
		extension = ".png"

	} else if formato == "data:image/jpg;base64" || formato == "data:image/jpeg;base64" {
		im, err = jpeg.Decode(r)
		extension = ".jpeg"
	} else {
		return "", errors.New("Error en formato de imagen")
	}

	if err != nil {
		return "", errors.New("Error al decodificar iamgen")
	}
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	path := "public/img/" + pathLocal + nombre + " " + currentTime + extension

	path = strings.ReplaceAll(path, " ", "-")
	path = strings.ReplaceAll(path, ":", "-")
	f, err := os.Create(path)
	if err != nil {
		return "", err
	}
	if extension == ".png" {
		err = png.Encode(f, im)
	} else {
		err = jpeg.Encode(f, im, nil)
	}

	if err != nil {
		return "", errors.New("Error al guardar imagen")
	}
	return path, err
}
