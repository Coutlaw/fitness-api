package models

import (
	u "fitness-api/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"` //The user that this contact belongs to
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (contact *Contact) Validate() (string, bool) {

	if contact.Name == "" {
		return "Contact name should be on the payload", false
	}

	if contact.Phone == "" {
		return "Phone number should be on the payload", false
	}

	if contact.UserId <= 0 {
		return "User is not recognized", false
	}

	//All the required parameters are present
	return "success", true
}

func (contact *Contact) Create() map[string]interface{} {

	if resp, ok := contact.Validate(); !ok {
		return u.Message(false, resp)
	}

	GetDB().Create(contact)

	resp := u.Message(true, "success")
	resp["contact"] = contact
	return resp
}

func GetContact(contactId uint, userId uint) *Contact {

	contact := &Contact{}
	err := GetDB().Table("contacts").Where("user_id = ?", userId).Where("id = ?", contactId).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func DeleteContact(contactId uint, userId uint) *Contact {

	contact := &Contact{}
	err := GetDB().Table("contacts").Where("user_id = ?", userId).Where("id = ?", contactId).Delete(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func DeleteContacts(userId uint) *Contact {

	contact := &Contact{}
	err := GetDB().Table("contacts").Where("user_id = ?", userId).Delete(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(userId uint) []*Contact {

	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", userId).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}
