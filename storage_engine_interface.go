package toggle

type storageEngine interface {
	GetFeatureStatusValue(featureSignature string) (status string, err error)
}
