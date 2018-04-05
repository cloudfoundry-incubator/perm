package apitest_test

import (
	"net"

	. "code.cloudfoundry.org/perm/pkg/api/apitest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestServer", func() {
	var (
		subject *TestServer
	)

	BeforeEach(func() {
		subject = NewTestServer()
	})

	Describe("#Serve", func() {
		It("fails if the server has already been stopped", func() {
			listener, err := net.Listen("tcp", "localhost:0")
			Expect(err).NotTo(HaveOccurred())

			defer listener.Close()

			go subject.Serve(listener)
			subject.Stop()

			err = subject.Serve(listener)
			Expect(err).To(MatchError("perm: the server has been stopped"))
		})

		It("fails when the listener is unable to accept connections", func() {
			listener, err := net.Listen("tcp", "localhost:0")
			Expect(err).NotTo(HaveOccurred())

			listener.Close()

			err = subject.Serve(listener)
			Expect(err).To(MatchError("perm: the server failed to start"))
		})
	})
})
