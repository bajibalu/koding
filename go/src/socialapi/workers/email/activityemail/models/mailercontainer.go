package models

import (
	"errors"
	"fmt"

	socialmodels "socialapi/models"
	"socialapi/workers/notification/models"
	"time"
)

type MailerContainer struct {
	Activity        *models.NotificationActivity
	Content         *models.NotificationContent
	AccountId       int64
	Message         string
	Slug            string
	ActivityMessage string
	ObjectType      string
	Group           GroupContent
	CreatedAt       time.Time
}

func NewMailerContainer() *MailerContainer {
	return &MailerContainer{}
}

func (mc *MailerContainer) PrepareContainer() error {
	if err := mc.validateContainer(); err != nil {
		return err
	}

	// if content type not valid return
	contentType, err := mc.Content.GetContentType()
	if err != nil {
		return err
	}

	var target *socialmodels.ChannelMessage

	switch mc.Content.TypeConstant {
	case models.NotificationContent_TYPE_COMMENT:
		target, err = fetchMessageTarget(mc.Activity.MessageId)
		if err != nil {
			return err
		}
		mc.Message = target.Body
		target, err = fetchMessageParent(mc.Activity.MessageId)
		if err != nil {
			return err
		}
	default:
		target, err = fetchMessageTarget(mc.Content.TargetId)
		if err != nil {
			return err
		}
		mc.Message = target.Body
	}

	mc.prepareGroup(target)
	mc.prepareSlug(target)
	mc.prepareObjectType(target)
	mc.CreatedAt = target.CreatedAt

	contentType.SetActorId(target.AccountId)
	contentType.SetListerId(mc.AccountId)

	mc.ActivityMessage = contentType.GetActivity()

	return nil
}

func fetchMessageTarget(messageId int64) (*socialmodels.ChannelMessage, error) {
	target := socialmodels.NewChannelMessage()

	if err := target.ById(messageId); err != nil {
		return nil, err
	}

	return target, nil
}

func fetchMessageParent(messageId int64) (*socialmodels.ChannelMessage, error) {
	mr := socialmodels.NewMessageReply()
	mr.ReplyId = messageId

	return mr.FetchParent()
}

func (mc *MailerContainer) validateContainer() error {
	if mc.AccountId == 0 {
		return errors.New("account id is not set")
	}
	if mc.Activity == nil {
		return errors.New("activity is not set")
	}
	if mc.Content == nil {
		return errors.New("content is not set")
	}

	return nil
}

func (mc *MailerContainer) prepareGroup(cm *socialmodels.ChannelMessage) {
	c := socialmodels.NewChannel()
	if err := c.ById(cm.InitialChannelId); err != nil {
		return
	}
	// TODO fix these Slug and Name
	mc.Group = GroupContent{
		Slug: c.GroupName,
		Name: c.GroupName,
	}
}

func (mc *MailerContainer) prepareSlug(cm *socialmodels.ChannelMessage) {
	switch cm.TypeConstant {
	case socialmodels.ChannelMessage_TYPE_REPLY:
		// TODO we need append something like comment id to parent message slug
		mc.Slug = fetchRepliedMessage(cm.Id).Slug
	case socialmodels.ChannelMessage_TYPE_PRIVATE_MESSAGE:
		mc.Slug = fetchPrivateChannelSlug(cm.Id)
	default:
		mc.Slug = cm.Slug
	}
}

func (mc *MailerContainer) prepareObjectType(cm *socialmodels.ChannelMessage) {
	switch cm.TypeConstant {
	case socialmodels.ChannelMessage_TYPE_POST:
		mc.ObjectType = "post"
	case socialmodels.ChannelMessage_TYPE_REPLY:
		mc.ObjectType = "comment"
	case socialmodels.ChannelMessage_TYPE_PRIVATE_MESSAGE:
		mc.ObjectType = "chat message"
	}
}

func fetchPrivateChannelSlug(messageId int64) string {
	cml := socialmodels.NewChannelMessageList()
	ids, err := cml.FetchMessageChannelIds(messageId)
	if err != nil {
		return ""
	}

	if len(ids) == 0 {
		return ""
	}

	return fmt.Sprintf("Message/%d", ids[0])
}

func fetchRepliedMessage(replyId int64) *socialmodels.ChannelMessage {
	mr := socialmodels.NewMessageReply()
	mr.ReplyId = replyId

	parent, err := mr.FetchParent()
	if err != nil {
		parent = socialmodels.NewChannelMessage()
	}

	return parent
}
