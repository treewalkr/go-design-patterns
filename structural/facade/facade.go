package facade

import (
	"errors"
	"fmt"
)

// BankApi defines the interface for different bank APIs.
type BankApi interface {
	Login() error
	Transfer(toAcc, toBank string, amount float32) (string, error)
	Confirm(confToken string) error
	IsLoggedIn() bool
}

// BaseBankApi provides common functionality for bank APIs.
type BaseBankApi struct {
	AccountNo string
	Bank      string
	Pin       string
	Token     string
}

// Login verifies account details and generates a token.
func (b *BaseBankApi) Login(validAccount, validPin, validBank string) error {
	if b.AccountNo == validAccount && b.Pin == validPin && b.Bank == validBank {
		b.Token = "valid-token"
		return nil
	}
	return fmt.Errorf("invalid account details for bank: %s", b.Bank)
}

// IsLoggedIn checks if the token is valid.
func (b *BaseBankApi) IsLoggedIn() bool {
	return b.Token == "valid-token"
}

// BankTransferFacade encapsulates the complexity of bank transfers
type BankTransferFacade struct {
	banks map[string]func(accountNo, bank, pin string) BankApi
}

// NewBankTransferFacade creates a new facade instance with registered banks
func NewBankTransferFacade() *BankTransferFacade {
	facade := &BankTransferFacade{
		banks: make(map[string]func(accountNo, bank, pin string) BankApi),
	}

	// Register supported banks
	facade.banks["SiamBank"] = func(accountNo, bank, pin string) BankApi {
		return NewSiamBankApi(accountNo, bank, pin)
	}
	facade.banks["KBank"] = func(accountNo, bank, pin string) BankApi {
		return NewKBankApi(accountNo, bank, pin)
	}

	return facade
}

// Transfer provides a simplified interface for transferring funds
func (f *BankTransferFacade) Transfer(accountNo, pin, fromBank, toAcc, toBank string, amount float32) error {
	// Get bank factory function
	bankFactory, exists := f.banks[fromBank]
	if !exists {
		return fmt.Errorf("unsupported bank: %s", fromBank)
	}

	// Create bank API instance
	bankApi := bankFactory(accountNo, fromBank, pin)

	// Execute transfer workflow
	if err := bankApi.Login(); err != nil {
		return fmt.Errorf("login failed: %w", err)
	}

	confToken, err := bankApi.Transfer(toAcc, toBank, amount)
	if err != nil {
		return fmt.Errorf("transfer failed: %w", err)
	}

	if err := bankApi.Confirm(confToken); err != nil {
		return fmt.Errorf("confirmation failed: %w", err)
	}

	return nil
}

// SiamBankApi implements BankApi for SiamBank
type SiamBankApi struct {
	BaseBankApi
}

func NewSiamBankApi(accountNo, bank, pin string) *SiamBankApi {
	return &SiamBankApi{BaseBankApi{AccountNo: accountNo, Bank: bank, Pin: pin}}
}

func (s *SiamBankApi) Login() error {
	return s.BaseBankApi.Login("0123456789", "1234", "SiamBank")
}

func (s *SiamBankApi) Transfer(toAcc, toBank string, amount float32) (string, error) {
	if !s.IsLoggedIn() {
		return "", errors.New("please login first")
	}
	fmt.Printf("Transferring %.2f from SiamBank to %s at %s\n", amount, toAcc, toBank)
	return "valid-confirm-token", nil
}

func (s *SiamBankApi) Confirm(confToken string) error {
	if !s.IsLoggedIn() {
		return errors.New("please login first")
	}
	if confToken != "valid-confirm-token" {
		return errors.New("invalid confirmation token")
	}
	fmt.Println("Transfer confirmed via SiamBank")
	return nil
}

// KBankApi implements BankApi for KBank
type KBankApi struct {
	BaseBankApi
}

func NewKBankApi(accountNo, bank, pin string) *KBankApi {
	return &KBankApi{BaseBankApi{AccountNo: accountNo, Bank: bank, Pin: pin}}
}

func (k *KBankApi) Login() error {
	return k.BaseBankApi.Login("0123456789", "1234", "KBank")
}

func (k *KBankApi) Transfer(toAcc, toBank string, amount float32) (string, error) {
	if !k.IsLoggedIn() {
		return "", errors.New("please login first")
	}
	fmt.Printf("Transferring %.2f from KBank to %s at %s\n", amount, toAcc, toBank)
	return "valid-confirm-token", nil
}

func (k *KBankApi) Confirm(confToken string) error {
	if !k.IsLoggedIn() {
		return errors.New("please login first")
	}
	if confToken != "valid-confirm-token" {
		return errors.New("invalid confirmation token")
	}
	fmt.Println("Transfer confirmed via KBank")
	return nil
}

// Example usage:
func Example() {
	facade := NewBankTransferFacade()

	// Perform transfer using the facade
	err := facade.Transfer("0123456789", "1234", "SiamBank", "9876543210", "KBank", 1000.00)
	if err != nil {
		fmt.Printf("Transfer failed: %v\n", err)
		return
	}
}
