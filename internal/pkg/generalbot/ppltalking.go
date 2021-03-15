package generalbot

import (
	"github.com/gempir/go-twitch-irc/v2"
)

func (gn *Generalbot) salute(message twitch.PrivateMessage) {

	switch message.User.DisplayName {
	// case "Nightbot":
	// 	message := "Nightbot calla la boca, equí 'l bot soy yo"
	// 	gn.Com.Client.Say(gn.Com.Channel, message)
	// 	break
	// case "JavvyoYT":
	// 	if !gn.JavvyoYTesta {
	// 		message := "Hola JavvyoYT, collaciu youtuberil, afayate!"
	// 		gn.Com.Client.Say(gn.Com.Channel, message)
	// 		gn.JavvyoYTesta = true
	// 	}
	// 	break
	// case "HannyaYT":
	// 	if !gn.HannyaYTesta {
	// 		message := "Hola HannyaYT, collacia youtuberil, afayate!"
	// 		gn.Com.Client.Say(gn.Com.Channel, message)
	// 		gn.HannyaYTesta = true
	// 	}
	// 	break
	// case "Trollchu":
	// 	if !gn.Trollchuesta {
	// 		message := "HOMEEE Gonzalooo, afayate!"
	// 		gn.Com.Client.Say(gn.Com.Channel, message)
	// 		gn.Trollchuesta = true
	// 	}
	// 	break
	// case "Haz_A":
	// 	if !gn.Haz_Aesta {
	// 		message := "Hola Haz_A, afayate!"
	// 		gn.Com.Client.Say(gn.Com.Channel, message)
	// 		gn.Haz_Aesta = true
	// 	}
	// 	break
	// case "chisssei":
	// 	if !gn.chisseiesta {
	// 		message := "Hola chisssei, collacia hot, afayate!"
	// 		gn.Com.Client.Say(gn.Com.Channel, message)
	// 		gn.chisseiesta = true
	// 	}
	// 	break
	// case "zaraaify":
	// 	if !gn.zaraaify {
	// 		message := "HeyGuys zaraaify, collacia, afayate!"
	// 		gn.Com.Client.Say(gn.Com.Channel, message)
	// 		gn.zaraaify = true
	// 	}
	// 	break
	// case "MarianaMarrana":
	// 	if !gn.mariana {
	// 		message := "Hola Mariana, collacia, afayate!"
	// 		gn.Com.Client.Say(gn.Com.Channel, message)
	// 		gn.mariana = true
	// 	}
	// 	break
	// case "miamaguila":
	// 	if !gn.miamaguila {
	// 		message := "Miameguia, collacia, afayate!"
	// 		gn.Com.Client.Say(gn.Com.Channel, message)
	// 		gn.miamaguila = true
	// 	}
	// 	break
	case "MickDiaz_":
		if !gn.gallegu {
			message := "⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿ ⣿⣿⣿⣿⣿⠀⠶⣾⠀⣿⠀⠶⢶⡇⠐⠤⣵⣦⠀⣶⣿⢁⡀⢿⣿⣿⣿⣿⣿⣿ ⣿⣿⣿⣿⣿⣀⣿⣿⣀⣿⣀⣛⣛⣏⡙⣂⣼⣿⣀⣿⣇⣠⣤⣘⣿⣿⣿⣿⣿⣿ ⣿⡿⠋⣉⠙⣿⡏⠉⢿⡏⢹⣿⣿⠉⣿⣿⠉⣉⣉⣿⠋⣉⠙⣿⡏⠉⢿⣿⣿⣿ ⣿⣇⠸⠏⠉⡿⠀⠃⠘⡇⠸⠿⢿⠀⠿⢿⠀⠤⠤⣇⠸⠏⠉⡿⠀⠃⠘⣿⣿⣿ ⣿⣿⣶⣶⣾⣷⣾⣿⣶⣷⣶⣶⣾⣶⣶⣾⣶⣶⣶⣿⣷⣶⣾⣷⣾⣿⣶⣿⣿⣿"
			gn.Com.Client.Say(gn.Com.Channel, message)
			gn.gallegu = true
		}
		break
	}
}
