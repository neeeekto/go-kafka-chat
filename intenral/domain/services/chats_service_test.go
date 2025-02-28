package services

import (
	"github.com/stretchr/testify/mock"
	"go-kafka-chat/intenral/domain/entities"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-kafka-chat/intenral/domain/contracts"
)

var _ = Describe("Chats Service", Ordered, Label("chats"), func() {
	var (
		repository *contracts.MockChatsRepository
		notifier   *contracts.MockMessageNotifier
		userRepo   *contracts.MockUsersRepository
		sut        *ChatsService
	)

	BeforeEach(func(ctx SpecContext) {
		repository = contracts.NewMockChatsRepository(GinkgoT())
		notifier = contracts.NewMockMessageNotifier(GinkgoT())
		userRepo = contracts.NewMockUsersRepository(GinkgoT())

		sut = NewChatService(repository, notifier, userRepo)
	})

	Context("Creation", func() {
		It("cant create chat for non exists users", func(ctx SpecContext) {
			userID1 := entities.NewUserID()
			userID2 := entities.NewUserID()
			userRepo.On("CheckExists", ctx, userID1, userID2).Return(false, nil).Once()

			_, err := sut.Create(ctx, "test", []entities.UserID{userID1, userID2})

			Expect(err).Should(HaveOccurred())
		})

		It("can create chat", func(ctx SpecContext) {
			name := "test"
			user1 := entities.NewUserID()

			repository.On("Save", ctx, mock.Anything).Return(nil).Once()
			userRepo.On("CheckExists", ctx, user1).Return(true, nil)

			res, err := sut.Create(ctx, name, []entities.UserID{user1})

			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.Users).Should(ContainElement(user1))
			Expect(res.Name).To(Equal(name))
		})
	})
})
