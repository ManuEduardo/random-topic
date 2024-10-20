package domain

import "encoding/xml"

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	SOAPNS  string   `xml:"xmlns:soap,attr"`
	Body    SOAPBody `xml:"soap:Body"`
}

type SOAPBody struct {
	GetUserRequest        *GetUserRequest        `xml:",omitempty"`
	GetUserResponse       *GetUserResponse       `xml:",omitempty"`
	CreateUserRequest     *CreateUserRequest     `xml:",omitempty"`
	CreateUserResponse    *CreateUserResponse    `xml:",omitempty"`
	CreateCardRequest     *CreateCardRequest     `xml:",omitempty"`
	CreateCardResponse    *CreateCardResponse    `xml:",omitempty"`
	GetRandomCardRequest  *GetRandomCardRequest  `xml:",omitempty"` // Asegúrate de que este campo esté presente
	GetRandomCardResponse *GetRandomCardResponse `xml:",omitempty"`
}

type GetUserRequest struct {
	XMLName xml.Name `xml:"GetUser"`
	UserID  int64    `xml:"UserID"`
}

type GetUserResponse struct {
	XMLName xml.Name `xml:"GetUserResponse"`
	User    User     `xml:"User"`
}

type CreateUserRequest struct {
	XMLName xml.Name `xml:"CreateUser"`
	User    User     `xml:"User"`
}

type CreateUserResponse struct {
	XMLName xml.Name `xml:"CreateUserResponse"`
	Success bool     `xml:"Success"`
}

type CreateCardRequest struct {
	XMLName xml.Name `xml:"CreateCard"`
	Card    Card     `xml:"Card"`
}

type CreateCardResponse struct {
	XMLName xml.Name `xml:"CreateCardResponse"`
	Success bool     `xml:"Success"`
}

type GetRandomCardRequest struct {
	XMLName xml.Name `xml:"GetRandomCard"` // Asegúrate de que coincida exactamente
	UserID  int64    `xml:"UserID"`
}

type GetRandomCardResponse struct {
	XMLName xml.Name `xml:"GetRandomCardResponse"`
	Card    Card     `xml:"Card"`
}
