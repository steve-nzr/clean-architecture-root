package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Movie entity validation", func() {
	var (
		movie Movie
		err   error
		v     = validator.New()

		validMovie = Movie{
			Entity: Entity{
				ID: "c505bbd9-0082-49b7-954b-7702118a2480",
			},
			Title:      "Harry Potter",
			ReleasedAt: time.Now(),
		}
	)

	BeforeEach(func() {
		movie = validMovie
	})

	When("The movie entity contains invalid data", func() {
		It("should be rejected if its id is not a uuid v4", func() {
			movie.Entity.ID = ""

			err = movie.Validate(v)
			Expect(err).Should(Not(BeNil()))
		})

		It("should be rejected if its title is empty", func() {
			movie.Title = ""

			err = movie.Validate(v)
			Expect(err).Should(Not(BeNil()))
		})

		It("should be rejected if its release data isn't set", func() {
			movie.ReleasedAt = time.Time{}

			err = movie.Validate(v)
			Expect(err).Should(Not(BeNil()))
		})
	})
})
