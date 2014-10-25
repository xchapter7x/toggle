package toggle_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xchapter7x/toggle"
)

var _ = Describe("toggle package", func() {
	controlNamespace := "hi"

	BeforeEach(func() {
		toggle.Init(controlNamespace)
	})

	Describe("RegisterFeature Function", func() {
		It("Should inject a new feature", func() {
			initialFeatureCount := len(toggle.ShowFeatures())
			featureName := "sampleFeature"
			toggle.RegisterFeature(featureName)
			currentFeatureCount := len(toggle.ShowFeatures())
			Expect(initialFeatureCount).NotTo(Equal(currentFeatureCount))
		})

		It("Should add feature record for referencing", func() {
			initialFeatureList := toggle.ShowFeatures()
			featureName := "sampleFeature"
			_, controlExists := initialFeatureList[featureName]
			toggle.RegisterFeature(featureName)
			currentFeatureList := toggle.ShowFeatures()
			_, currentExists := currentFeatureList[featureName]
			Expect(controlExists).NotTo(Equal(currentExists))
		})

		It("Should ignore duplicate register calls", func() {
			featureName := "sampleFeature"
			toggle.RegisterFeature(featureName)
			initialFeatureCount := len(toggle.ShowFeatures())
			toggle.RegisterFeature(featureName)
			currentFeatureCount := len(toggle.ShowFeatures())
			Expect(initialFeatureCount).To(Equal(currentFeatureCount))
		})

	})
})
