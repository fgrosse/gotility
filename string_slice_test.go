package gotility_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fgrosse/gotility"
	"testing"
)

func BenchmarkAddBuiltinStringSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := []string{}
		s = append(s, "test")
	}
}

func BenchmarkAddStringSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := gotility.StringSlice{}
		s.Add("test")
	}
}

var _ = Describe("StringSlice", func() {
	It("should add one element", func() {
		s := gotility.StringSlice{}
		s.Add("foo")
		Expect(s).To(ConsistOf([]string{"foo"}))
	})

	It("should add all elements", func() {
		s := gotility.StringSlice{}
		s.AddAll("foo", "bar", "baz")
		Expect(s).To(ConsistOf([]string{"foo", "bar", "baz"}))
	})

	It("should delete elements by index", func() {
		s := gotility.StringSlice{}
		Expect(s.DeleteByIndex(0)).To(BeFalse())
		Expect(s.DeleteByIndex(-1)).To(BeFalse())
		Expect(s.DeleteByIndex(42)).To(BeFalse())

		s = gotility.StringSlice{"foo", "bar", "baz"}
		Expect(s).To(ConsistOf([]string{"foo", "bar", "baz"}))
		Expect(s.DeleteByIndex(1)).To(BeTrue())
		Expect(s).To(ConsistOf([]string{"foo", "baz"}))
		Expect(s.DeleteByIndex(0)).To(BeTrue())
		Expect(s).To(ConsistOf([]string{"baz"}))
		Expect(s.DeleteByIndex(0)).To(BeTrue())
		Expect(s).To(ConsistOf([]string{}))
	})

	It("should search and delete elements", func() {
		s := gotility.StringSlice{}
		Expect(s.DeleteByValue("foo")).To(BeFalse())

		s = gotility.StringSlice{"foo", "bar", "baz"}
		Expect(s).To(ConsistOf([]string{"foo", "bar", "baz"}))
		Expect(s.DeleteByValue("bar")).To(BeTrue())
		Expect(s).To(ConsistOf([]string{"foo", "baz"}))
		Expect(s.DeleteByValue("foo")).To(BeTrue())
		Expect(s).To(ConsistOf([]string{"baz"}))
		Expect(s.DeleteByValue("baz")).To(BeTrue())
		Expect(s).To(ConsistOf([]string{}))
	})

	It("should know if it contains a string", func() {
		s := gotility.StringSlice{}
		Expect(s.Contains("foo")).To(BeFalse())

		s = gotility.StringSlice{"foo", "bar"}
		Expect(s.Contains("foo")).To(BeTrue())
		Expect(s.Contains("bar")).To(BeTrue())
		Expect(s.Contains("baz")).To(BeFalse())
	})

	It("should be able to reverse its elements", func() {
		s := gotility.StringSlice{}
		s.Reverse()
		Expect(s).To(BeEmpty())

		s = gotility.StringSlice{"foo", "bar", "baz", "1", "2"}
		s.Reverse()
		Expect(s).To(Equal(gotility.StringSlice{"2", "1", "baz", "bar", "foo"}))
		s.Reverse()
		Expect(s).To(Equal(gotility.StringSlice{"foo", "bar", "baz", "1", "2"}))

		s = gotility.StringSlice{"1", "2", "3", "4"}
		s.Reverse()
		Expect(s).To(Equal(gotility.StringSlice{"4", "3", "2", "1"}))
	})
})
