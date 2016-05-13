package session

type State uint

type Session struct {
	ChatID int64
	User   string  // User name
	State  []State // What interface user is using
}

// int = User.ID
var sessions = map[int]*Session{}

func NewSession(userID int) *Session {
	s := &Session{}
	sessions[userID] = s
	return s
}

func GetSession(userID int) *Session {
	s, ok := sessions[userID]
	if !ok {
		return nil
	}
	return s
}
