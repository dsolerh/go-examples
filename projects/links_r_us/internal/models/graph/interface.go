package graph

import (
	"linkrus/internal/models/iterator"
	"time"

	"github.com/google/uuid"
)

type Graph interface {
	// links
	UpsertLink(link *Link) error
	FindLink(id uuid.UUID) (*Link, error)
	Links(fromID, toID uuid.UUID, retrieveBefore time.Time) (LinkIterator, error)
	// edges
	UpsertEdge(edge *Edge) error
	RemoveStaleEdges(fromID uuid.UUID, updatedBefore time.Time) error
	Edges(fromID, toID uuid.UUID, updatedBefore time.Time) (EdgeIterator, error)
}

type Link struct {
	Id          uuid.UUID
	Url         string
	RetrievedAt time.Time
}

func (l *Link) GetID() uuid.UUID {
	return l.Id
}

type Edge struct {
	Id        uuid.UUID
	Src, Dst  uuid.UUID
	UpdatedAt time.Time
}

func (e *Edge) GetID() uuid.UUID {
	return e.Id
}

type LinkIterator = iterator.Iterator[*Link]

type EdgeIterator = iterator.Iterator[*Edge]
