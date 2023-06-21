package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/stretchr/testify/assert"
)

//func randomString(length int) string {
//	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
//	b := make([]byte, length+2)
//	generator.Read(b)
//	return fmt.Sprintf("%x", b)[2 : length+2]
//}

//func TestWebhooksService_Create(t *testing.T) {
//	// Act
//	storeID := "11559"
//	discount, response, err := client.Webhooks.Create(context.Background(), &lemonsqueezy.WebhookCreateParams{
//		Name:       "10% Off",
//		Code:       strings.ToUpper(randomString(10)),
//		Amount:     10,
//		AmountType: "percent",
//		StoreID:    storeID,
//	})
//
//	// Assert
//	assert.Nil(t, err)
//
//	assert.Equal(t, http.StatusCreated, response.HTTPResponse.StatusCode)
//	assert.Equal(t, storeID, fmt.Sprintf("%d", discount.Data.Attributes.StoreID))
//
//	// Teardown
//	//_, _ = client.Webhooks.Delete(context.Background(), discount.Data.ID)
//}

//	func TestWebhooksService_Get(t *testing.T) {
//		// Arrange
//		//storeID := "11559"
//		//expectedWebhook, response, err := client.Webhooks.Create(context.Background(), &lemonsqueezy.WebhookCreateParams{
//		//	Name:       "10% Off",
//		//	Code:       strings.ToUpper(randomString(10)),
//		//	Amount:     10,
//		//	AmountType: "percent",
//		//	StoreID:    storeID,
//		//})
//
//		//spew.Dump(expectedWebhook.Data.ID)
//
//		// Act
//		discount, response, err := client.Webhooks.Get(context.Background(), "11199") //expectedWebhook.Data.ID)
//
//		spew.Dump(discount)
//
//		// Assert
//		assert.Nil(t, err)
//
//		assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
//		//assert.Equal(t, expectedWebhook.Data.ID, discount.Data.ID)
//
//		// Teardown
//		//_, _ = client.Webhooks.Delete(context.Background(), discount.Data.ID)
//	}
//func TestWebhooksService_Delete(t *testing.T) {
//	// Arrange
//	storeID := "11211"
//	expectedWebhook, response, err := client.Webhooks.Create(context.Background(), &lemonsqueezy.WebhookCreateParams{
//		URL:     "https://httpstat.us",
//		Events:  []string{"order_created", "subscription_created"},
//		Secret:  "SIGNING_SECRET",
//		StoreID: storeID,
//	})
//
//	if err != nil {
//		t.Error(err)
//	}
//
//	// Act
//	response, err = client.Webhooks.Delete(context.Background(), expectedWebhook.Data.ID)
//
//	// Assert
//	assert.Nil(t, err)
//
//	assert.Equal(t, http.StatusNoContent, response.HTTPResponse.StatusCode)
//}

func TestWebhooksService_List(t *testing.T) {
	// Act
	webhooks, response, err := client.Webhooks.List(context.Background())

	spew.Dump(webhooks)

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 2, len(webhooks.Data))
}
