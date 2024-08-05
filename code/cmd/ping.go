/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/go-ping/ping"
	"github.com/spf13/cobra"
)

var easyMode bool

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping [host]",
	Short: "send ping",
	Long:  `Check if the necessary lines are connected to WOL`,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		var host = args[0]
		pinghost(host, easyMode)

	},
}

func pinghost(host string, easyMode bool) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		panic(err)
	}
	pinger.Count = 5
	//タイムアウトを5秒に設定
	pinger.Timeout = time.Second * 5

	if easyMode {
		fmt.Printf("PING %s (%s):\n", host, pinger.Addr())
		err = pinger.Run()
		if err != nil {
			panic(err)
		}
	} else {
		pinger.OnRecv = func(pkt *ping.Packet) {
			fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
				pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
		}
		pinger.OnFinish = func(stats *ping.Statistics) {
			fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
			fmt.Printf("%d packets transmitted, %d packets received, %.2f%% packet loss\n",
				stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
			fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
				stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
		}

		fmt.Printf("PING %s (%s):\n", host, pinger.Addr())
		err = pinger.Run()
		if err != nil {
			panic(err)
		}

	}

}

func init() {
	pingCmd.Flags().BoolVarP(&easyMode, "easy", "e", false, "表示の簡略化")
	rootCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
