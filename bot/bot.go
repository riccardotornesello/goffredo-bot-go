package bot

import (
	"fmt"
	"path/filepath"
	"sync"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

func Run(wg *sync.WaitGroup) {
	botToken, _ := beego.AppConfig.String("discord::bottoken")

	discord, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	discord.AddHandler(userJoinHandler)

	// Open Websocket
	err = discord.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer discord.Close()
	defer fmt.Println("Bot shut down.")

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	wg.Wait()
}

func userJoinHandler(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	// The user was not logged in
	if v.BeforeUpdate != nil {
		return
	}

	u, err := s.User(v.UserID)
	if err != nil {
		err := fmt.Errorf("error searching user: %w", err)
		fmt.Println(err)
		return
	}

	if u.Bot {
		return
	}

	sound := GetSound(v.UserID)
	if sound == nil {
		return
	}

	dgv, err := s.ChannelVoiceJoin(v.GuildID, v.ChannelID, false, true)
	if err != nil {
		err := fmt.Errorf("error joining voice channel: %w", err)
		fmt.Println(err)
		return
	}

	soundsDirectory, _ := beego.AppConfig.String("soundsDirectory")
	soundIdStr := fmt.Sprintf("%v", sound.Id)
	soundPath := filepath.Join(soundsDirectory, soundIdStr)

	dgvoice.PlayAudioFile(dgv, soundPath, make(chan bool))
	dgv.Disconnect()
}
