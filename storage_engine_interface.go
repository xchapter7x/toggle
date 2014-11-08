package toggle

type storageEngine interface {
	//SetFeatureList(featureList map[string]*feature) (err error)
	//RefreshAll() (err error)
	//RefreshOne(featureName string) (err error)
	//ResetOne(featureName string) (err error)
	//Subscribe() (err error)
	//UnSubscribe() (err error)
	GetFeatureStatusValue(featureSignature string) (status string, err error)
}
