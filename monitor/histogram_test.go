package monitor_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "code.cloudfoundry.org/perm/monitor"
)

var _ = Describe("Histogram", func() {
	var (
		subject *Histogram
	)

	BeforeEach(func() {
		subject = NewHistogram(5, 1*time.Nanosecond, 1*time.Second, 5)
	})

	Describe("#Max", func() {
		It("returns the highest recorded value", func() {
			Expect(subject.Max()).To(Equal(int64(0)))

			subject.RecordValue(10)
			subject.RecordValue(12345)
			subject.RecordValue(-30)
			subject.RecordValue(678)

			Expect(subject.Max()).To(Equal(int64(12345)))
		})
	})

	Describe("#ValueAtQuantile", func() {
		It("returns the value at the given quantile", func() {
			Expect(subject.ValueAtQuantile(50)).To(Equal(int64(0)))

			subject.RecordValue(1)
			subject.RecordValue(2)
			subject.RecordValue(3)

			Expect(subject.ValueAtQuantile(84)).To(Equal(int64(3)))
			Expect(subject.ValueAtQuantile(50)).To(Equal(int64(2)))
		})
	})

	Describe("#Rotate", func() {
		It("resets the values", func() {
			Expect(subject.ValueAtQuantile(50)).To(Equal(int64(0)))

			subject.RecordValue(1)
			subject.RecordValue(2)
			subject.RecordValue(3)

			Expect(subject.ValueAtQuantile(50)).To(Equal(int64(2)))

			subject.Rotate()

			Expect(subject.ValueAtQuantile(50)).To(Equal(int64(0)))
		})
	})
})