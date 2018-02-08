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
		channelRepo channel.Repository
		userRepo    user.Repository
		messageRepo message.Repository
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
		{
			name: "Success",
			fields: fields{
				channelRepo: &ChannelRepositoryMem{},
				userRepo:    &UserRepositoryMem{},
				messageRepo: &MessageRepositoryMem{},
			},
			args: args{},
			want: &channel.Channel{
				ID:      "test_channel_id",
				Name:    "test_channel_name",
				OwnerID: "test_owner_id",
				Owner: &user.User{
					ID:   "test_owner_id",
					Name: "test_owner_name",
				},
				Messages: []*message.Message{
					{
						ID:        "test_mess_id",
						Text:      "test_mess_text",
						UserID:    "test_owner_id",
						ChannelID: "test_channel_id",
					},
				},
			},
			wantErr: false,
		},
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
				t.Errorf("Channel.InitialDisplay() error = %+V, wantErr %+V", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Channel.InitialDisplay() = %+V, want %+V", got, tt.want)
			}
		})
	}
}

type ChannelRepositoryMem struct{}

func (_ ChannelRepositoryMem) Find(id _type.ChannelID) (*channel.Channel, error) {
	return &channel.Channel{
		ID:      "test_channel_id",
		Name:    "test_channel_name",
		OwnerID: "test_owner_id",
	}, nil
}
func (_ ChannelRepositoryMem) Add(c *channel.Channel) (*channel.Channel, error) {
	return nil, nil
}

type UserRepositoryMem struct{}

func (_ UserRepositoryMem) Find(id _type.UserID) (*user.User, error) {
	return &user.User{
		ID:   "test_owner_id",
		Name: "test_owner_name",
	}, nil
}
func (_ UserRepositoryMem) Add(entity *user.User) (*user.User, error) {
	return nil, nil
}

type MessageRepositoryMem struct{}

func (_ MessageRepositoryMem) Set(key string, entity *message.Message) error {
	return nil
}
func (_ MessageRepositoryMem) Add(c *message.Message) (*message.Message, error) {
	return nil, nil
}
func (_ MessageRepositoryMem) Find(id _type.MessageID) (*message.Message, error) {
	return nil, nil
}
func (_ MessageRepositoryMem) FindAllByChannelID(id _type.ChannelID) ([]*message.Message, error) {
	mess := &message.Message{
		ID:        "test_mess_id",
		Text:      "test_mess_text",
		UserID:    "test_owner_id",
		ChannelID: "test_channel_id",
		User:      nil,
	}
	return []*message.Message{mess}, nil
}
