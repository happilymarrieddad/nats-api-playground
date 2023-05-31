package repos_test

import (
	"context"

	. "github.com/happilymarrieddad/nats-api-playground/api/internal/repos"
	"github.com/happilymarrieddad/nats-api-playground/api/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UsersRepo", func() {
	var repo Users
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.Background()

		repo = NewUsers(db)
		Expect(repo).NotTo(BeNil())

		clearDatabase("users")
	})

	Context("Create", func() {
		It("should successfully create a user", func() {
			usr := &types.User{
				FirstName: "Nick",
				LastName:  "Kotenberg",
				Email:     "nick@mail.com",
				Password:  "somegarbage", // this is normally a hashed value but in a test we don't care
			}

			Expect(repo.Create(ctx, usr)).To(Succeed())

			Expect(usr.ID).To(BeNumerically(">", 0))
		})
	})

	Context("Find", func() {
		BeforeEach(func() {
			Expect(repo.Create(ctx, &types.User{
				FirstName: "Nick",
				LastName:  "Kotenberg",
				Email:     "nick@mail.com",
				Password:  "somegarbage", // this is normally a hashed value but in a test we don't care
			})).To(Succeed())
			Expect(repo.Create(ctx, &types.User{
				FirstName: "Nick 2",
				LastName:  "Kotenberg",
				Email:     "nick2@mail.com",
				Password:  "somegarbage2", // this is normally a hashed value but in a test we don't care
			})).To(Succeed())
			Expect(repo.Create(ctx, &types.User{
				FirstName: "Nick 3",
				LastName:  "Kotenberg",
				Email:     "nick3@mail.com",
				Password:  "somegarbage3", // this is normally a hashed value but in a test we don't care
			})).To(Succeed())
		})

		It("should successfully get all users", func() {
			usrs, count, err := repo.Find(ctx, 2, 0)
			Expect(err).To(BeNil())
			Expect(count).To(BeNumerically("==", 3))
			Expect(usrs).To(HaveLen(2))
		})
	})

	Context("Get", func() {
		var usr *types.User
		BeforeEach(func() {
			usr = &types.User{
				FirstName: "Nick",
				LastName:  "Kotenberg",
				Email:     "nick@mail.com",
				Password:  "somegarbage", // this is normally a hashed value but in a test we don't care
			}
			Expect(repo.Create(ctx, usr)).To(Succeed())
			Expect(repo.Create(ctx, &types.User{
				FirstName: "Nick 2",
				LastName:  "Kotenberg",
				Email:     "nick2@mail.com",
				Password:  "somegarbage2", // this is normally a hashed value but in a test we don't care
			})).To(Succeed())
		})

		It("should fail to find an invalid usr", func() {
			existingUser, exists, err := repo.Get(ctx, usr.ID+1000)
			Expect(err).To(BeNil())
			Expect(exists).To(BeFalse())
			Expect(existingUser).To(BeNil())
		})

		It("should successfully find a user", func() {
			existingUser, exists, err := repo.Get(ctx, usr.ID)
			Expect(err).To(BeNil())
			Expect(exists).To(BeTrue())
			Expect(existingUser.FirstName).To(Equal(usr.FirstName))
		})
	})

	Context("Update", func() {
		var usr *types.User
		BeforeEach(func() {
			usr = &types.User{
				FirstName: "Nick",
				LastName:  "Kotenberg",
				Email:     "nick@mail.com",
				Password:  "somegarbage", // this is normally a hashed value but in a test we don't care
			}
			Expect(repo.Create(ctx, usr)).To(Succeed())
		})

		It("should successfully update", func() {
			newFirstName := "dfsdfgd"

			oldFirstName := usr.FirstName
			oldLastName := usr.LastName

			usr.FirstName = newFirstName

			newUsr, err := repo.Update(ctx, types.UserUpdate{
				ID:        usr.ID,
				FirstName: func(str string) *string { return &str }(newFirstName),
			})
			Expect(err).To(BeNil())
			Expect(newUsr).NotTo(BeNil())

			Expect(newUsr.FirstName).To(Equal(newFirstName))
			Expect(newUsr.FirstName).NotTo(Equal(oldFirstName))

			Expect(newUsr.LastName).To(Equal(oldLastName))
		})
	})

	Context("Delete", func() {
		var usr *types.User
		BeforeEach(func() {
			usr = &types.User{
				FirstName: "Nick",
				LastName:  "Kotenberg",
				Email:     "nick@mail.com",
				Password:  "somegarbage", // this is normally a hashed value but in a test we don't care
			}
			Expect(repo.Create(ctx, usr)).To(Succeed())
			Expect(repo.Create(ctx, &types.User{
				FirstName: "Nick 2",
				LastName:  "Kotenberg",
				Email:     "nick2@mail.com",
				Password:  "somegarbage2", // this is normally a hashed value but in a test we don't care
			})).To(Succeed())
		})

		It("should successfully delete the user", func() {
			usrs, count, err := repo.Find(ctx, 10, 0)
			Expect(err).To(BeNil())
			Expect(count).To(BeNumerically("==", 2))
			Expect(usrs).To(HaveLen(2))

			Expect(repo.Delete(ctx, usr.ID)).To(Succeed())

			usrs, count, err = repo.Find(ctx, 10, 0)
			Expect(err).To(BeNil())
			Expect(count).To(BeNumerically("==", 1))
			Expect(usrs).To(HaveLen(1))
		})
	})
})
