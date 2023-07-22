package services

import (
	"context"
	"gogin/model"
)

type IContactService interface {
	GetAllContacts(ctx context.Context) ([]*model.Contact, error)
}
type ContactService struct {
	contactModel *model.ContactModel
}

func NewContactService(contactModel *model.ContactModel) *ContactService {
	return &ContactService{
		contactModel: contactModel,
	}
}

func (s *ContactService) GetAllContacts(ctx context.Context) ([]*model.Contact, error) {
	return s.contactModel.FindAllContact(ctx)
}
