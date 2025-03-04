package inmemory

import (
	"fmt"
	"linkrus/internal/models/graph"
	"sync"
	"time"

	"github.com/google/uuid"
)

var _ graph.Graph = (*InMemoryGraph)(nil)

type edgeList []uuid.UUID

type InMemoryGraph struct {
	mu           sync.RWMutex
	links        map[uuid.UUID]*graph.Link
	edges        map[uuid.UUID]*graph.Edge
	linkURLIndex map[string]*graph.Link
	linkEdgeMap  map[uuid.UUID]edgeList
}

// Edges implements graph.Graph.
func (i *InMemoryGraph) Edges(fromID uuid.UUID, toID uuid.UUID, updatedBefore time.Time) (graph.EdgeIterator, error) {
	panic("unimplemented")
}

// FindLink implements graph.Graph.
func (i *InMemoryGraph) FindLink(id uuid.UUID) (*graph.Link, error) {
	panic("unimplemented")
}

// Links implements graph.Graph.
func (i *InMemoryGraph) Links(fromID uuid.UUID, toID uuid.UUID, retrieveBefore time.Time) (graph.LinkIterator, error) {
	panic("unimplemented")
}

// RemoveStaleEdges implements graph.Graph.
func (i *InMemoryGraph) RemoveStaleEdges(fromID uuid.UUID, updatedBefore time.Time) error {
	panic("unimplemented")
}

// UpsertEdge implements graph.Graph.
func (i *InMemoryGraph) UpsertEdge(edge *graph.Edge) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	_, srcExists := i.links[edge.Src]
	_, dstExists := i.links[edge.Dst]
	if !srcExists || !dstExists {
		return fmt.Errorf("upsert edge: %w", graph.ErrUnknownEdgeLinks)
	}

	// Scan edge list from source
	for _, edgeID := range i.linkEdgeMap[edge.Src] {
		existingEdge := i.edges[edgeID]
		if existingEdge.Src == edge.Src && existingEdge.Dst == edge.Dst {
			existingEdge.UpdatedAt = time.Now()
			*edge = *existingEdge
			return nil
		}
	}

	// Insert new edge
	for {
		edge.Id = uuid.New()
		if i.edges[edge.Id] == nil {
			break
		}
	}

	return nil
}

// UpsertLink creates a new link or updates an existing link.
func (i *InMemoryGraph) UpsertLink(link *graph.Link) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	// Check if a link with the same URL already exists. If so, convert
	// this into an update and point the link ID to the existing link.
	if existing := i.linkURLIndex[link.Url]; existing != nil {
		link.Id = existing.Id
		// update the RetrievedAt of the existing record if the new one is
		// newer
		if link.RetrievedAt.After(existing.RetrievedAt) {
			existing.RetrievedAt = link.RetrievedAt
		}
		return nil
	}

	// Assign new ID and insert link
	for {
		link.Id = uuid.New()
		// make sure there's no link with the id already
		if i.links[link.Id] == nil {
			break
		}
	}

	// make a copy of the link to store in memory
	// to prevent the pointer being updated outside
	// the scope of this function
	lCopy := new(graph.Link)
	*lCopy = *link
	// update the url index
	i.linkURLIndex[lCopy.Url] = lCopy
	i.links[lCopy.Id] = lCopy

	return nil
}
