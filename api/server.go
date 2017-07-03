package api

import (
	"log"
	"sync"

	oldctx "golang.org/x/net/context"

	phonebook "github.com/iheanyi/grpc-phonebook"
)

type server struct {
	contactsByNameMu sync.Mutex
	contactsByName   map[string]*phonebook.Contact
}

func New() PhoneBookServer {
	return &server{}
}

func (svc *server) CreateContact(ctx oldctx.Context, req *CreateContactReq) (*CreateContactRes, error) {
	log.Printf("in the create contact method, with req: %f", req)

	contact := &phonebook.Contact{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumbers,
	}

	svc.contactsByNameMu.Lock()
	defer svc.contactsByNameMu.Unlock()
	svc.contactsByName[contact.Name] = contact
	res := &CreateContactRes{
		Contact: contact,
	}

	log.Printf("Created the following contact: %v", contact)
	return res, nil
}

func (svc *server) ListContacts(ctx oldctx.Context, req *ListContactsReq) (*ListContactsRes, error) {

	return nil, nil
}
