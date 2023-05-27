package nats_test

import (
	"time"

	. "github.com/happilymarrieddad/nats-api-playground/api/internal/nats"
	"github.com/nats-io/nats.go"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("nats client", func() {

	var (
		client1 Client
		client2 Client
		client3 Client
	)

	BeforeEach(func() {
		var err error

		client1, err = NewClient(nats.DefaultURL, "usr", "pass")
		Expect(err).To(BeNil())
		Expect(client1).NotTo(BeNil())

		client2, err = NewClient(nats.DefaultURL, "usr", "pass")
		Expect(err).To(BeNil())
		Expect(client2).NotTo(BeNil())

		client3, err = NewClient(nats.DefaultURL, "usr", "pass")
		Expect(err).To(BeNil())
		Expect(client3).NotTo(BeNil())
	})

	Context("multiple clients", func() {
		It("should only respond once", func() {
			ch := "example-channel"
			data := []byte(`{"success": true}`)

			var handles int

			_, err := client1.HandleRequest(ch, "api", func(m *nats.Msg) {
				handles++
				client1.Respond(m.Reply, data)
			})
			Expect(err).To(BeNil())

			_, err = client2.HandleRequest(ch, "api", func(m *nats.Msg) {
				handles++
				client2.Respond(m.Reply, data)
			})
			Expect(err).To(BeNil())

			time.Sleep(time.Second) // simulate normal delay

			res, err := client3.Request(ch, nil, nil)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(data))

			// Should have only 1 handle
			Expect(handles).To(Equal(1))
		})
	})
})
