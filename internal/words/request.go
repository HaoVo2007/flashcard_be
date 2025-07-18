package words

type CreateWordRequest struct {
	TopicID    string  `json:"topic_id" bson:"topic_id"`
	Word       string  `json:"word" bson:"word"`
	Definition string  `json:"definition" bson:"definition"`
	Example    *string `json:"example" bson:"example"`
	WordType   string  `json:"word_type" bson:"word_type"`
}

type UpdateWordRequest struct {
	Word       *string `json:"word" bson:"word"`
	Definition *string `json:"definition" bson:"definition"`
	Example    *string `json:"example" bson:"example"`
	WordType   *string `json:"word_type" bson:"word_type"`
	IsTrue     *bool   `json:"is_true" bson:"is_true"`
}

type SearchWordRequest struct {
	TopicID *string `json:"topic_id" bson:"topic_id"`
	Word    *string `json:"word" bson:"word"`
	Istrue  *bool   `json:"is_true" bson:"is_true"`
}
