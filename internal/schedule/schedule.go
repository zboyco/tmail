package schedule

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"time"
	"tmail/ent"
	"tmail/ent/attachment"
	"tmail/ent/envelope"
)

type Scheduler struct {
	db *ent.Client
}

func New(db *ent.Client) *Scheduler {
	return &Scheduler{db: db}
}

func (s *Scheduler) Run() {
	go s.cleanUpExpired()
}

func (s *Scheduler) cleanUpExpired() {
	run(func() {
		expired := time.Now().Add(-time.Hour * 240)
		list, err := s.db.Attachment.Query().Where(attachment.HasOwnerWith(envelope.CreatedAtLT(expired))).All(context.TODO())
		if err != nil {
			log.Err(err).Msg("Attachment Query")
			return
		}
		for _, a := range list {
			_ = os.Remove(a.Filepath)
		}
		count, err := s.db.Attachment.Delete().Where(attachment.HasOwnerWith(envelope.CreatedAtLT(expired))).Exec(context.TODO())
		if err != nil {
			log.Err(err).Msg("Attachment Delete")
			return
		}
		if count > 0 {
			log.Info().Msgf("clean up attachment %d", count)
		}
		count, err = s.db.Envelope.Delete().Where(envelope.CreatedAtLT(expired)).Exec(context.TODO())
		if err != nil {
			log.Err(err).Msg("Envelope Delete")
			return
		}
		if count > 0 {
			log.Info().Msgf("clean up expired %d", count)
		}
	}, time.Hour)
}

func run(fn func(), dur time.Duration) {
	for {
		select {
		case <-time.Tick(dur):
			go func() {
				defer func() {
					if err := recover(); err != nil {
						log.Error().Msg(fmt.Sprint(err))
					}
				}()
				fn()
			}()
		}
	}
}
