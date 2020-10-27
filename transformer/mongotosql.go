package transformer

import (
	"api-reflist-migration-service/model"
	"time"
)

func Transform(mreflists []model.ReferenceListMongo) []model.ReferenceList{
	var reflists []model.ReferenceList
	for _, mreflist := range mreflists {
		var reflist model.ReferenceList
		reflist.DateUpdated = time.Now()
		reflist.ID = mreflist.ID
		reflist.DateCreated = mreflist.DateCreated
		reflist.Description=mreflist.Description
		reflist.Status=mreflist.Status
		reflist.Code=mreflist.Code
		reflist.Type=mreflist.Type
		reflist.Name=mreflist.Name
		reflists = append(reflists, reflist)
	}
	return reflists
}
