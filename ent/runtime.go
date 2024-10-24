// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/mayocream/twitter/ent/directmessage"
	"github.com/mayocream/twitter/ent/notification"
	"github.com/mayocream/twitter/ent/schema"
	"github.com/mayocream/twitter/ent/tweet"
	"github.com/mayocream/twitter/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	directmessageFields := schema.DirectMessage{}.Fields()
	_ = directmessageFields
	// directmessageDescSentAt is the schema descriptor for sent_at field.
	directmessageDescSentAt := directmessageFields[1].Descriptor()
	// directmessage.DefaultSentAt holds the default value on creation for the sent_at field.
	directmessage.DefaultSentAt = directmessageDescSentAt.Default.(func() time.Time)
	notificationFields := schema.Notification{}.Fields()
	_ = notificationFields
	// notificationDescCreatedAt is the schema descriptor for created_at field.
	notificationDescCreatedAt := notificationFields[1].Descriptor()
	// notification.DefaultCreatedAt holds the default value on creation for the created_at field.
	notification.DefaultCreatedAt = notificationDescCreatedAt.Default.(func() time.Time)
	tweetFields := schema.Tweet{}.Fields()
	_ = tweetFields
	// tweetDescCreatedAt is the schema descriptor for created_at field.
	tweetDescCreatedAt := tweetFields[1].Descriptor()
	// tweet.DefaultCreatedAt holds the default value on creation for the created_at field.
	tweet.DefaultCreatedAt = tweetDescCreatedAt.Default.(func() time.Time)
	// tweetDescUpdatedAt is the schema descriptor for updated_at field.
	tweetDescUpdatedAt := tweetFields[2].Descriptor()
	// tweet.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	tweet.DefaultUpdatedAt = tweetDescUpdatedAt.Default.(func() time.Time)
	// tweet.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	tweet.UpdateDefaultUpdatedAt = tweetDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescMetadata is the schema descriptor for metadata field.
	userDescMetadata := userFields[5].Descriptor()
	// user.DefaultMetadata holds the default value on creation for the metadata field.
	user.DefaultMetadata = userDescMetadata.Default.(map[string]interface{})
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[6].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[7].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
