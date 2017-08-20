package server

import (
	"log"
	"sync"

	oldctx "golang.org/x/net/context"

	api "github.com/iheanyi/grpc-phonebook/api"
)

type server struct {
	contactsByNameMu sync.Mutex
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
	}

	svc.contactsByNameMu.Lock()
	defer svc.contactsByNameMu.Unlock()
	svc.contactsByName[contact.Name] = contact
	svc.contacts = append(svc.contacts, contact)
	res := &api.CreateContactRes{
		Contact: contact,
	}

	log.Printf("Done creating contact %v", contact)
	return res, nil
}

func (svc *server) ListContacts(ctx oldctx.Context, req *api.ListContactsReq) (*api.ListContactsRes, error) {
	res := &api.ListContactsRes{
		Contacts: svc.contacts,
	}

	return res, nil
}
