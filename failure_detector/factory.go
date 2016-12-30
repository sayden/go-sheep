package failure_detector

import "github.com/sayden/go-sheep"

func NewFailureDetector() go_sheep.SWIM {
	return swim{}
}
