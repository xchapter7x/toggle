package localpubsub

import (
	"os"

	"github.com/xchapter7x/toggle/engines/localengine"
	"github.com/xchapter7x/toggle/engines/storageinterface"
)

func NewLocalPubSubEngine() (engine storageinterface.StorageEngine) {
	le := &localengine.LocalEngine{
		Getenv: os.Getenv,
	}
	engine = &LocalPubSubEngine{
		LocalEngine: le,
	}
	return
}

type LocalPubSubEngine struct {
	LocalEngine *localengine.LocalEngine
}

func (s *LocalPubSubEngine) GetFeatureStatusValue(featureSignature string) (status string, err error) {
	status, err = s.LocalEngine.GetFeatureStatusValue(featureSignature)
	return
}
