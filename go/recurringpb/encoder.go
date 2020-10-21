package recurringpb

import (
	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MgoExpire struct {
	Month string `bson:"month" json:"month"`
	Year  string `bson:"year" json:"year"`
}

type MgoMultiLang struct {
	Lang  string `bson:"lang"`
	Value string `bson:"value"`
}

type MgoSavedCard struct {
	Id          primitive.ObjectID `bson:"_id"`
	Token       string             `bson:"token"`
	ProjectId   primitive.ObjectID `bson:"project_id"`
	MerchantId  primitive.ObjectID `bson:"merchant_id"`
	MaskedPan   string             `bson:"masked_pan"`
	CardHolder  string             `bson:"card_holder"`
	Expire      *MgoExpire         `bson:"expire"`
	RecurringId string             `bson:"recurring_id"`
	IsActive    bool               `bson:"is_active"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func (s *SavedCard) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(s.Id)

	if err != nil {
		return nil, err
	}

	projectOid, err := primitive.ObjectIDFromHex(s.ProjectId)

	if err != nil {
		return nil, err
	}

	merchantOid, err := primitive.ObjectIDFromHex(s.MerchantId)

	if err != nil {
		return nil, err
	}

	st := &MgoSavedCard{
		Id:         oid,
		Token:      s.Token,
		ProjectId:  projectOid,
		MerchantId: merchantOid,
		MaskedPan:  s.MaskedPan,
		CardHolder: s.CardHolder,
		Expire: &MgoExpire{
			Month: s.Expire.Month,
			Year:  s.Expire.Year,
		},
		RecurringId: s.RecurringId,
		IsActive:    s.IsActive,
	}

	t, err := ptypes.Timestamp(s.CreatedAt)

	if err != nil {
		return nil, err
	}

	st.CreatedAt = t

	t, err = ptypes.Timestamp(s.UpdatedAt)

	if err != nil {
		return nil, err
	}

	st.UpdatedAt = t

	return bson.Marshal(st)
}

func (s *SavedCard) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoSavedCard)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	s.Id = decoded.Id.Hex()
	s.Token = decoded.Token
	s.ProjectId = decoded.ProjectId.Hex()
	s.MerchantId = decoded.MerchantId.Hex()
	s.MaskedPan = decoded.MaskedPan
	s.CardHolder = decoded.CardHolder
	s.Expire = &CardExpire{
		Month: decoded.Expire.Month,
		Year:  decoded.Expire.Year,
	}
	s.RecurringId = decoded.RecurringId
	s.IsActive = decoded.IsActive
	s.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)

	if err != nil {
		return err
	}

	s.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

type MgoCustomerInfo struct {
	ExternalId string `bson:"external_id"`
	Email      string `bson:"email"`
	Phone      string `bson:"phone"`
}

type MgoSubscription struct {
	Id                    primitive.ObjectID `bson:"_id"`
	OrderId               primitive.ObjectID `bson:"order_id"`
	CustomerId            primitive.ObjectID `bson:"customer_id"`
	CustomerUuid          string             `bson:"customer_uuid"`
	Period                string             `bson:"period"`
	MerchantId            primitive.ObjectID `bson:"merchant_id"`
	ProjectId             primitive.ObjectID `bson:"project_id"`
	Amount                float64            `bson:"amount"`
	Currency              string             `bson:"currency"`
	ItemType              string             `bson:"item_type"`
	ItemList              []string           `bson:"item_list"`
	IsActive              bool               `bson:"is_active"`
	MaskedPan             string             `bson:"masked_pan"`
	CardpaySubscriptionId string             `bson:"cardpay_subscription_id"`
	CustomerInfo          *MgoCustomerInfo   `bson:"customer_info"`
	ExpireAt              time.Time          `bson:"expire_at"`
	LastPaymentAt         time.Time          `bson:"last_payment_at"`
	CreatedAt             time.Time          `bson:"created_at"`
	UpdatedAt             time.Time          `bson:"updated_at"`
	ProjectName           []*MgoMultiLang    `bson:"project_name"`
}

func (s *Subscription) MarshalBSON() ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(s.Id)

	if err != nil {
		return nil, err
	}

	orderOid, err := primitive.ObjectIDFromHex(s.OrderId)

	if err != nil {
		return nil, err
	}

	customerOid, err := primitive.ObjectIDFromHex(s.CustomerId)

	if err != nil {
		return nil, err
	}

	projectOid, err := primitive.ObjectIDFromHex(s.ProjectId)

	if err != nil {
		return nil, err
	}

	merchantOid, err := primitive.ObjectIDFromHex(s.MerchantId)

	if err != nil {
		return nil, err
	}

	st := &MgoSubscription{
		Id:                    oid,
		OrderId:               orderOid,
		CustomerId:            customerOid,
		CustomerUuid:          s.CustomerUuid,
		Period:                s.Period,
		MerchantId:            merchantOid,
		ProjectId:             projectOid,
		Amount:                s.Amount,
		Currency:              s.Currency,
		ItemType:              s.ItemType,
		ItemList:              s.ItemList,
		IsActive:              s.IsActive,
		MaskedPan:             s.MaskedPan,
		CardpaySubscriptionId: s.CardpaySubscriptionId,
	}

	if s.ProjectName != nil {
		for k, v := range s.ProjectName {
			st.ProjectName = append(st.ProjectName, &MgoMultiLang{Lang: k, Value: v})
		}
	}

	if s.CustomerInfo != nil {
		st.CustomerInfo = &MgoCustomerInfo{
			ExternalId: s.CustomerInfo.ExternalId,
			Email:      s.CustomerInfo.Email,
			Phone:      s.CustomerInfo.Phone,
		}
	}

	if s.LastPaymentAt != nil {
		t, err := ptypes.Timestamp(s.LastPaymentAt)

		if err != nil {
			return nil, err
		}

		st.LastPaymentAt = t
	}

	t, err := ptypes.Timestamp(s.ExpireAt)

	if err != nil {
		return nil, err
	}

	st.ExpireAt = t

	t, err = ptypes.Timestamp(s.CreatedAt)

	if err != nil {
		return nil, err
	}

	st.CreatedAt = t

	t, err = ptypes.Timestamp(s.UpdatedAt)

	if err != nil {
		return nil, err
	}

	st.UpdatedAt = t

	return bson.Marshal(st)
}

func (s *Subscription) UnmarshalBSON(raw []byte) error {
	decoded := new(MgoSubscription)
	err := bson.Unmarshal(raw, decoded)

	if err != nil {
		return err
	}

	s.Id = decoded.Id.Hex()
	s.OrderId = decoded.OrderId.Hex()
	s.CustomerId = decoded.CustomerId.Hex()
	s.CustomerUuid = decoded.CustomerUuid
	s.Period = decoded.Period
	s.MerchantId = decoded.MerchantId.Hex()
	s.ProjectId = decoded.ProjectId.Hex()
	s.Amount = decoded.Amount
	s.Currency = decoded.Currency
	s.ItemType = decoded.ItemType
	s.ItemList = decoded.ItemList
	s.MaskedPan = decoded.MaskedPan
	s.CardpaySubscriptionId = decoded.CardpaySubscriptionId
	s.IsActive = decoded.IsActive

	if decoded.ProjectName != nil {
		nameLen := len(decoded.ProjectName)
		if nameLen > 0 {
			s.ProjectName = make(map[string]string, nameLen)

			for _, v := range decoded.ProjectName {
				s.ProjectName[v.Lang] = v.Value
			}
		}
	}

	if decoded.CustomerInfo != nil {
		s.CustomerInfo = &CustomerInfo{
			ExternalId: decoded.CustomerInfo.ExternalId,
			Email:      decoded.CustomerInfo.Email,
			Phone:      decoded.CustomerInfo.Phone,
		}
	}

	if !decoded.LastPaymentAt.IsZero() {
		s.LastPaymentAt, err = ptypes.TimestampProto(decoded.LastPaymentAt)

		if err != nil {
			return err
		}
	}

	s.ExpireAt, err = ptypes.TimestampProto(decoded.ExpireAt)

	if err != nil {
		return err
	}

	s.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)

	if err != nil {
		return err
	}

	s.UpdatedAt, err = ptypes.TimestampProto(decoded.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}
