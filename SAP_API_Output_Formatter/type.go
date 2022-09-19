package sap_api_output_formatter

type BusinessPartner struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	APISchema       string `json:"api_schema"`
	BusinessPartner string `json:"business_partner_code"`
	Deleted         bool   `json:"deleted"`
}

type General struct {
	BusinessPartner                string      `json:"BusinessPartner"`
	Customer                       string      `json:"Customer"`
	Supplier                       string      `json:"Supplier"`
	AcademicTitle                  string      `json:"AcademicTitle"`
	AuthorizationGroup             string      `json:"AuthorizationGroup"`
	BusinessPartnerCategory        string      `json:"BusinessPartnerCategory"`
	BusinessPartnerFullName        string      `json:"BusinessPartnerFullName"`
	BusinessPartnerGrouping        string      `json:"BusinessPartnerGrouping"`
	BusinessPartnerName            string      `json:"BusinessPartnerName"`
	CorrespondenceLanguage         string      `json:"CorrespondenceLanguage"`
	CreationDate                   string      `json:"CreationDate"`
	CreationTime                   string      `json:"CreationTime"`
	FirstName                      string      `json:"FirstName"`
	Industry                       string      `json:"Industry"`
	IsFemale                       bool        `json:"IsFemale"`
	IsMale                         bool        `json:"IsMale"`
	IsNaturalPerson                string      `json:"IsNaturalPerson"`
	IsSexUnknown                   bool        `json:"IsSexUnknown"`
	GenderCodeName                 string      `json:"GenderCodeName"`
	Language                       string      `json:"Language"`
	LastChangeDate                 string      `json:"LastChangeDate"`
	LastChangeTime                 string      `json:"LastChangeTime"`
	LastName                       string      `json:"LastName"`
	OrganizationBPName1            string      `json:"OrganizationBPName1"`
	OrganizationBPName2            string      `json:"OrganizationBPName2"`
	OrganizationBPName3            string      `json:"OrganizationBPName3"`
	OrganizationBPName4            string      `json:"OrganizationBPName4"`
	OrganizationFoundationDate     string      `json:"OrganizationFoundationDate"`
	OrganizationLiquidationDate    string      `json:"OrganizationLiquidationDate"`
	SearchTerm1                    string      `json:"SearchTerm1"`
	SearchTerm2                    string      `json:"SearchTerm2"`
	AdditionalLastName             string      `json:"AdditionalLastName"`
	BirthDate                      string      `json:"BirthDate"`
	BusinessPartnerBirthplaceName  string      `json:"BusinessPartnerBirthplaceName"`
	BusinessPartnerDeathDate       string      `json:"BusinessPartnerDeathDate"`
	BusinessPartnerIsBlocked       bool        `json:"BusinessPartnerIsBlocked"`
	BusinessPartnerType            string      `json:"BusinessPartnerType"`
	GroupBusinessPartnerName1      string      `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2      string      `json:"GroupBusinessPartnerName2"`
	IndependentAddressID           string      `json:"IndependentAddressID"`
	MiddleName                     string      `json:"MiddleName"`
	NameCountry                    string      `json:"NameCountry"`
	PersonFullName                 string      `json:"PersonFullName"`
	PersonNumber                   string      `json:"PersonNumber"`
	IsMarkedForArchiving           bool        `json:"IsMarkedForArchiving"`
	BusinessPartnerIDByExtSystem   string      `json:"BusinessPartnerIDByExtSystem"`
	TradingPartner                 string      `json:"TradingPartner"`
	ToRole                         string      `json:"to_BusinessPartnerRole"`
	ToAddress                      string      `json:"to_BusinessPartnerAddress"`
	ToBank                         string      `json:"to_BusinessPartnerBank"`
	ToContact                      string      `json:"to_BusinessPartnerContact"`
}

type Role struct {
	BusinessPartner     string `json:"BusinessPartner"`
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	ValidFrom           string `json:"ValidFrom"`
	ValidTo             string `json:"ValidTo"`
}

type Address struct {
	BusinessPartner   string `json:"BusinessPartner"`
	AddressID         string `json:"AddressID"`
	ValidityStartDate string `json:"ValidityStartDate"`
	ValidityEndDate   string `json:"ValidityEndDate"`
	Country           string `json:"Country"`
	Region            string `json:"Region"`
	StreetName        string `json:"StreetName"`
	CityName          string `json:"CityName"`
	PostalCode        string `json:"PostalCode"`
	Language          string `json:"Language"`
}

type Bank struct {
	BusinessPartner          string      `json:"BusinessPartner"`
	BankIdentification       string      `json:"BankIdentification"`
	BankCountryKey           string      `json:"BankCountryKey"`
	BankName                 string      `json:"BankName"`
	BankNumber               string      `json:"BankNumber"`
	SWIFTCode                string      `json:"SWIFTCode"`
	BankControlKey           string      `json:"BankControlKey"`
	BankAccountHolderName    string      `json:"BankAccountHolderName"`
	BankAccountName          string      `json:"BankAccountName"`
	ValidityStartDate        string      `json:"ValidityStartDate"`
	ValidityEndDate          string      `json:"ValidityEndDate"`
	IBAN                     string      `json:"IBAN"`
	IBANValidityStartDate    string      `json:"IBANValidityStartDate"`
	BankAccount              string      `json:"BankAccount"`
	BankAccountReferenceText string      `json:"BankAccountReferenceText"`
	CollectionAuthInd        bool        `json:"CollectionAuthInd"`
	CityName                 string      `json:"CityName"`
	AuthorizationGroup       string      `json:"AuthorizationGroup"`
}

type ToRole struct {
	BusinessPartner     string `json:"BusinessPartner"`
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	ValidFrom           string `json:"ValidFrom"`
	ValidTo             string `json:"ValidTo"`
}

type ToAddress struct {
	BusinessPartner   string `json:"BusinessPartner"`
	AddressID         string `json:"AddressID"`
	ValidityStartDate string `json:"ValidityStartDate"`
	ValidityEndDate   string `json:"ValidityEndDate"`
	Country           string `json:"Country"`
	Region            string `json:"Region"`
	StreetName        string `json:"StreetName"`
	CityName          string `json:"CityName"`
	PostalCode        string `json:"PostalCode"`
	Language          string `json:"Language"`
}

type ToBank struct {
	BusinessPartner          string      `json:"BusinessPartner"`
	BankIdentification       string      `json:"BankIdentification"`
	BankCountryKey           string      `json:"BankCountryKey"`
	BankName                 string      `json:"BankName"`
	BankNumber               string      `json:"BankNumber"`
	SWIFTCode                string      `json:"SWIFTCode"`
	BankControlKey           string      `json:"BankControlKey"`
	BankAccountHolderName    string      `json:"BankAccountHolderName"`
	BankAccountName          string      `json:"BankAccountName"`
	ValidityStartDate        string      `json:"ValidityStartDate"`
	ValidityEndDate          string      `json:"ValidityEndDate"`
	IBAN                     string      `json:"IBAN"`
	IBANValidityStartDate    string      `json:"IBANValidityStartDate"`
	BankAccount              string      `json:"BankAccount"`
	BankAccountReferenceText string      `json:"BankAccountReferenceText"`
	CollectionAuthInd        bool        `json:"CollectionAuthInd"`
	CityName                 string      `json:"CityName"`
	AuthorizationGroup       string      `json:"AuthorizationGroup"`
}

type ToContact struct {
	BusinessPartner               string `json:"BusinessPartner"`
	RelationshipNumber            string `json:"RelationshipNumber"`
	BusinessPartnerCompany        string `json:"BusinessPartnerCompany"`
	BusinessPartnerPerson         string `json:"BusinessPartnerPerson"`
	ValidityEndDate               string `json:"ValidityEndDate"`
	ValidityStartDate             string `json:"ValidityStartDate"`
	IsStandardRelationship        string `json:"IsStandardRelationship"`
	RelationshipCategory          string `json:"RelationshipCategory"`
}
