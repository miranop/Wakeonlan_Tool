/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register [MAC]",
	Short: "Register the MAC address you want to use.",
	Long: `To make WOL easier, I thought the best thing to do would be to register it in a database 
	so that it could be called easily.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		address := args[0]
		Register(address)

	},
}

func Register(mac string) {
	db, err := sql.Open("sqlite3", "./data.sql")
	if err != err {
		log.Fatal(err)
	}
	defer db.Close()

	createTable := `CREATE TABLE IF NOT EXISTS macadr (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"mac_address" TEXT NOT NULL UNIQUE
		);`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal()
	}

	registerMAC := `INSERT INTO macadr (mac_address) VALUES(?) `

	address := mac

	_, err = db.Exec(registerMAC, address)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MACアドレスが正常に登録されました")
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
