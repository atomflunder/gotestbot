package usercommands

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/phxenix-w/gotestbot/internal/inits"
)

type Dice struct{}

func (c *Dice) Invokes() []string {
	return []string{"dice", "roll", "r"}
}

func (c *Dice) Description() string {
	return "Rolls a dice in NdN format and tells you the results."
}

func (c *Dice) AdminPermission() bool {
	return false
}

func (c *Dice) Exec(ctx *inits.Context) error {
	//the dice will be input in NdN format, like: 2d100
	//so we need to split the first arg to get the sides and amount
	amount_s := ""
	sides_s := ""
	//this checks if the input is in the NdN format
	//otherwise the whole bot exits when we try to assign these values, which is obviously the worst case for a discord bot
	if len(strings.Split(ctx.Args[0], "d")) > 1 {
		amount_s = strings.Split(ctx.Args[0], "d")[0]
		sides_s = strings.Split(ctx.Args[0], "d")[1]
	} else {
		//returns if the input is not valid
		_, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			"Please input a valid dice in the NdN format.")
		if err != nil {
			return err
		}
		return nil
	}

	//then we convert both to integers
	amount, err := strconv.Atoi(amount_s)
	if err != nil {
		return err
	}
	sides, err := strconv.Atoi(sides_s)
	if err != nil {
		return err
	}

	//checks if the input is not stupidly high or worse, below zero
	if sides > 10000 || amount > 1000 || sides < 1 || amount < 1 {
		_, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			"Please do not use any values too high or too low for this dice.")
		if err != nil {
			return err
		}
		return nil
	}

	//setting up the total and individual roll vars
	result := 0
	var result_list []string

	for amount > 0 {
		//first we need a different seed each time we run this loop, since amount changes everytime i use that
		//also we need something that is different each time we run the command, so i use the current time too
		//no clue why i even have to do this in go. i mean wtf
		rand.Seed(int64(amount) * time.Now().UnixNano())
		//then we get the random number
		r := (rand.Intn(sides-1+1) + 1)
		//add this to the total result
		result += r
		//and also append it to the string of results so the user sees each dice roll separately again
		rs := strconv.Itoa(r)
		result_list = append(result_list, rs)
		amount -= 1
	}

	//now we need to convert the result and the list to a string to send it in a message
	result_s := strconv.Itoa(result)
	result_list_s := strings.Join(result_list, ", ")

	//sends all this stuff back to the user in one message
	//we can use the original sides and amount vars since all checks passed and they are correct inputs
	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
		"Rolled a "+sides_s+" sided dice "+amount_s+" times:\nTotal: "+result_s+"\nRolls: "+result_list_s)

	if err != nil {
		return err
	}

	return nil
}
