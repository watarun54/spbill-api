package usecase

import (
	"errors"
	"fmt"
	"github.com/watarun54/spbill-api/server/domain"
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
		BillRepository       IBillRepository
		RoomRepository       IRoomRepository
		RoomMemberRepository IRoomMemberRepository
	}
)

type RoomMemberPayment struct {
	RoomMemberId int
	Amount       int
}

func (interactor *BillInteractor) ConvertBillFormToBill(billForm domain.BillForm) (bill domain.Bill, err error) {
	payees, err := interactor.RoomMemberRepository.FindByIds(billForm.PayeeIds)
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

func (interactor *BillInteractor) UserPayments(b domain.Bill) (memberPaymentsRes []domain.RoomMemberPaymentRes, err error) {
	room, err := interactor.RoomRepository.FindOne(domain.Room{
		ID: b.RoomID,
	})
	if err != nil {
		return
	}
	roomMembers := room.RoomMembers
	bills, err := interactor.BillRepository.FindAll(b)
	if err != nil {
		return
	}

	// memberId と 支払い差分 のペア
	response := map[int]int{}
	for _, member := range roomMembers {
		response[member.ID] = 0
	}
	for _, bill := range bills {
		response[bill.Payer.ID] += bill.Amount
		for _, payee := range bill.Payees {
			response[payee.ID] -= bill.Amount / (len(bill.Payees) + 1)
		}
	}

	memberIdToMembers := map[int]domain.RoomMember{}
	for _, member := range roomMembers {
		memberIdToMembers[member.ID] = member
	}

	memberPayments := []RoomMemberPayment{}
	for memberId, amount := range response {
		if amount != 0 {
			memberPayments = append(memberPayments, RoomMemberPayment{RoomMemberId: memberId, Amount: amount})
		}
	}

	var first RoomMemberPayment
	var last RoomMemberPayment
	var diff int

	for {
		sort.SliceStable(memberPayments, func(i, j int) bool { return memberPayments[i].Amount < memberPayments[j].Amount })

		fmt.Println("memberPayments", memberPayments)

		if len(memberPayments) <= 1 {
			break
		}

		first = memberPayments[0]
		memberPayments = memberPayments[1:]

		last = memberPayments[len(memberPayments)-1]
		memberPayments = memberPayments[:len(memberPayments)-1]

		if first.Amount > 0 && last.Amount > 0 {
			break
		}

		diff = first.Amount + last.Amount

		if diff > 0 {
			fmt.Printf("from: %v to: %v amount: %v\n", first.RoomMemberId, last.RoomMemberId, -(first.Amount))

			memberPaymentsRes = append(memberPaymentsRes, domain.RoomMemberPaymentRes{
				FromMember: memberIdToMembers[first.RoomMemberId],
				ToMember:   memberIdToMembers[last.RoomMemberId],
				Amount:     -(first.Amount),
			})
			last.Amount = diff
			memberPayments = append(memberPayments, last)
		} else if diff < 0 {
			fmt.Printf("from: %v to: %v amount: %v\n", last.RoomMemberId, first.RoomMemberId, last.Amount)

			memberPaymentsRes = append(memberPaymentsRes, domain.RoomMemberPaymentRes{
				FromMember: memberIdToMembers[last.RoomMemberId],
				ToMember:   memberIdToMembers[first.RoomMemberId],
				Amount:     last.Amount,
			})
			first.Amount = diff
			memberPayments = append(memberPayments, first)
		} else {
			fmt.Printf("from: %v to: %v amount: %v\n", last.RoomMemberId, first.RoomMemberId, last.Amount)

			memberPaymentsRes = append(memberPaymentsRes, domain.RoomMemberPaymentRes{
				FromMember: memberIdToMembers[last.RoomMemberId],
				ToMember:   memberIdToMembers[first.RoomMemberId],
				Amount:     last.Amount,
			})
		}

		fmt.Println("memberPayments", memberPayments)
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
