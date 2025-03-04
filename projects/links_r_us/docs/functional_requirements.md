## Functional Requirements

### User story – link submission

As an end user,

I need to be able to submit new links to Links 'R' Us,

so as to update the link graph and make their contents searchable.

The acceptance criteria for this user story are as follows:
- A frontend or API endpoint is provided for facilitating the link submission journey for the end users.
- Submitted links have the following criteria:
    - Must be added to the graph.
    - Must be crawled by the system and added to their index.
- Already submitted links should be accepted by the backend but not inserted twice to the graph.

### User story – search

As an end user,

I need to be able to submit full-text search queries,

so as to to retrieve a list of relevant matching results from the content indexed by Links 'R' Us.

The acceptance criteria for this user story are as follows:
- A frontend or API endpoint is provided for the users to submit a full-text query.
- If the query matches multiple items, they are returned as a list that the end user can paginate through.
- Each entry in the result list must contain the following items: title or link description, the link to the content, and a timestamp indicating when the link was last crawled. If feasible, the link may also contain a relevance score expressed as a percentage.
- When the query does not match any item, an appropriate response should be returned to the end user.

### User story – crawl link graph

As the crawler backend system,

I need to be able to obtain a list of sanitized links from the link graph, 

so as to fetch and index their contents while at the same time expanding the link graph with newly discovered links.

The acceptance criteria for this user story are as follows:
- The crawler can query the link graph and receive a list of stale links that need to be crawled.
- Links received by the crawler are retrieved from the remote hosts unless the remote server provides an ETag or Last Modified header that the crawler has already seen before.
- Retrieved content is scanned for links and the link graph gets updated.
- Retrieved content is indexed and added to the search corpus.

### User story – calculate PageRank scores

As the PageRank calculator backend system,

I need to be able to access the link graph,

so as to calculate and persist the PageRank score for each link.

The acceptance criteria for this user story are as follows:
- The PageRank calculator can obtain an immutable snapshot of the entire link graph.
- A PageRank score is assigned to every link in the graph.
- The search corpus entries are annotated with the updated PageRank scores.

### User story – monitor Links 'R' Us health

As a member of the Links 'R' Us Site Reliability Engineering (SRE) team, 

I need to be able to monitor the health of all Links 'R' Us services,

so as to detect and address issues that cause degraded service performance.

The acceptance criteria for this user story are as follows:
- All Links 'R' Us services should periodically submit health- and performance- related metrics to a centralized metrics collection system.
- A monitoring dashboard is created for each service.
- A high-level monitoring dashboard tracks the overall system health.
- Metric-based alerts are defined and linked to a paging service. Each alert comes with its own playbook with a set of steps that need to be performed by a member of the SRE team that is on-call.
