package todoist

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/saimib/code-kata-be-go-solution/models"
	"github.com/saimib/code-kata-be-go-solution/utils"
)

var _ = Describe("Command", func() {
	var (
		mockController *gomock.Controller
		mockUtils      *utils.MockUtils
		todoist        Todoist
	)

	BeforeEach(func() {
		mockController = gomock.NewController(GinkgoT())
		mockUtils = utils.NewMockUtils(mockController)
		todoist = NewTodoist(mockUtils)
	})

	When("Execute is called", func() {
		It("should return list of even numbered Todos", func() {

			mockUtils.EXPECT().Fetch(gomock.Any()).DoAndReturn(func(resp any) error {
				result := resp.(*[]models.Todo)
				*result = append(*result, models.Todo{ID: 2})
				return nil
			})
			todos := todoist.GetEvenTodos()

			Expect(len(todos)).ToNot(BeZero())
		})

		It("should return error if could not fetch todos", func() {

			mockUtils.EXPECT().Fetch(gomock.Any()).Do(func(resp any) error {
				return errors.New("fetch error")
			})
			todos := todoist.GetEvenTodos()

			Expect(len(todos)).To(BeZero())
		})

		It("should return empty list if there are no even numbered todos", func() {
			mockUtils.EXPECT().Fetch(gomock.Any()).Do(func(resp any) error {
				result := resp.(*[]models.Todo)
				*result = append(*result, models.Todo{ID: 1})
				return nil
			})
			todos := todoist.GetEvenTodos()

			Expect(len(todos)).To(BeZero())
		})
	})
})
