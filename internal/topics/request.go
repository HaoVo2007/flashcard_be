package topics

type CreateTopicRequest struct {
	TopicName        string `json:"name" bson:"name"`
	TopicDescription string `json:"description" bson:"description"`
	Color            string `json:"color" bson:"color"`
}

type UpdateTopicRequest struct {
	TopicName        *string `json:"name" bson:"name"`
	TopicDescription *string `json:"description" bson:"description"`
	Color            *string `json:"color" bson:"color"`
}