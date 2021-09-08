package admin

import (
	"strconv"

	"github.com/phxenix-w/gotestbot/internal/inits"
)

type Clear struct{}

func (c *Clear) Invokes() []string {
	return []string{"clear", "c"}
}

func (c *Clear) Description() string {
	return "Clear command. Clears the last X messages in this channel."
}

func (c *Clear) AdminPermission() bool {
	return true
}

func (c *Clear) Exec(ctx *inits.Context) error {
	//checking the input. if the user does not provide any or a non-integer input, the amount gets set to 1.
	//otherwise just the input
	amount := 0
	if len(ctx.Args) < 1 {
		amount = 1
	} else {
		val, err := strconv.Atoi(ctx.Args[0])
		if err != nil {
			amount = 1
		} else {
			amount = val + 1
		}
	}
	//check if the input is a positive int
	if amount < 1 {
		amount = 1
	}

	var last_message string

	//you can only delete 100 messages at once, so we have to loop this code here
	for amount > 100 {
		//gets the message array for the correct amount, have to have empty args apparently in go?
		messages, err := ctx.Session.ChannelMessages(ctx.Message.ChannelID, 100, last_message, "", "")
		if err != nil {
			return err
		}

		//makes an empty array and appends the message IDs to it for the delete function
		var message_IDs []string
		for _, message := range messages {
			message_IDs = append(message_IDs, string(message.ID))
		}

		//finally deletes the messages with the IDs from above
		err = ctx.Session.ChannelMessagesBulkDelete(ctx.Message.ChannelID, message_IDs)
		if err != nil {
			return err
		}

		//pass in the last message ID into the variable so the ChannelMessages function looks for the correct messages, not the ones that were already deleted
		last_message = message_IDs[len(message_IDs)-1]
		//and finally subtracts 100 from the amount
		amount -= 100
	}

	//the code from above is repeated for the last remaining amount of messages, under 100 so we pass in the amount variable
	messages, err := ctx.Session.ChannelMessages(ctx.Message.ChannelID, amount, last_message, "", "")
	if err != nil {
		return err
	}

	var message_IDs []string
	for _, message := range messages {
		message_IDs = append(message_IDs, string(message.ID))
	}

	err = ctx.Session.ChannelMessagesBulkDelete(ctx.Message.ChannelID, message_IDs)
	if err != nil {
		return err
	}

	return nil

}
