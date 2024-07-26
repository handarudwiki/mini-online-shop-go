package transaction

import (
	"encoding/json"
	"handarudwiki/mini-online-shop-go/infra/response"
	"time"
)

type TransactionStatus uint8

const (
	TransactionStatusCreated    TransactionStatus = 1
	TransactionStatusProgres    TransactionStatus = 10
	TransactionStatusInDelivery TransactionStatus = 15
	TransactionStatusCompleted  TransactionStatus = 20

	TRX_CREATED     string = "CREATED"
	TRX_ON_PROGRESS string = "ON PROGRESS"
	TRX_IN_DELIVERY string = "IN DELIVERY"
	TRX_COMPLETED   string = "COMPLETED"
	TRX_UNKNOWN     string = "UNKNOWN"
)

var (
	MappingTransactionStatus = map[TransactionStatus]string{
		TransactionStatusCreated:    TRX_CREATED,
		TransactionStatusInDelivery: TRX_IN_DELIVERY,
		TransactionStatusProgres:    TRX_ON_PROGRESS,
		TransactionStatusCompleted:  TRX_COMPLETED,
	}
)

type Transaction struct {
	ID           int               `db:"id"`
	UserPublicID string            `db:"user_public_id"`
	ProductID    uint              `db:"product_id"`
	ProductPrice uint              `db:"product_price"`
	Amount       uint8             `db:"amount"`
	Subtotal     uint              `db:"subtotal"`
	PlatformFee  uint              `db:"platform_fee"`
	GrandTotal   uint              `db:"grand_total"`
	Status       TransactionStatus `db:"status"`
	ProductJSON  json.RawMessage   `db:"product_snapshot"`
	Created_At   time.Time         `db:"created_at"`
	UpdatedAt    time.Time         `db:"updated_at"`
}

func NewTransactionFromReq(req CreateTransactionRequestPayload) Transaction {
	return Transaction{
		UserPublicID: req.UserPublicID,
		Amount:       req.Amount,
		Status:       TransactionStatusCreated,
		Created_At:   time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (t Transaction) Validate(stock uint8) (err error) {
	err = t.ValidateAmount()
	if err != nil {
		return
	}

	err = t.ValidateStock(stock)
	if err != nil {
		return
	}
	return
}

func (t Transaction) ValidateAmount() (err error) {
	if t.Amount == 0 {
		return response.ErrAmountInvalid
	}
	return
}

func (t Transaction) ValidateStock(stock uint8) (err error) {
	if stock < t.Amount {
		return response.ErrAmountGreaterThanStock
	}
	return
}

func (t *Transaction) SetSubtotal() {
	if t.Subtotal == 0 {
		t.Subtotal = uint(t.Amount) * t.ProductPrice
	}
}

func (t *Transaction) SetPlatformFee(platformFee uint) {
	t.PlatformFee = platformFee
}

func (t *Transaction) SetGrandTotal() {
	if t.GrandTotal == 0 {
		t.SetSubtotal()
		t.GrandTotal = t.PlatformFee + t.Subtotal
	}
}

//set product id,price, and json

func (t *Transaction) FromProduct(Product Product) *Transaction {
	t.ProductID = uint(Product.ID)
	t.ProductPrice = uint(Product.Price)

	t.SetProductJson(Product)
	return t
}

func (t *Transaction) SetProductJson(product Product) (err error) {
	productJSON, err := json.Marshal(product)
	if err != nil {
		return
	}
	t.ProductJSON = productJSON
	return
}

func (t Transaction) GetProduct() (product Product, err error) {
	err = json.Unmarshal(t.ProductJSON, &product)
	if err != nil {
		return
	}
	return
}

func (t Transaction) GetStatus() (status string) {
	status, ok := MappingTransactionStatus[t.Status]
	if !ok {
		return TRX_UNKNOWN
	}

	return
}

func (t Transaction) ToTransactionHistoryResponse() TransactionHistoryResponse {
	product, err := t.GetProduct()

	if err != nil {
		return TransactionHistoryResponse{}
	}

	return TransactionHistoryResponse{
		ID:           t.ID,
		UserPublicID: t.UserPublicID,
		ProductID:    t.ProductID,
		ProductPrice: t.ProductPrice,
		Amount:       t.Amount,
		Subtotal:     t.Subtotal,
		PlatformFee:  t.PlatformFee,
		GrandTotal:   t.GrandTotal,
		Status:       t.GetStatus(),
		CreatedAt:    t.Created_At,
		UpdatedAt:    t.UpdatedAt,
		Product:      product,
	}
}
