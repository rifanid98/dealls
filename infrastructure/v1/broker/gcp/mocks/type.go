package mocks

import (
	"cloud.google.com/go/iam"
	"cloud.google.com/go/pubsub"
)

type ClientMock struct {
	CreateSubscription       *Subscription
	CreateSubscriptionErr    error
	CreateTopic              *Topic
	CreateTopicErr           error
	CreateTopicWithConfig    *Topic
	CreateTopicWithConfigErr error
	DetachSubscription       pubsub.DetachSubscriptionResult
	DetachSubscriptionErr    error
	Subscription             *Subscription
	SubscriptionInProject    *Subscription
	Subscriptions            *pubsub.SubscriptionIterator
	Topic                    *Topic
	TopicInProject           *Topic
	Topics                   *pubsub.TopicIterator
}

type SubscriptionMock struct {
	String    string
	ID        string
	Delete    error
	Exists    bool
	ExistsErr error
	Config    pubsub.SubscriptionConfig
	ConfigErr error
	Update    pubsub.SubscriptionConfig
	UpdateErr error
	IAM       *iam.Handle
	Receive   error
}

type TopicMock struct {
	Config        pubsub.TopicConfig
	ConfigErr     error
	Update        pubsub.TopicConfig
	UpdateErr     error
	ID            string
	String        string
	Delete        error
	Exists        bool
	ExistsErr     error
	IAM           *iam.Handle
	Subscriptions *pubsub.SubscriptionIterator
	Publish       *pubsub.PublishResult
}
