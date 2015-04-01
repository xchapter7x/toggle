package toggle

import "github.com/xchapter7x/goutil"

func NewTgl() *Tgl {
	return new(Tgl)
}

type Tgl struct {
	flagName     string
	onFeature    interface{}
	offFeature   interface{}
	args         []interface{}
	returnValues []interface{}
}

func (s *Tgl) Flag(name string) *Tgl {
	s.flagName = name
	return s
}

func (s *Tgl) Off(feature interface{}) *Tgl {
	s.offFeature = feature
	return s
}

func (s *Tgl) On(feature interface{}) *Tgl {
	s.onFeature = feature
	return s
}

func (s *Tgl) Args(args ...interface{}) *Tgl {
	s.args = args
	return s
}

func (s *Tgl) Returns(vals ...interface{}) *Tgl {
	s.returnValues = vals
	return s
}

func (s *Tgl) Run() (res []interface{}, err error) {
	res = Flip(s.flagName,
		s.offFeature,
		s.onFeature,
		s.args...)
	err = goutil.UnpackArray(res, s.returnValues)
	return
}
