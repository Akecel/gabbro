package utils

import (
	"strings"
)

func ReconstructImgURL(url string) string {
	baseURL := "https://images.igdb.com/igdb/image/upload/t_cover_big/"
	splitURL := strings.Split(url, "/")
	imgName := splitURL[len(splitURL)-1]
	return baseURL + imgName
}
