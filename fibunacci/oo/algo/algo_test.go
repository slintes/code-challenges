package algo_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/slintes/code-challenges/fibunacci/oo/algo"
)

var _ = Describe("Algo", func() {

	When("generating fibs", func() {

		var generator algo.FibGen

		test := func() {
			It("should return correct number of elements", func() {
				for i := range []int{0, 1, 2, 3, 10, 42} {
					Expect(generator.Generate(i)).To(HaveLen(i))
				}
			})

			DescribeTable("should return correct elements",
				func(nrOfElements int, elements []int) {
					Expect(generator.Generate(nrOfElements)).To(Equal(elements))
				},
				Entry("1 element", 1, []int{0}),
				Entry("2 elements", 2, []int{0, 1}),
				Entry("3 elements", 3, []int{0, 1, 1}),
				Entry("4 elements", 4, []int{0, 1, 1, 2}),
				Entry("5 elements", 5, []int{0, 1, 1, 2, 3}),
				Entry("6 elements", 6, []int{0, 1, 1, 2, 3, 5}),
			)
		}

		Context("with loop algo", func() {
			BeforeEach(func() {
				generator = algo.Loop{}
			})

			test()
		})

		Context("with recursive algo", func() {
			BeforeEach(func() {
				generator = algo.Recursive{}
			})

			test()
		})

	})

})
