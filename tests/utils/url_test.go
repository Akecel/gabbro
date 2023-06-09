package tests

import (
	"testing"

	"github.com/akecel/gabbro/utils"
)

func TestReconstructImgURL(t *testing.T) {
	url := "//images.igdb.com/igdb/image/upload/t_thumb/co65ac.jpg"
	expected := "https://images.igdb.com/igdb/image/upload/t_cover_big/co65ac.jpg"
	result := utils.ReconstructImgURL(url)

	if result != expected {
		t.Errorf("Result=%s; Expected=%s", result, expected)
	}
}
