package usecase

import (
	"errors"
	"fmt"
	"github.com/watarun54/serverless-skill-manager/server/domain"
	"sort"
)

type (
	IBillRepository interface {
		FindOne(c domain.Bill) (domain.Bill, error)
		FindAll(c domain.Bill) (domain.Bills, error)
		Store(domain.Bill) (domain.Bill, error)
		Update(domain.Bill) (domain.Bill, error)
		DeleteById(domain.Bill) error
	}

	BillInteractor struct {
		BillRepository IBillRepository
		UserRepository IUserRepository
	}
)

type UserPayment struct {
	UserId int
	Amount int
}

func (interactor *BillInteractor) ConvertBillFormToBill(billForm domain.BillForm) (bill domain.Bill, err error) {
	payees, err := interactor.UserRepository.FindByIds(billForm.PayeeIds)
	if err != nil {
		return
	}
	bill.ID = billForm.ID
	bill.Name = billForm.Name
	bill.Amount = billForm.Amount
	bill.RoomID = billForm.RoomID
	bill.PayerID = billForm.PayerID
	bill.Payees = payees
	return
}

func (interactor *BillInteractor) Bill(b domain.Bill) (bill domain.Bill, err error) {
	bill, err = interactor.BillRepository.FindOne(b)
	return
}

func (interactor *BillInteractor) Bills(b domain.Bill) (bills domain.Bills, err error) {
	bills, err = interactor.BillRepository.FindAll(b)
	return
}

func (interactor *BillInteractor) UserPayments(b domain.Bill) (userPaymentsRes []domain.UserPaymentRes, err error) {
	bills, err := interactor.BillRepository.FindAll(b)
	if err != nil {
		return
	}

	users, err := interactor.UserRepository.FindAll()
	if err != nil {
		return
	}

	response := map[int]int{}
	for _, user := range users {
		response[user.ID] = 0
	}
	for _, bill := range bills {
		response[bill.Payer.ID] += bill.Amount
		for _, payee := range bill.Payees {
			response[payee.ID] -= bill.Amount / len(bill.Payees)
		}
	}

	userIdToUsers := map[int]domain.User{}
	for _, user := range users {
		userIdToUsers[user.ID] = user
	}

	userPayments := []UserPayment{}
	for userId, amount := range response {
		if amount != 0 {
			userPayments = append(userPayments, UserPayment{UserId: userId, Amount: amount})
		}
	}

	var first UserPayment
	var last UserPayment
	var diff int

	for {
		sort.SliceStable(userPayments, func(i, j int) bool { return userPayments[i].Amount < userPayments[j].Amount })
		fmt.Println("userPayments", userPayments)
		first = userPayments[0]
		userPayments = userPayments[1:]
		last = userPayments[len(userPayments)-1]
		userPayments = userPayments[:len(userPayments)-1]
		diff = first.Amount + last.Amount
		if diff > 0 {
			fmt.Printf("from: %v to: %v amount: %v\n", first.UserId, last.UserId, -(first.Amount))
			userPaymentsRes = append(userPaymentsRes, domain.UserPaymentRes{
				FromUser: userIdToUsers[first.UserId],
				ToUser:   userIdToUsers[last.UserId],
				Amount:   -(first.Amount)})
			last.Amount = diff
			userPayments = append(userPayments, last)
		} else if diff < 0 {
			fmt.Printf("from: %v to: %v amount: %v\n", last.UserId, first.UserId, last.Amount)
			userPaymentsRes = append(userPaymentsRes, domain.UserPaymentRes{
				FromUser: userIdToUsers[last.UserId],
				ToUser:   userIdToUsers[first.UserId],
				Amount:   last.Amount})
			first.Amount = diff
			userPayments = append(userPayments, first)
		} else {
			fmt.Printf("from: %v to: %v amount: %v\n", last.UserId, first.UserId, last.Amount)
			userPaymentsRes = append(userPaymentsRes, domain.UserPaymentRes{
				FromUser: userIdToUsers[last.UserId],
				ToUser:   userIdToUsers[first.UserId],
				Amount:   last.Amount})
		}
		fmt.Println("userPayments", userPayments)
		if len(userPayments) == 0 {
			break
		}
	}

	return
}

func (interactor *BillInteractor) Add(b domain.Bill) (bill domain.Bill, err error) {
	// if err = validate(b); err != nil {
	//   return
	// }
	bill, err = interactor.BillRepository.Store(b)
	return
}

func (interactor *BillInteractor) Update(b domain.Bill) (bill domain.Bill, err error) {
	bill, err = interactor.BillRepository.Update(b)
	return
}

func (interactor *BillInteractor) DeleteById(b domain.Bill) (err error) {
	err = interactor.BillRepository.DeleteById(b)
	return
}

func validate(bill domain.Bill) error {
	if len(bill.Name) == 0 {
		return errors.New("nameを入力してください")
	}
	return nil
}
