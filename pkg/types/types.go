package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д.)
type Money int64

// PaymentCategory представляет собой категорию, в коротой был савершён поатёж (авто, аптеки, рестораны и т.д.)
type PaymentCategory string

// PaymentStatus представляет собой статус платежа.
type PaymentStatus string

// Currency представляет код валюты
type Currency string

// Код валюты
const(
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// Предопределённые статусы платежей.
const (
	PaymentStatusOk			PaymentStatus = "OK"
	PaymentStatusFail		PaymentStatus = "FAIL"
	PaymentStatusInProgress	PaymentStatus = "INPROGRESS"
)

// PAN представляет номер карты
type PAN string

// Payment представляет информацию о платеже.
type Payment struct{
	ID			string
	AccountID	int64
	Amount		Money
	Category	PaymentCategory
	Status		PaymentStatus
}

type Phone string

// Card представляет информацию о платёжной карте.
type Card struct {
	ID			int
	PAN			PAN
	Balance		Money // использовали Money
	MinBalance	Money // использовали Money
	Currency	Currency
	Color		string
	Name		string
	Active		bool
}

// Account представляет информацию о счёте пользователя.
type Account struct {
	ID		int64
	Phone	Phone
	Balance	Money
}

type erorr interface{
	Error() string
}

type Error string

func (e Error) Error() string {
	return string(e)
}


type Messenger interface {
	Send(message string) bool
	Reseive() (message string, ok bool)
}

type Telegram struct{

}

func (t *Telegram) Send(message string) bool {
	return true
}

func (t *Telegram) Reseive() (message string, ok bool) {
	return "", true
}


func New (text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
