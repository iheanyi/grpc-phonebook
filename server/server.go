package server

import (
	"log"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	oldctx "golang.org/x/net/context"

	api "github.com/iheanyi/grpc-phonebook/api"
)

type server struct {
	contactsByNameMu sync.RWMutex
	contactsByName   map[string]*api.Contact

	contacts []*api.Contact
}

func New() api.PhoneBookServer {
	return &server{
		contactsByName: make(map[string]*api.Contact),
		contacts: []*api.Contact{
			{
				Name:  "Iheanyi Ekechukwu",
				Email: "me@iheanyi.com",
				PhoneNumbers: []*api.PhoneNumber{
					{
						Number: "123-456-7890",
						Type:   api.PhoneNumber_HOME,
					},
				},
				Home: &api.PhoneNumber{
					Number: "843-340-0830",
					Type:   api.PhoneNumber_HOME,
				},
			},
		},
	}
}

func (svc *server) CreateContact(ctx oldctx.Context, req *api.CreateContactReq) (*api.CreateContactRes, error) {
	log.Printf("Creating contact %v", req)
	contact := &api.Contact{
		Name:         req.Name,
		Email:        req.Email,
		PhoneNumbers: req.PhoneNumbers,
		Home:         req.Home,
		Mobile:       req.Mobile,
		Work:         req.Work,
	}

	svc.contactsByNameMu.Lock()
	defer svc.contactsByNameMu.Unlock()

	if existingContact, ok := svc.contactsByName[contact.Name]; !ok {
		svc.contacts = append(svc.contacts, contact)
	} else if existingContact == contact {
		return nil, api.ErrorDuplicateContact
	}

	svc.contactsByName[contact.Name] = contact
	res := &api.CreateContactRes{
		Contact: contact,
	}

	log.Printf("Done creating contact %v", contact)
	return res, nil
}

func (svc *server) ListContacts(ctx oldctx.Context, req *api.ListContactsReq) (*api.ListContactsRes, error) {
	var contacts []*api.Contact
	// We're going to map the keys to an array.
	for _, contact := range svc.contactsByName {
		contacts = append(contacts, contact)
	}

	res := &api.ListContactsRes{
		Contacts: contacts,
	}

	return res, nil
}

func (svc *server) DeleteContact(ctx oldctx.Context, req *api.DeleteContactReq) (*api.DeleteContactRes, error) {
	svc.contactsByNameMu.Lock()
	defer svc.contactsByNameMu.Unlock()

	contact, ok := svc.contactsByName[req.Name]
	if !ok {
		return nil, api.ErrorNotFound
	}

	delete(svc.contactsByName, req.Name)
	res := &api.DeleteContactRes{
		Contact: contact,
	}

	return res, nil
}

func (svc *server) ShowContact(ctx oldctx.Context, req *api.ShowContactReq) (*api.ShowContactRes, error) {
	svc.contactsByNameMu.RLock()
	defer svc.contactsByNameMu.RUnlock()

	contact, ok := svc.contactsByName[req.Name]
	if !ok {
		return nil, api.ErrorNotFound
	}

	res := &api.ShowContactRes{
		Contact: contact,
	}

	return res, nil
}

func (svc *server) UpdateContact(ctx oldctx.Context, req *api.UpdateContactReq) (*api.UpdateContactRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}
