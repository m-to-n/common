package sessions

import "context"

const SESSION_ACTOR_TYPE = "M2NSessionActor"

type SessionActorClientStub struct {
	ActorId     string
	SendMessage func(context.Context, string) (string, error)
}

func NewSessionActorClientStub(actorId string) *SessionActorClientStub {
	return &SessionActorClientStub{
		ActorId: actorId,
	}
}

func (a *SessionActorClientStub) Type() string {
	return SESSION_ACTOR_TYPE
}

func (a *SessionActorClientStub) ID() string {
	return a.ActorId
}
