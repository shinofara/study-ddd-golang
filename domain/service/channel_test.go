package service

import (
	"reflect"
	"testing"

	"gitlab.com/shinofara/alpha/domain/channel"
	"gitlab.com/shinofara/alpha/domain/message"
	"gitlab.com/shinofara/alpha/domain/type"
	"gitlab.com/shinofara/alpha/domain/user"
)

func TestChannel_InitialDisplay(t *testing.T) {
	type fields struct {
		channelRepo *channel.Repository
		userRepo    *user.Repository
		messageRepo *message.Repository
	}
	type args struct {
		channelID _type.ChannelID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *channel.Channel
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Channel{
				channelRepo: tt.fields.channelRepo,
				userRepo:    tt.fields.userRepo,
				messageRepo: tt.fields.messageRepo,
			}
			got, err := c.InitialDisplay(tt.args.channelID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Channel.InitialDisplay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Channel.InitialDisplay() = %v, want %v", got, tt.want)
			}
		})
	}
}
