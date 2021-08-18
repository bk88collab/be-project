package linktrailers

import (
	"movie/businesses/linktrailers"
	"time"
)

type LinkUrl struct {
	Id_trailer int
	Url        string
	Created_At time.Time
	Updated_At time.Time
}

func (record *LinkUrl) toDomain() linktrailers.Domain {
	return linktrailers.Domain{
		Id_trailer: record.Id_trailer,
		Url:        record.Url,
		Created_At: record.Created_At,
		Updated_At: record.Updated_At,
	}
}

func fromDomainToRecord(linkDomain *linktrailers.Domain) *LinkUrl {
	return &LinkUrl{
		Id_trailer: linkDomain.Id_trailer,
		Url:        linkDomain.Url,
		Created_At: linkDomain.Created_At,
		Updated_At: linkDomain.Updated_At,
	}
}
