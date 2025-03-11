package membership

import "coworking/internal/core/events"

type MembershipEventPublisher struct {
	publishedEvents []events.DomainEvent[MembershipEventTypes, any]
}

func NewMembershipEventPublisher() *MembershipEventPublisher {
	return &MembershipEventPublisher{
		publishedEvents: []events.DomainEvent[MembershipEventTypes, any]{},
	}
}

func (p *MembershipEventPublisher) Publish(e events.DomainEvent[MembershipEventTypes, any]) error {
	p.publishedEvents = append(p.publishedEvents, e)
	return nil
}
