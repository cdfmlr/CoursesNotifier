package models

type Relationship struct {
	Sid string
	Cid string
}

func NewRelationship(sid string, cid string) *Relationship {
	return &Relationship{Sid: sid, Cid: cid}
} 
