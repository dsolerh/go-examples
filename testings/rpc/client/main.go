package main

import (
	"errors"
	"log"
	"net/rpc"
	"sort"
	"sync"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:12098")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer client.Close()

	const n = 50_000
	stats := benchmark(client, n)
	log.Printf("requests=%d errors=%d (%.2f%%)\n", stats.total, stats.errors, stats.errorRate*100)
	log.Printf("avg=%s p50=%s p95=%s\n", stats.avg, stats.p50, stats.p95)
}

// benchmark fires n concurrent Echo calls and collects timing statistics.
func benchmark(client *rpc.Client, n int) stats {
	data := make(chan info, n)

	services := []string{"EchoDelay"}
	// services := []string{"EchoOne", "EchoTwo", "EchoThree"}
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				call(client, services[i%len(services)], data)
			}
		}()
	}

	// Close the channel once every call has reported, so collect can range to completion.
	go func() {
		wg.Wait()
		close(data)
	}()

	return collect(data)
}

type stats struct {
	total     int
	errors    int
	errorRate float64
	avg       time.Duration
	p50       time.Duration
	p95       time.Duration
}

// collect drains the channel and computes avg, p50, p95 of the delays plus the error rate.
func collect(data <-chan info) stats {
	var s stats
	delays := make([]time.Duration, 0)
	var sum time.Duration

	for in := range data {
		s.total++
		if in.err != nil {
			s.errors++
			continue
		}
		delays = append(delays, in.delay)
		sum += in.delay
	}

	if s.total > 0 {
		s.errorRate = float64(s.errors) / float64(s.total)
	}
	if len(delays) == 0 {
		return s
	}

	sort.Slice(delays, func(i, j int) bool { return delays[i] < delays[j] })
	s.avg = sum / time.Duration(len(delays))
	s.p50 = percentile(delays, 0.50)
	s.p95 = percentile(delays, 0.95)
	return s
}

// percentile returns the value at the given quantile of a pre-sorted slice
// using nearest-rank.
func percentile(sorted []time.Duration, q float64) time.Duration {
	if len(sorted) == 0 {
		return 0
	}
	rank := int(q * float64(len(sorted)))
	if rank >= len(sorted) {
		rank = len(sorted) - 1
	}
	return sorted[rank]
}

type (
	payload struct {
		Data string
	}
	info struct {
		delay time.Duration
		err   error
	}
)

func call(client *rpc.Client, methodName string, data chan<- info) {
	args := payload{Data: largePayload()}
	var reply payload

	var err error
	start := time.Now()
	call := client.Go("service."+methodName, args, &reply, make(chan *rpc.Call, 1))
	select {
	case <-call.Done:
		err = call.Error
	case <-time.After(500 * time.Millisecond):
		err = errors.New("Service timeout")
	}
	data <- info{delay: time.Since(start), err: err}
}

func largePayload() string {
	return `The standard Lorem Ipsum passage, used since 1966
"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

Section 1.10.32 of "de Finibus Bonorum et Malorum", written by Cicero in 45 BC
"Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?"

1914 translation by H. Rackham
"But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain was born and I will give you a complete account of the system, and expound the actual teachings of the great explorer of the truth, the master-builder of human happiness. No one rejects, dislikes, or avoids pleasure itself, because it is pleasure, but because those who do not know how to pursue pleasure rationally encounter consequences that are extremely painful. Nor again is there anyone who loves or pursues or desires to obtain pain of itself, because it is pain, but because occasionally circumstances occur in which toil and pain can procure him some great pleasure. To take a trivial example, which of us ever undertakes laborious physical exercise, except to obtain some advantage from it? But who has any right to find fault with a man who chooses to enjoy a pleasure that has no annoying consequences, or one who avoids a pain that produces no resultant pleasure?"

Section 1.10.33 of "de Finibus Bonorum et Malorum", written by Cicero in 45 BC
"At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum deleniti atque corrupti quos dolores et quas molestias excepturi sint occaecati cupiditate non provident, similique sunt in culpa qui officia deserunt mollitia animi, id est laborum et dolorum fuga. Et harum quidem rerum facilis est et expedita distinctio. Nam libero tempore, cum soluta nobis est eligendi optio cumque nihil impedit quo minus id quod maxime placeat facere possimus, omnis voluptas assumenda est, omnis dolor repellendus. Temporibus autem quibusdam et aut officiis debitis aut rerum necessitatibus saepe eveniet ut et voluptates repudiandae sint et molestiae non recusandae. Itaque earum rerum hic tenetur a sapiente delectus, ut aut reiciendis voluptatibus maiores alias consequatur aut perferendis doloribus asperiores repellat."

1914 translation by H. Rackham
"On the other hand, we denounce with righteous indignation and dislike men who are so beguiled and demoralized by the charms of pleasure of the moment, so blinded by desire, that they cannot foresee the pain and trouble that are bound to ensue; and equal blame belongs to those who fail in their duty through weakness of will, which is the same as saying through shrinking from toil and pain. These cases are perfectly simple and easy to distinguish. In a free hour, when our power of choice is untrammelled and when nothing prevents our being able to do what we like best, every pleasure is to be welcomed and every pain avoided. But in certain circumstances and owing to the claims of duty or the obligations of business it will frequently occur that pleasures have to be repudiated and annoyances accepted. The wise man therefore always holds in these matters to this principle of selection: he rejects pleasures to secure other greater pleasures, or else he endures pains to avoid worse pains."
`
}
