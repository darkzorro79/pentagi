// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "pentagi/pkg/observability/langfuse/api/internal"
	time "time"
)

type CreateCommentRequest struct {
	// The id of the project to attach the comment to.
	ProjectId string `json:"projectId" url:"-"`
	// The type of the object to attach the comment to (trace, observation, session, prompt).
	ObjectType string `json:"objectType" url:"-"`
	// The id of the object to attach the comment to. If this does not reference a valid existing object, an error will be thrown.
	ObjectId string `json:"objectId" url:"-"`
	// The content of the comment. May include markdown. Currently limited to 500 characters.
	Content string `json:"content" url:"-"`
	// The id of the user who created the comment.
	AuthorUserId *string `json:"authorUserId,omitempty" url:"-"`
}

type CommentsGetRequest struct {
	// Page number, starts at 1.
	Page *int `json:"-" url:"page,omitempty"`
	// Limit of items per page. If you encounter api issues due to too large page sizes, try to reduce the limit
	Limit *int `json:"-" url:"limit,omitempty"`
	// Filter comments by object type (trace, observation, session, prompt).
	ObjectType *string `json:"-" url:"objectType,omitempty"`
	// Filter comments by object id. If objectType is not provided, an error will be thrown.
	ObjectId *string `json:"-" url:"objectId,omitempty"`
	// Filter comments by author user id.
	AuthorUserId *string `json:"-" url:"authorUserId,omitempty"`
}

type Comment struct {
	Id           string            `json:"id" url:"id"`
	ProjectId    string            `json:"projectId" url:"projectId"`
	CreatedAt    time.Time         `json:"createdAt" url:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt" url:"updatedAt"`
	ObjectType   CommentObjectType `json:"objectType" url:"objectType"`
	ObjectId     string            `json:"objectId" url:"objectId"`
	Content      string            `json:"content" url:"content"`
	AuthorUserId *string           `json:"authorUserId,omitempty" url:"authorUserId,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *Comment) GetId() string {
	if c == nil {
		return ""
	}
	return c.Id
}

func (c *Comment) GetProjectId() string {
	if c == nil {
		return ""
	}
	return c.ProjectId
}

func (c *Comment) GetCreatedAt() time.Time {
	if c == nil {
		return time.Time{}
	}
	return c.CreatedAt
}

func (c *Comment) GetUpdatedAt() time.Time {
	if c == nil {
		return time.Time{}
	}
	return c.UpdatedAt
}

func (c *Comment) GetObjectType() CommentObjectType {
	if c == nil {
		return ""
	}
	return c.ObjectType
}

func (c *Comment) GetObjectId() string {
	if c == nil {
		return ""
	}
	return c.ObjectId
}

func (c *Comment) GetContent() string {
	if c == nil {
		return ""
	}
	return c.Content
}

func (c *Comment) GetAuthorUserId() *string {
	if c == nil {
		return nil
	}
	return c.AuthorUserId
}

func (c *Comment) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *Comment) UnmarshalJSON(data []byte) error {
	type embed Comment
	var unmarshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"createdAt"`
		UpdatedAt *internal.DateTime `json:"updatedAt"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = Comment(unmarshaler.embed)
	c.CreatedAt = unmarshaler.CreatedAt.Time()
	c.UpdatedAt = unmarshaler.UpdatedAt.Time()
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *Comment) MarshalJSON() ([]byte, error) {
	type embed Comment
	var marshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"createdAt"`
		UpdatedAt *internal.DateTime `json:"updatedAt"`
	}{
		embed:     embed(*c),
		CreatedAt: internal.NewDateTime(c.CreatedAt),
		UpdatedAt: internal.NewDateTime(c.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (c *Comment) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CommentObjectType string

const (
	CommentObjectTypeTrace       CommentObjectType = "TRACE"
	CommentObjectTypeObservation CommentObjectType = "OBSERVATION"
	CommentObjectTypeSession     CommentObjectType = "SESSION"
	CommentObjectTypePrompt      CommentObjectType = "PROMPT"
)

func NewCommentObjectTypeFromString(s string) (CommentObjectType, error) {
	switch s {
	case "TRACE":
		return CommentObjectTypeTrace, nil
	case "OBSERVATION":
		return CommentObjectTypeObservation, nil
	case "SESSION":
		return CommentObjectTypeSession, nil
	case "PROMPT":
		return CommentObjectTypePrompt, nil
	}
	var t CommentObjectType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CommentObjectType) Ptr() *CommentObjectType {
	return &c
}

type CreateCommentResponse struct {
	// The id of the created object in Langfuse
	Id string `json:"id" url:"id"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateCommentResponse) GetId() string {
	if c == nil {
		return ""
	}
	return c.Id
}

func (c *CreateCommentResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateCommentResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateCommentResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateCommentResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateCommentResponse) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type GetCommentsResponse struct {
	Data []*Comment         `json:"data,omitempty" url:"data,omitempty"`
	Meta *UtilsMetaResponse `json:"meta,omitempty" url:"meta,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetCommentsResponse) GetData() []*Comment {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetCommentsResponse) GetMeta() *UtilsMetaResponse {
	if g == nil {
		return nil
	}
	return g.Meta
}

func (g *GetCommentsResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetCommentsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetCommentsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetCommentsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetCommentsResponse) String() string {
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}
