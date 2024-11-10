package app

import "context"

type funcWithErr func() error

type startStopAdpt struct {
	ctx      context.Context
	launcher funcWithErr
	lander   funcWithErr
}

func NewStartStopAdpt(ctx context.Context, launcher func() error, lander func() error) *startStopAdpt {
	return &startStopAdpt{
		ctx:      ctx,
		launcher: launcher,
		lander:   lander,
	}
}

func (ssa startStopAdpt) Start() error {
	if err := ssa.launcher(); err != nil {
		return err
	}
	<-ssa.ctx.Done()
	return nil
}

func (ssa startStopAdpt) Stop() error {
	return ssa.lander()
}
