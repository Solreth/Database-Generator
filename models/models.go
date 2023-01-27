package models

import (
	"github.com/brianvoe/gofakeit/v6"
)

type User struct {
	ID            int    `json:"ID" gorm:"primaryKey" `
	Name          string `json:"name" fake:"{name}"`
	Email         string `json:"email" fake:"{email}"`
	Gender        string `json:"gender" fake:"{gender}"`
	SSN           int    `json:"ssn" fake:"{ssn}"`
	City          string `json:"city" fake:"{city}"`
	Street        string `json:"street" fake:"{street}"`
	FavoriteColor string `json:"color" fake:"{color}"`
	CreditCard    string `json:"credit_card" gorm:"type:text" fake:"{creditcard}"`
	Image         string `json:"image_url" fake:"{imageurl}"`
	Hobby         string `json:"hobby" fake:"{hobby}"`
	Occupation    string `json:"occupation" fake:"{job}"`
	Phone         string `json:"phone" fake:"{phoneformatted}"` //{phone for just the integers, may add toggle to format / not in future}
}

type FakeUser struct {
	Name   string `json:"name" fake:"{name}"`
	Email  string `json:"email" fake:"{email}"`
	Gender string `json:"gender" fake:"{gender}"`
}

var Users = []User{
	{Name: gofakeit.Name(), Email: gofakeit.Email(), Gender: gofakeit.Gender()},
	{Name: gofakeit.Name(), Email: gofakeit.Email(), Gender: gofakeit.Gender()},
	{Name: gofakeit.Name(), Email: gofakeit.Email(), Gender: gofakeit.Gender()},
	{Name: gofakeit.Name(), Email: gofakeit.Email(), Gender: gofakeit.Gender()},
	{Name: gofakeit.Name(), Email: gofakeit.Email(), Gender: gofakeit.Gender()},
}

var FakeUsers = []FakeUser{
	{Name: gofakeit.Name(), Email: gofakeit.Email(), Gender: gofakeit.Gender()},
	{Name: gofakeit.Name(), Email: gofakeit.Email(), Gender: gofakeit.Gender()},
	{Name: gofakeit.Name(), Email: gofakeit.Email(), Gender: gofakeit.Gender()},
	{Name: gofakeit.Name(), Email: gofakeit.Email(), Gender: gofakeit.Gender()},
	{Name: gofakeit.Name(), Email: gofakeit.Email(), Gender: gofakeit.Gender()},
}

// type Foo struct {
// 	Str           string
// 	Int           int
// 	Pointer       *int
// 	Name          string         `fake:"{firstname}"`  // Any available function all lowercase
// 	Sentence      string         `fake:"{sentence:3}"` // Can call with parameters
// 	RandStr       string         `fake:"{randomstring:[hello,world]}"`
// 	Number        string         `fake:"{number:1,10}"`       // Comma separated for multiple values
// 	Regex         string         `fake:"{regex:[abcdef]{5}}"` // Generate string from regex
// 	Map           map[string]int `fakesize:"2"`
// 	Array         []string       `fakesize:"2"`
// 	Bar           Bar
// 	Skip          *string   `fake:"skip"` // Set to "skip" to not generate data for
// 	Created       time.Time // Can take in a fake tag as well as a format tag
// 	CreatedFormat time.Time `fake:"{year}-{month}-{day}" format:"2006-01-02"`
// }

// type Bar struct {
// 	Name   string
// 	Number int
// 	Float  float32
// }

// CreditCard int      `json:"credit card" fake:"{creditcardinfo}"`
// SSN        int      `json:"ssn" fake:"{ssn}" `
// Friends    []string `json:"friends" fakesize:"3" fake:"{friendname}"`
