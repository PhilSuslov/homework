package integrations

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	inventory_v1 "github.com/PhilSuslov/homework/shared/pkg/proto/inventory/v1"
)

var _ = Describe("InventoryService", func(){
	var (
		ctx context.Context
		cancel context.CancelFunc
		inventoryClient inventory_v1.InventoryServiceClient
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(suiteCtx)

		conn, err := grpc.NewClient(
			env.App.Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)

		Expect(err).ToNot(HaveOccurred(), "Ожидали успешное подключение к gRPC приложению")

		inventoryClient = inventory_v1.NewInventoryServiceClient(conn)
	})

	AfterEach(func(){
		err := env.ClearNoteCollection(ctx)
		Expect(err).ToNot(HaveOccurred(), "Ожидали успешную очистку коллекции note")

		cancel()
	})

	Describe("Get", func() {
		var orderUUID string
		
		BeforeEach(func() {
			var err error
			orderUUID, err = env.InsertTestNote(ctx)
			Expect(err).ToNot(HaveOccurred(), "Ожидали успешную вставку тестового GET")
		})

		It("Должен успешно возвращать order", func(){
			resp, err := inventoryClient.GetPart(ctx, &inventory_v1.GetPartRequest{
				Uuid: orderUUID,
			})

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.GetPart()).ToNot(BeNil())
			Expect(resp.GetPart().Uuid).To(Equal(orderUUID))
			Expect(resp.GetPart().Category).ToNot(BeNil())
			Expect(resp.GetPart().Price).ToNot(BeEmpty())
			Expect(resp.GetPart().Description).ToNot(BeEmpty())
			Expect(resp.GetPart().GetCreatedAt()).ToNot(BeNil())
		})
	})

	Describe("List Part", func(){
		var part *inventory_v1.ListPartsRequest

		BeforeEach(func(){
			part = env.GetListPartsNoteInfo()
		})
		It("Должен успешно возвращать part", func(){
			resp, err := inventoryClient.ListParts(ctx, &inventory_v1.ListPartsRequest{
				Filter: part.Filter,

			})

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.Parts).ToNot(BeNil())
		})
	})
})