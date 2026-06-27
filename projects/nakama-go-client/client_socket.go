package nkclient

type Socket struct{}
type Event struct{}
type Notification struct{}
type MatchData struct{}
type MatchPresenceEvent struct{}

func (c *Client) CreateSocket(session *Session) (*Socket, error) {
	panic("Not implemented")
}

func (s *Socket) Connect() error
func (s *Socket) Disconnect() error
func (s *Socket) JoinMatch(matchId string) error
func (s *Socket) SendMatchState(matchId string, opCode int, data any) error
func (s *Socket) LeaveMatch(matchId string) error
func (s *Socket) Rpc(id string, httpKey string, payload any) error
func (s *Socket) OnError(fn func(evt Event))
func (s *Socket) OnNotification(fn func(evt Notification))
func (s *Socket) OnMatchdata(fn func(evt MatchData))
func (s *Socket) OnMatchPresence(fn func(evt MatchPresenceEvent))
