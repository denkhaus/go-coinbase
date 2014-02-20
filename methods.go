package coinbase

// Actual client functions to call

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Account changes
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Account
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) AccountBalance() (Amount, error) {
	body, err := c.Get("account/balance", nil)
	if err != nil {
		return Amount{}, err
	}

	// parse into json
	var response Amount
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Amount{}, err
	}

	return response, nil
}

func (c *Client) AccountReceiveAddress() (interface{}, error) {
	body, err := c.Get("account/receive_address", nil)
	if err != nil {
		return nil, err
	}

	type Response struct {
		Success     bool
		Address     string
		CallbackUrl string `json:"callback_url"`
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) AccountGenerateReceiveAddress(callbackURL string) (interface{}, error) {
	params := make(map[string]interface{})
	if callbackURL != "" {
		params["address"] = map[string]string{
			"callback_url": callbackURL,
		}
	}

	body, err := c.PostJSON("account/generate_receive_address", params)
	if err != nil {
		return nil, err
	}

	type Response struct {
		Success     bool
		Address     string
		CallbackUrl string `json:"callback_url"`
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Addresses
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Addresses(page int, limit int, query string) (interface{}, error) {

	params := url.Values{}
	if page != 0 {
		params.Set("page", strconv.Itoa(page))
	}

	if limit != 0 {
		params.Set("limit", strconv.Itoa(limit))
	}

	if query != "" {
		params.Set("query", query)
	}

	body, err := c.Get("addresses", params)
	if err != nil {
		return nil, err
	}

	type Address struct {
		Address struct {
			Address     string
			CallbackUrl string `json:"callback_url"`
			Label       string
			CreatedAt   string `json:"created_at"`
		}
	}

	type Response struct {
		Addresses   []Address
		TotalCount  int `json:"total_count"`
		NumPages    int `json:"num_pages"`
		CurrentPage int `json:"current_page"`
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Buttons
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Buys
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Buys(quantity float32, agree_btc_amount_varies bool) (interface{}, error) {
	params := url.Values{}
	params.Set("qty", fmt.Sprintf("%.8f", quantity))
	if agree_btc_amount_varies {
		params.Set("agree_btc_amount_varies", "true")
	}

	body, err := c.PostForm("buys", params)
	if err != nil {
		return nil, err
	}

	type CentsAmount struct {
		Cents       int
		CurrencyISO string `json:"currency_iso"`
	}

	type Transfer struct {
		ID            string
		Type          string
		Code          string
		CreatedAt     string `json:"created_at"`
		Fees          map[string]CentsAmount
		Status        string
		PayoutDate    string `json:"payout_date"`
		TransactionID string `json:"transaction_id"`
		BTC           Amount
		Subtotal      Amount
		Total         Amount
		Description   string
	}

	type Response struct {
		Success  bool
		Errors   []string
		Transfer Transfer
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Contacts
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Contacts() (interface{}, error) {
	body, err := c.Get("contacts", nil)
	if err != nil {
		return nil, err
	}

	type Contact struct {
		Email string
	}

	type Response struct {
		Contacts    []Contact
		TotalCount  int
		NumPages    int
		CurrentPage int
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Currencies
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Currencies() (interface{}, error) {
	body, err := c.Get("currencies", nil)
	if err != nil {
		return nil, err
	}

	// parse into json
	var response [][]string
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CurrenciesExchangeRates() (interface{}, error) {
	body, err := c.Get("currencies/exchange_rates", nil)
	if err != nil {
		return nil, err
	}

	type Rates map[string]string

	// parse into json
	var response Rates
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Orders
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Orders() (interface{}, error) {
	body, err := c.Get("orders", nil)
	if err != nil {
		return nil, err
	}

	type TotalCurrency struct {
		Cents       int
		CurrencyISO string
	}

	type Button struct {
		Type        string
		Name        string
		Description string
		ID          string
	}

	type Transaction struct {
		ID            string
		Hash          string
		Confirmations int
	}

	type Order struct {
		ID        string
		CreatedAt string
		Status    string

		TotalBTC    TotalCurrency
		TotalNative TotalCurrency

		Custom      string
		Button      Button
		Transaction Transaction
	}

	type Response struct {
		Orders      []Order
		TotalCount  int
		NumPages    int
		CurrentPage int
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Prices
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) PricesBuy() (interface{}, error) {
	body, err := c.Get("prices/buy", nil)
	if err != nil {
		return nil, err
	}

	type Response struct {
		SubTotal Amount
		Fees     []map[string]Amount
		Total    Amount
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) PricesSell() (interface{}, error) {
	body, err := c.Get("prices/sell", nil)
	if err != nil {
		return nil, err
	}

	type Response struct {
		SubTotal Amount
		Fees     []map[string]Amount
		Total    Amount
		Amount   string
		Currency string
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) PricesSpotRate() (Amount, error) {
	body, err := c.Get("prices/spot_rate", nil)
	if err != nil {
		return Amount{}, err
	}

	// parse into json
	var response Amount
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Amount{}, err
	}

	return response, nil
}

func (c *Client) PricesHistorical(page int) (string, error) {
	params := url.Values{}
	if page != 0 {
		params.Set("page", strconv.Itoa(page))
	}

	body, err := c.Get("prices/historical", params)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Recurring Payments
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Sells
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Sells(quantity float32) (interface{}, error) {
	params := url.Values{}
	params.Set("qty", fmt.Sprintf("%.8f", quantity))

	body, err := c.PostForm("sells", params)
	if err != nil {
		return nil, err
	}

	type CentsAmount struct {
		Cents       int
		CurrencyISO string `json:"currency_iso"`
	}

	type Transfer struct {
		ID            string
		Type          string
		Code          string
		CreatedAt     string `json:"created_at"`
		Fees          map[string]CentsAmount
		Status        string
		PayoutDate    string `json:"payout_date"`
		TransactionID string `json:"transaction_id"`
		BTC           Amount
		Subtotal      Amount
		Total         Amount
		Description   string
	}

	type Response struct {
		Success  bool
		Errors   []string
		Transfer Transfer
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Subscribers
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Tokens
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Transactions
///////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Transfers
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Transfers(page int, limit int) (interface{}, error) {
	params := url.Values{}
	if page != 0 {
		params.Set("page", strconv.Itoa(page))
	}

	if limit != 0 {
		params.Set("limit", strconv.Itoa(limit))
	}

	body, err := c.Get("transfers", params)
	if err != nil {
		return nil, err
	}

	type CentsAmount struct {
		Cents       int
		CurrencyISO string `json:"currency_iso"`
	}

	type Transfer struct {
		ID            string
		Type          string
		Code          string
		CreatedAt     string `json:"created_at"`
		Fees          map[string]CentsAmount
		Status        string
		PayoutDate    string `json:"payout_date"`
		TransactionID string `json:"transaction_id"`
		BTC           Amount
		Subtotal      Amount
		Total         Amount
		Description   string
	}

	type Response struct {
		Transfers   []map[string]Transfer
		TotalCount  int `json:"total_count"`
		NumPages    int `json:"num_pages"`
		CurrentPage int `json:"current_page"`
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////
// Users
///////////////////////////////////////////////////////////////////////////////////////////////////////

func (c *Client) Users() (interface{}, error) {
	body, err := c.Get("users", nil)
	if err != nil {
		return nil, err
	}

	type User struct {
		User struct {
			ID             string
			Name           string
			Email          string
			TimeZone       string `json:"time_zone"`
			NativeCurrency string `json:"native_currency"`
			Balance        Amount
			BuyLevel       int    `json:"buy_level"`
			SellLevel      int    `json:"sell_level"`
			BuyLimit       Amount `json:"buy_limit"`
			SellLimit      Amount `json:"sell_limit"`
		}
	}

	type Response struct {
		Users []User
	}

	// parse into json
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
