package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/antelman107/net-wait-go/wait"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/phayes/freeport"
	"github.com/redis/go-redis/v9"
)

func TestSqlettus(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sqlettus Suite")
}

var _ = Describe("CLI", func() {
	It("can start the server", func() {
		port, err := freeport.GetFreePort()
		Expect(err).NotTo(HaveOccurred())

		cli := &CLI{
			Port:     uint(port),
			Filename: "file:test.db?cache=shared&mode=memory",
			Workers:  1,
		}
		go func() {
			defer GinkgoRecover()

			err := cli.Run()
			Expect(err).NotTo(HaveOccurred())
		}()

		ok := wait.New().Do([]string{fmt.Sprintf("localhost:%d", port)})
		Expect(ok).To(BeTrue())

		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("localhost:%d", port),
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		Expect(client.Ping(context.Background()).Err()).NotTo(HaveOccurred())
		Expect(client.Set(context.Background(), "name", "value", time.Minute).Err()).NotTo(HaveOccurred())
		Expect(client.Get(context.Background(), "name").Err()).NotTo(HaveOccurred())
		Expect(client.FlushAll(context.Background()).Err()).NotTo(HaveOccurred())
	})
})
