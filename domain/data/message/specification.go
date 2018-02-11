package message

type Specification interface {
	IsSatisfiedBy(mess *Message) bool
}

// PostSpecification 投稿する場所に応じた投稿仕様
type PostSpecification struct {
	MinLength int
	MaxLength int
}

// IsSatisfiedBy 投稿されたメッセージテキストがルールを準拠しているか確認
func (s *PostSpecification) IsSatisfiedBy(mess *Message) bool {
	if len([]rune(mess.Text)) < s.MinLength {
		return false
	}

	if len([]rune(mess.Text)) > s.MaxLength {
		return false
	}

	return true
}
