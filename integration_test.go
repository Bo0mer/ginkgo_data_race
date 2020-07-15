package datarace_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/oauth2"
)

var _ = Describe("Data race", func() {
	BeforeEach(func() {
		config := &oauth2.Config{
			ClientID:     "cid",
			ClientSecret: "cs",
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://accounts.google.com/o/oauth2/auth",
				TokenURL: "https://accounts.google.com/o/oauth2/token",
			},
		}
		ts := oauth2.ReuseTokenSource(nil, config.TokenSource(
			oauth2.NoContext,
			&oauth2.Token{RefreshToken: "refresh token"},
		))
		_, err := ts.Token()
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("should succeed", func() {
		Ω(true).Should(BeTrue())
	})
})
