package application

import "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/domain"

type EventBus interface {
	Publish(event domain.Event) error
}
