package response

// PdvPostResponse representa a resposta de um post de PDV
type PdvPostResponse struct {
	ID                 string  `json:"id"`
	RepresentativeUUID string  `json:"representative_uuid"`
	PdvName            string  `json:"pdv_name"`
	PostDate           string  `json:"post_date"`
	Month              string  `json:"month"`
	Platform           string  `json:"platform"`
	Link               *string `json:"link"`
	ProofUrl           *string `json:"proof_url"`
	Status             string  `json:"status"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
}

// RecurrentPdvResponse representa a resposta de um PDV recorrente
type RecurrentPdvResponse struct {
	ID                 string  `json:"id"`
	RepresentativeUUID string  `json:"representative_uuid"`
	Name               string  `json:"name"`
	City               *string `json:"city"`
	Followers          *int    `json:"followers"`
	InstagramProfile   *string `json:"instagram_profile"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
}

// PdvPostListData representa os dados da lista de posts de PDV
type PdvPostListData struct {
	Posts []PdvPostResponse `json:"posts"`
	Meta  MetaResponse      `json:"meta"`
}

// RecurrentPdvListData representa os dados da lista de PDVs recorrentes
type RecurrentPdvListData struct {
	Pdvs []RecurrentPdvResponse `json:"pdvs"`
	Meta MetaResponse           `json:"meta"`
}
