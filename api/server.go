package api

import (
	"log"
	"strings"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

	var phoneType phonebook.PhoneNumber_PhoneType
	switch strings.ToLower(req.PhoneNumberType) {
	case "mobile":
		phoneType = phonebook.PhoneNumber_MOBILE
	case "home":
		phoneType = phonebook.PhoneNumber_HOME
	case "work":
		phoneType = phonebook.PhoneNumber_WORK
	default:
		return nil, status.Errorf(codes.InvalidArgument, "invalid phone type: %v", req.PhoneNumberType)
	}
	phoneNumber := &phonebook.PhoneNumber{
		Number: req.PhoneNumber,
		Type:   phoneType,
	}

	contact := &phonebook.Contact{
		Name:  req.Name,
		Email: req.Email,
		PhoneNumber: []*phonebook.PhoneNumber{
			phoneNumber,
		},
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
