package validators

import "go-kafka-chat/intenral/domain"

func ValidateMessageContent(content string) error {
	if content == "" {
		return domain.NewDomainValidationError("message is empty")
	}
	return nil
}
