package gotility_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/fgrosse/gotility"
)

var _ = Describe("StringSlice", func() {
	It("should add elements to the trace", func() {
		t := gotility.StringSlice{}
		t.Add("foo", "bar")
		Expect(t).To(ConsistOf([]string{"foo", "bar"}))
	})

	It("should delete elements from the trace", func() {
		t := gotility.StringSlice{}
		t.Delete("foo")
		Expect(t).To(ConsistOf([]string{}))

		t = gotility.StringSlice{"foo", "bar", "baz"}
		Expect(t).To(ConsistOf([]string{"foo", "bar", "baz"}))
		t.Delete("bar")
		Expect(t).To(ConsistOf([]string{"foo", "baz"}))
		t.Delete("foo")
		Expect(t).To(ConsistOf([]string{"baz"}))
		t.Delete("baz")
		Expect(t).To(ConsistOf([]string{}))
	})

	It("should know if it contains a string", func() {
		t := gotility.StringSlice{}
		Expect(t.Contains("foo")).To(BeFalse())

		t = gotility.StringSlice{"foo", "bar"}
		Expect(t.Contains("foo")).To(BeTrue())
		Expect(t.Contains("bar")).To(BeTrue())
		Expect(t.Contains("baz")).To(BeFalse())
	})

	It("should be able to reverse its elements", func() {
		t := gotility.StringSlice{}
		t.Reverse()
		Expect(t).To(BeEmpty())

		t = gotility.StringSlice{"foo", "bar", "baz", "1", "2"}
		t.Reverse()
		Expect(t).To(Equal(gotility.StringSlice{"2", "1", "baz", "bar", "foo"}))
		t.Reverse()
		Expect(t).To(Equal(gotility.StringSlice{"foo", "bar", "baz", "1", "2"}))

		t = gotility.StringSlice{"1", "2", "3", "4"}
		t.Reverse()
		Expect(t).To(Equal(gotility.StringSlice{"4", "3", "2", "1"}))
	})
})
