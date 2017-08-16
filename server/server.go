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
	return &server{}
}

func (svc *server) CreateContact(ctx oldctx.Context, req *api.CreateContactReq) (*api.CreateContactRes, error) {
	log.Printf("in the create contact method, with req: %f", req)

	contact := &api.Contact{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumbers,
	}

	svc.contactsByNameMu.Lock()
	defer svc.contactsByNameMu.Unlock()
	svc.contactsByName[contact.Name] = contact
	svc.contacts = append(svc.contacts, contact)
	res := &api.CreateContactRes{
		Contact: contact,
	}

	log.Printf("Created the following contact: %v", contact)
	return res, nil
}

func (svc *server) ListContacts(ctx oldctx.Context, req *api.ListContactsReq) (*api.ListContactsRes, error) {
	return nil, nil
}
