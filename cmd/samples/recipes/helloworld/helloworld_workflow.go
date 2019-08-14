package main

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

/**
 * This is the hello world workflow sample.
 */

// ApplicationName is the task list for this sample
const ApplicationName = "helloWorldGroup"

// This is registration process where you register all your workflows
// and activity function handlers.
func init() {
	workflow.Register(Workflow)
	activity.Register(helloworldActivity)
}

// Workflow workflow decider
func Workflow(ctx workflow.Context, name string) error {
	ao := workflow.LocalActivityOptions{
		ScheduleToCloseTimeout: time.Hour * 10,
		//RetryPolicy: &cadence.RetryPolicy{
		//	InitialInterval:    time.Second,
		//	BackoffCoefficient: 2.0,
		//	MaximumInterval:    time.Minute,
		//	ExpirationInterval: time.Minute * 30,
		//	MaximumAttempts:    0,
		//},
	}
	ctx = workflow.WithLocalActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("helloworld workflow started")
	var helloworldResult string
	err := workflow.ExecuteLocalActivity(ctx, helloworldActivity, name).Get(ctx, &helloworldResult)
	if err != nil {
		logger.Error("Activity failed.", zap.Error(err))
		return err
	}

	logger.Info("Workflow completed.", zap.String("Result", helloworldResult))

	return nil
}

func helloworldActivity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("helloworld activity started")
	time.Sleep(time.Hour)
	return "", fmt.Errorf("error for retry")
}
