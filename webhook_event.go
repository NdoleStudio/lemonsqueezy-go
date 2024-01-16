package lemonsqueezy

const (
	// WebhookEventOrderCreated occurs when a new order is successfully placed.
	WebhookEventOrderCreated = "order_created"

	// WebhookEventOrderRefunded occurs when a full or partial refund is made on an order.
	WebhookEventOrderRefunded = "order_refunded"

	// WebhookEventSubscriptionCreated occurs when a new subscription is successfully created.
	WebhookEventSubscriptionCreated = "subscription_created"

	// WebhookEventSubscriptionUpdated occurs when a subscription's data is changed or updated.
	WebhookEventSubscriptionUpdated = "subscription_updated"

	// WebhookEventSubscriptionCancelled occurs when a subscription is cancelled manually by the customer or store owner.
	WebhookEventSubscriptionCancelled = "subscription_cancelled"

	// WebhookEventSubscriptionResumed occurs when a subscription is resumed after being previously cancelled.
	WebhookEventSubscriptionResumed = "subscription_resumed"

	// WebhookEventSubscriptionExpired occurs when a subscription has ended after being previously cancelled, or once dunning has been completed for past_due subscriptions.
	WebhookEventSubscriptionExpired = "subscription_expired"

	// WebhookEventSubscriptionPaused occurs when a subscription's payment collection is paused.
	WebhookEventSubscriptionPaused = "subscription_paused"

	// WebhookEventSubscriptionUnpaused occurs when a subscription's payment collection is resumed after being previously paused.
	WebhookEventSubscriptionUnpaused = "subscription_unpaused"

	// WebhookEventSubscriptionPaymentSuccess occurs when a subscription payment is successful.
	WebhookEventSubscriptionPaymentSuccess = "subscription_payment_success"

	// WebhookEventSubscriptionPaymentFailed occurs when a subscription renewal payment fails.
	WebhookEventSubscriptionPaymentFailed = "subscription_payment_failed"

	// WebhookEventSubscriptionPaymentRecovered occurs when a subscription has a successful payment after a failed payment.
	WebhookEventSubscriptionPaymentRecovered = "subscription_payment_recovered"

	// WebhookEventSubscriptionPaymentRefunded occurs when a subscription payment is refunded.
	WebhookEventSubscriptionPaymentRefunded = "subscription_payment_refunded"

	// WebhookEventLicenseKeyCreated occurs when a license key is created from a new order.
	WebhookEventLicenseKeyCreated = "license_key_created"

	// WebhookEventLicenseKeyUpdated occurs when a license key is updated.
	WebhookEventLicenseKeyUpdated = "license_key_updated"
)
