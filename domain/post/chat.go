package post

import (
	"context"

	"gitlab.com/shinofara/alpha/domain/user"

	"cloud.google.com/go/firestore"
	"gitlab.com/shinofara/alpha/domain/internal"
	"gitlab.com/shinofara/alpha/domain/type"
	"google.golang.org/api/iterator"
)

type Post struct {
	ID        uint32
	Text      string
	UserID    _type.UserID
	ChannelID _type.ChannelID

	User *user.User        `firestore:"-"`
	ctx  context.Context   `firestore:"-"`
	cli  *firestore.Client `firestore:"-"`
}

// User 投稿したユーザ情報を取得
func (p *Post) GetUser() (*user.User, error) {
	return user.Find(p.UserID)
}

type Repository struct {
	cli *firestore.Client
	ctx context.Context
}

const Collection = "post"

func New(cli *firestore.Client, ctx context.Context) *Repository {
	return &Repository{
		cli: cli,
		ctx: ctx,
	}
}

// Set アイテムを追加する
func (r *Repository) Set(key string, entity *Post) error {
	_, err := r.cli.Collection(Collection).Doc(key).Set(r.ctx, entity)

	return err
}

// Add アイテムを追加するKeyは自動で振られる
func (r *Repository) Add(entity *Post) error {
	_, _, err := r.cli.Collection(Collection).Add(r.ctx, entity)

	return err
}

func (r *Repository) Find(key string) (*Post, error) {
	ref, err := r.cli.Collection(Collection).Doc(key).Get(r.ctx)
	if err != nil {
		return nil, err
	}

	c := new(Post)
	if err := internal.Convert(ref, &c); err != nil {
		return nil, err
	}

	return c, nil
}

func (r *Repository) FindAllByChannelID(id _type.ChannelID) ([]*Post, error) {
	var posts []*Post

	iter := r.cli.Collection(Collection).Documents(r.ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		c := new(Post)
		if err := internal.Convert(doc, &c); err != nil {
			return nil, err
		}
		posts = append(posts, c)
	}

	return posts, nil
}
