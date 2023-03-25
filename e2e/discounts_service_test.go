package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

//func randomString(length int) string {
//	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
//	b := make([]byte, length+2)
//	generator.Read(b)
//	return fmt.Sprintf("%x", b)[2 : length+2]
//}

//func TestDiscountsService_Create(t *testing.T) {
//	// Act
//	storeID := "11559"
//	discount, response, err := client.Discounts.Create(context.Background(), &lemonsqueezy.DiscountCreateParams{
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
//	//_, _ = client.Discounts.Delete(context.Background(), discount.Data.ID)
//}

//func TestDiscountsService_Get(t *testing.T) {
//	// Arrange
//	//storeID := "11559"
//	//expectedDiscount, response, err := client.Discounts.Create(context.Background(), &lemonsqueezy.DiscountCreateParams{
//	//	Name:       "10% Off",
//	//	Code:       strings.ToUpper(randomString(10)),
//	//	Amount:     10,
//	//	AmountType: "percent",
//	//	StoreID:    storeID,
//	//})
//
//	//spew.Dump(expectedDiscount.Data.ID)
//
//	// Act
//	discount, response, err := client.Discounts.Get(context.Background(), "11199") //expectedDiscount.Data.ID)
//
//	spew.Dump(discount)
//
//	// Assert
//	assert.Nil(t, err)
//
//	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
//	//assert.Equal(t, expectedDiscount.Data.ID, discount.Data.ID)
//
//	// Teardown
//	//_, _ = client.Discounts.Delete(context.Background(), discount.Data.ID)
//}
//func TestDiscountsService_Delete(t *testing.T) {
//	// Arrange
//	storeID := "11559"
//	expectedDiscount, response, err := client.Discounts.Create(context.Background(), &lemonsqueezy.DiscountCreateParams{
//		Name:       "10% Off",
//		Code:       strings.ToUpper(randomString(10)),
//		Amount:     10,
//		AmountType: "percent",
//		StoreID:    storeID,
//	})
//
//	// Act
//	response, err = client.Discounts.Delete(context.Background(), expectedDiscount.Data.ID)
//
//	// Assert
//	assert.Nil(t, err)
//
//	assert.Equal(t, http.StatusNoContent, response.HTTPResponse.StatusCode)
//}

func TestDiscountsService_List(t *testing.T) {
	// Act
	discounts, response, err := client.Discounts.List(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, 0, len(discounts.Data))
}
