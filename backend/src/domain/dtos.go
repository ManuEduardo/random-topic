package domain

import "encoding/xml"

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	SOAPNS  string   `xml:"xmlns:soap,attr"`
	Body    SOAPBody `xml:"Body"`
}

type SOAPBody struct {
	GetUserRequest        *GetUserRequest        `xml:",omitempty"`
	GetUserResponse       *GetUserResponse       `xml:",omitempty"`
	CreateUser            *CreateUser            `xml:",omitempty"`
	CreateUserResponse    *CreateUserResponse    `xml:",omitempty"`
	CreateCardRequest     *CreateCardRequest     `xml:",omitempty"`
	CreateCardResponse    *CreateCardResponse    `xml:",omitempty"`
	GetRandomCard         *GetRandomCard         `xml:",omitempty"`
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

type CreateUser struct {
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

type GetRandomCard struct {
	XMLName xml.Name `xml:"GetRandomCard"`
	UserID  int64    `xml:"UserID"`
}

type GetRandomCardResponse struct {
	XMLName xml.Name `xml:"GetRandomCardResponse"`
	Card    Card     `xml:"Card"`
}
