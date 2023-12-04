package main

import (
	"fmt"
    "log"
	"os"

	"github.com/boltdb/bolt"
	"github.com/heyymrdj/Boltlib/pkg/database"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "BoltCLI"}

var createCmd = &cobra.Command{
    Use:   "create",
    Short: "Create a new key-value pair",
    Run:   create,
}

var getCmd = &cobra.Command{
    Use:   "get",
    Short: "Get a key's value",
    Run:   get,
}

var updateCmd = &cobra.Command{
    Use:   "update",
    Short: "Update an existing key",
    Run:   update,
}

var deleteCmd = &cobra.Command{
    Use:   "delete",
    Short: "Delete an existing key",
    Run:   delete,
}

var getAllCmd = &cobra.Command{
    Use:   "getall",
    Short: "Get all keys and values",
    Run:   getAll,
}

func create(cmd *cobra.Command, args []string) {
    db, _ := bolt.Open("my.db", 0600, nil)
    defer db.Close()
    err := database.CreateKey(db, args[0], args[1])
    if err != nil {
        log.Fatal(err)
    }
}

func get(cmd *cobra.Command, args []string) {
    db, _ := bolt.Open("my.db", 0600, nil)
    defer db.Close()
    key, err := database.GetKey(db, args[0])
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(key)
}

func update(cmd *cobra.Command, args []string) {
    db, _ := bolt.Open("my.db", 0600, nil)
    defer db.Close()
    err := database.UpdateKey(db, args[0], args[1])
    if err != nil {
        log.Fatal(err)
    }
}

func delete(cmd *cobra.Command, args []string) {
    db, _ := bolt.Open("my.db", 0600, nil)
    defer db.Close()
    err := database.DeleteKey(db, args[0])
    if err != nil {
        log.Fatal(err)
    }
}

func getAll(cmd *cobra.Command, args []string) {
    db, _ := bolt.Open("my.db", 0600, nil)
    defer db.Close()

    keys, err := database.GetAllKey(db)
    if err != nil {
        log.Fatal(err)
    }
    for k, v := range(keys) {
        fmt.Println(k, v)
    }
}

func main() {
    rootCmd.AddCommand(createCmd)
    rootCmd.AddCommand(getCmd)
    rootCmd.AddCommand(updateCmd)
    rootCmd.AddCommand(deleteCmd)
    rootCmd.AddCommand(getAllCmd)

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    return 
}
