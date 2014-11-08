package toggle

import (
	"fmt"
	"os"
)

func NewDefaultEngine() (engine storageEngine) {
	engine = &DefaultEngine{
		Getenv: os.Getenv,
	}
	return
}

type DefaultEngine struct {
	Getenv func(string) string
}

func (s *DefaultEngine) GetFeatureStatusValue(featureSignature string) (status string, err error) {
	status = s.Getenv(featureSignature)

	if status == "" {
		err = fmt.Errorf("toggle value not set")
	}
	return
}
