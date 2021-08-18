package request

import (
	"movie/businesses/linktrailers"
)

type LinkTrailer struct {
	Id_trailer int    `json:"id_trailer"`
	Url        string `json:"url"`
}

func (request *LinkTrailer) ToDomain() *linktrailers.Domain {
	return &linktrailers.Domain{
		Id_trailer: request.Id_trailer,
		Url:        request.Url,
	}
}
