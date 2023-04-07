package telegram

import (
	"context"
	"fmt"
	"github.com/karakala1427/telegram-bot-pocket-golang/pkg/repository"
)

func (b *Bot) generateAuthorizationLink(chatID int64) (string, error) {
	redirectUrl := b.generateRedirectURL(chatID)
	requestToken, err := b.pocketClient.GetRequestToken(context.Background(), redirectUrl)
	if err != nil {
		return "", err
	}

	if err := b.tokenRepository.Save(chatID, requestToken, repository.RequestTokens); err != nil {
		return "", err
	}

	return b.pocketClient.GetAuthorizationURL(requestToken, redirectUrl)
}

func (b *Bot) generateRedirectURL(chatID int64) string {
	return fmt.Sprintf("%s?chat_id=%d", b.redirectURL, chatID)
}
