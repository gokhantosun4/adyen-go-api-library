/*
 * Adyen Recurring API
 *
 * The Recurring APIs allow you to manage and remove your tokens or saved payment details. Tokens should be created with validation during a payment request.  For more information, refer to our [Tokenization documentation](https://docs.adyen.com/checkout/tokenization). ## Authentication To connect to the Recurring API, you must use your basic authentication credentials. For this, create your web service user, as described in [How to get the WS user password](https://docs.adyen.com/user-management/how-to-get-the-web-service-ws-user-password). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@Company.YourCompany\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Recurring API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://pal-test.adyen.com/pal/servlet/Recurring/v49/disable ```
 *
 * API version: 49
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package recurring

// NotifyShopperResult struct for NotifyShopperResult
type NotifyShopperResult struct {
	// Reference of Pre-debit notification that is displayed to the shopper
	DisplayedReference string `json:"displayedReference,omitempty"`
	// A simple description of the `resultCode`.
	Message string `json:"message,omitempty"`
	// The unique reference that is associated with the request.
	PspReference string `json:"pspReference,omitempty"`
	// Reference of Pre-debit notification sent in my the merchant
	Reference string `json:"reference,omitempty"`
	// The code indicating the status of notification.
	ResultCode string `json:"resultCode,omitempty"`
	// The unique reference for the request sent downstream.
	ShopperNotificationReference string `json:"shopperNotificationReference,omitempty"`
	// Reference of Pre-debit notification that is displayed to the shopper
	StoredPaymentMethodID string `json:"storedPaymentMethodId,omitempty"`
}
