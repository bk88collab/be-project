package response

import "movie/businesses/linktrailers"

type LinkTrailer struct {
	Id_trailer int    `json:"id_trailer"`
	Url        string `json:"url"`
}

func FromDomain(domain linktrailers.Domain) LinkTrailer {
	return LinkTrailer{
		Id_trailer: domain.Id_trailer,
		Url:        domain.Url,
	}
}
