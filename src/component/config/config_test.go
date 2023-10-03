package config

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/spf13/viper"
)

var _ = Describe("GetAppConfig", func() {
	It("should returns", func() {
		viper.Set("app.env", "local")

		Expect(GetAppConfig()).To(Equal(AppConfig{
			Env: "local",
		}))
	})
})

var _ = Describe("GetHTTPConfig", func() {
	It("should returns", func() {
		viper.Set("http.port", "8000")

		Expect(GetHTTPConfig()).To(Equal(HTTPConfig{
			Port: 8000,
		}))
	})
})

var _ = Describe("GetLoggingConfig", func() {
	Context("local environment", func() {
		It("should returns", func() {
			viper.Set("app.env", "local")

			Expect(GetLoggingConfig()).To(Equal(LoggingConfig{
				Formatter: LoggingTextFormatter,
				Level:     LoggingDebugLevel,
			}))
		})
	})

	Context("staging environment", func() {
		It("should returns", func() {
			viper.Set("app.env", "staging")

			Expect(GetLoggingConfig()).To(Equal(LoggingConfig{
				Formatter: LoggingJSONFormatter,
				Level:     LoggingDebugLevel,
			}))
		})
	})

	Context("production environment", func() {
		It("should returns", func() {
			viper.Set("app.env", "production")

			Expect(GetLoggingConfig()).To(Equal(LoggingConfig{
				Formatter: LoggingJSONFormatter,
				Level:     LoggingInfoLevel,
			}))
		})
	})
})
