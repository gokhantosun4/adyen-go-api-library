/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/online-payments).  ## Authentication Each request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v67/payments ```
 *
 * API version: 67
 * Contact: developer-experience@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkout

// Mandate struct for Mandate
type Mandate struct {
	// The billing amount (in minor units) of the recurring transactions..
	Amount string `json:"amount"`
	// The limitation rule of the billing amount. Allowed values: * `max` * `exact`
	AmountRule string `json:"amountRule,omitempty"`
	// The rule to specify the period, within which the recurring debit can happen, relative to the mandate recurring date. Allowed values: * `on` * `before` * `after`
	BillingAttemptsRule string `json:"billingAttemptsRule,omitempty"`
	// The number of the day, on which the recurring debit can happen. Should be within the same calendar month as the mandate recurring date. Possible values: 1-31 based on the frequency.
	BillingDay string `json:"billingDay,omitempty"`
	// End date of the billing plan, in YYYY-MM-DD format.
	EndsAt string `json:"endsAt"`
	// The frequency with which a shopper should be charged. Allowed values: * `daily` * `weekly` * `biWeekly` * `monthly` * `quarterly` * `halfYearly` * `yearly` * `adhoc`.
	Frequency string `json:"frequency"`
	// The message shown by UPI to the shopper on the approval screen.
	Remarks string `json:"remarks,omitempty"`
	// Start date of the billing plan, in YYYY-MM-DD format. By default, the transaction date.
	StartsAt string `json:"startsAt,omitempty"`
}
