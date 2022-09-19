package responses

type ToContact struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			BusinessPartner               string `json:"BusinessPartner"`
			RelationshipNumber            string `json:"RelationshipNumber"`
			BusinessPartnerCompany        string `json:"BusinessPartnerCompany"`
			BusinessPartnerPerson         string `json:"BusinessPartnerPerson"`
			ValidityEndDate               string `json:"ValidityEndDate"`
			ValidityStartDate             string `json:"ValidityStartDate"`
			IsStandardRelationship        string `json:"IsStandardRelationship"`
			RelationshipCategory          string `json:"RelationshipCategory"`
		} `json:"results"`
	} `json:"d"`
}
