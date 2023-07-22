package api

import (
	"github.com/gin-gonic/gin"
	"gogin/services"
	"net/http"
)

type ContactAPI struct {
	contactService services.IContactService
}

func NewContactAPI(contactService services.IContactService) *ContactAPI {
	return &ContactAPI{
		contactService: contactService,
	}
}

func (api *ContactAPI) GetAllContactsHandler(c *gin.Context) {
	contacts, err := api.contactService.GetAllContacts(c)
	if api.contactService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "contactService is not initialized"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contacts)
}

// func (api *ContactAPI) GetAllContactsHandler(w http.ResponseWriter, r *http.Request) {
// 	contacts, err := api.contactService.GetAllContacts()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	jsonData, err := json.Marshal(contacts)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(jsonData)
// }
