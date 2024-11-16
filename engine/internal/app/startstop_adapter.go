package app

import "context"

type funcWithErr func() error

type startStopAdpt struct {
	ctx      context.Context
	launcher funcWithErr
	finisher funcWithErr
}

func NewStartStopAdpt(ctx context.Context, launcher func() error, finisher func() error) *startStopAdpt {
	return &startStopAdpt{
		ctx:      ctx,
		launcher: launcher,
		finisher: finisher,
	}
}

func (ssa startStopAdpt) Start() error {
	if ssa.launcher != nil {
		if err := ssa.launcher(); err != nil {
			return err
		}
	}
	<-ssa.ctx.Done()
	return nil
}

func (ssa startStopAdpt) Stop() error {
	if ssa.finisher == nil {
		return nil
	}
	return ssa.finisher()
}
