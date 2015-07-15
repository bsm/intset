package intset

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set", func() {
	var subject *Set

	BeforeEach(func() {
		subject = New(5)
		Expect(subject.Add(2)).To(BeTrue())
		Expect(subject.Add(4)).To(BeTrue())
		Expect(subject.Add(6)).To(BeTrue())
	})

	It("should have len", func() {
		Expect(subject.Len()).To(Equal(3))
	})

	It("should add data", func() {
		Expect(subject.Add(3)).To(BeTrue())
		Expect(subject.Add(1)).To(BeTrue())
		Expect(subject.Len()).To(Equal(5))

		Expect(subject.Add(2)).To(BeFalse())
		Expect(subject.Add(3)).To(BeFalse())
		Expect(subject.Add(4)).To(BeFalse())
		Expect(subject.Len()).To(Equal(5))
	})

	It("should remove data", func() {
		Expect(subject.Remove(3)).To(BeFalse())
		Expect(subject.Len()).To(Equal(3))
		Expect(subject.Remove(2)).To(BeTrue())
		Expect(subject.Len()).To(Equal(2))
		Expect(subject.Remove(2)).To(BeFalse())
		Expect(subject.Len()).To(Equal(2))
	})

	It("should check if exists", func() {
		Expect(subject.Exists(1)).To(BeFalse())
		Expect(subject.Exists(2)).To(BeTrue())
		Expect(subject.Exists(3)).To(BeFalse())
		Expect(subject.Exists(4)).To(BeTrue())
	})

	It("should check for intersections", func() {
		oth := New(3)
		Expect(subject.Intersects(oth)).To(BeFalse())

		oth.Add(3)
		oth.Add(5)
		Expect(subject.Intersects(oth)).To(BeFalse())

		oth.Add(7)
		oth.Add(4)
		Expect(subject.Intersects(oth)).To(BeTrue())
	})

})

// --------------------------------------------------------------------

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "intset")
}
