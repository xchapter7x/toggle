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

	Describe("IsActive function", func() {
		flagName := "bogusFlag"

		It("Should return false if given unregistered flag", func() {
			response := toggle.IsActive(flagName)
			Expect(response).To(Equal(false))
		})

		It("Should return false if given flag that is FEATURE_OFF status ", func() {
			toggle.RegisterFeature(flagName)
			toggle.SetFeatureStatus(flagName, toggle.FEATURE_OFF)
			response := toggle.IsActive(flagName)
			Expect(response).To(Equal(false))
		})

		It("Should return true if given flag that is FEATURE_ON status ", func() {
			toggle.RegisterFeature(flagName)
			toggle.SetFeatureStatus(flagName, toggle.FEATURE_ON)
			response := toggle.IsActive(flagName)
			Expect(response).To(Equal(true))
		})
	})

	Describe("SetFeatureStatus function", func() {
		flagName := "bogusFlag"

		It("Should return false if setting FEATURE_OFF status from default", func() {
			toggle.RegisterFeature(flagName)
			toggle.SetFeatureStatus(flagName, toggle.FEATURE_OFF)
			response := toggle.IsActive(flagName)
			Expect(response).To(Equal(false))
		})

		It("Should return true if setting FEATURE_ON status from default", func() {
			toggle.RegisterFeature(flagName)
			toggle.SetFeatureStatus(flagName, toggle.FEATURE_ON)
			response := toggle.IsActive(flagName)
			Expect(response).To(Equal(true))
		})

		It("Should return false if setting FEATURE_OFF status updating existing value", func() {
			toggle.RegisterFeature(flagName)
			toggle.SetFeatureStatus(flagName, toggle.FEATURE_ON)
			toggle.SetFeatureStatus(flagName, toggle.FEATURE_OFF)
			response := toggle.IsActive(flagName)
			Expect(response).To(Equal(false))
		})

		It("Should return true if setting FEATURE_ON status updating existing value", func() {
			toggle.RegisterFeature(flagName)
			toggle.SetFeatureStatus(flagName, toggle.FEATURE_OFF)
			toggle.SetFeatureStatus(flagName, toggle.FEATURE_ON)
			response := toggle.IsActive(flagName)
			Expect(response).To(Equal(true))
		})
	})

	Describe("Flip function", func() {
		flagName := "bogusFlag"

		It("Should select the default feature function to run when flag is inactive", func() {
			toggle.RegisterFeature(flagName)
			status := ""
			controlDefault := "default"
			controlNew := "new"
			toggle.Flip(flagName, func() {
				status = controlDefault
			}, func() {
				status = controlNew
			})
			Expect(status).To(Equal(controlDefault))
		})

		It("Should select the new feature function to run when flag is set to active", func() {
			toggle.RegisterFeature(flagName)
			toggle.SetFeatureStatus(flagName, toggle.FEATURE_ON)
			status := ""
			controlDefault := "default"
			controlNew := "new"
			toggle.Flip(flagName, func() {
				status = controlDefault
			}, func() {
				status = controlNew
			})
			Expect(status).To(Equal(controlNew))
		})
	})
})
