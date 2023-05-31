package users_test

import (
	"encoding/json"
	"fmt"

	"github.com/golang/mock/gomock"
	v1 "github.com/happilymarrieddad/nats-api-playground/api/internal/api/v1"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/auth"
	natspkg "github.com/happilymarrieddad/nats-api-playground/api/internal/nats"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/repos"
	repomocks "github.com/happilymarrieddad/nats-api-playground/api/internal/repos/mocks"
	"github.com/happilymarrieddad/nats-api-playground/api/types"
	"github.com/nats-io/nats.go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NATS: Users", func() {

	var (
		ctrl *gomock.Controller

		globalRepo *repomocks.MockGlobalRepo
		usersRepo  *repomocks.MockUsers

		natsServerClient natspkg.Client
		natsReqClient    natspkg.Client
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		globalRepo = repomocks.NewMockGlobalRepo(ctrl)
		usersRepo = repomocks.NewMockUsers(ctrl)

		var err error
		natsServerClient, err = natspkg.NewClient(nats.DefaultURL, "usr", "pass")
		Expect(err).To(BeNil())
		natsServerClient.SetDebug(true)

		natsReqClient, err = natspkg.NewClient(nats.DefaultURL, "usr", "pass")
		Expect(err).To(BeNil())

		Expect(globalRepo).NotTo(BeNil())
		Expect(usersRepo).NotTo(BeNil())

		globalRepo.EXPECT().Users().Return(usersRepo).AnyTimes()

		v1.SetupRoutes(globalRepo, natsServerClient)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("users.index", func() {
		var token string
		var ret []*types.User

		BeforeEach(func() {
			var err error
			token, err = auth.CreateToken(map[string]interface{}{
				"user": &types.User{
					ID: 1,
				},
			})
			Expect(err).To(BeNil())

			ret = []*types.User{
				{
					ID:        1,
					FirstName: "nick",
					LastName:  "kot",
					Email:     "test@mail.com",
				},
				{
					ID:        2,
					FirstName: "jack",
					LastName:  "burgendy",
					Email:     "test2@mail.com",
				},
			}
		})

		It("should return an error from the auth'd handler", func() {
			res, err := natsReqClient.Request("users.index", []byte(`{"limit": 10, "offset": 0}`), nil)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("unauthorized"))
			Expect(res).To(BeNil())
		})

		It("should successfully return a list of users", func() {
			usersRepo.EXPECT().Find(gomock.Any(), &repos.UserFindOpts{Limit: 2}).Return(ret, int64(len(ret)), nil).Times(1)

			res, err := natsReqClient.Request("users.index", []byte(`{"limit": 10, "offset": 0}`), map[string]string{
				"token": token,
			})
			Expect(err).To(BeNil())

			bts, err := json.Marshal(struct {
				Data  interface{} `json:"data"`
				Count int         `json:"count"`
			}{
				Data:  ret,
				Count: len(ret),
			})
			Expect(err).To(BeNil())
			Expect(res).To(Equal(bts))
		})
	})

	Context("users.get.<id>", func() {
		var token string
		var usr *types.User

		BeforeEach(func() {
			usr = &types.User{
				ID:        8273,
				FirstName: "first",
				LastName:  "last",
				Email:     "somemail@test.com",
			}

			var err error
			token, err = auth.CreateToken(map[string]interface{}{
				"user": &types.User{
					ID:        usr.ID,
					FirstName: usr.FirstName,
					LastName:  usr.LastName,
					Email:     usr.Email,
				},
			})
			Expect(err).To(BeNil())
		})

		// It("should return an error from the auth'd handler", func() {
		// 	res, err := natsReqClient.Request(fmt.Sprintf("users.get.%d", usr.ID), nil, nil)
		// 	Expect(err).NotTo(BeNil())
		// 	Expect(err.Error()).To(Equal("unauthorized"))
		// 	Expect(res).To(BeNil())
		// })

		It("should successfully get a user", func() {
			usersRepo.EXPECT().Get(gomock.Any(), usr.ID).Return(usr, true, nil).AnyTimes()

			res, err := natsReqClient.Request(fmt.Sprintf("users.get.%d", usr.ID), nil, map[string]string{
				"token": token,
			})
			Expect(err).To(BeNil())

			bts, err := json.Marshal(usr)
			Expect(err).To(BeNil())

			Expect(res).To(Equal(bts))
		})
	})

	Context("users.create", func() {

	})

	Context("users.update", func() {

	})
})
