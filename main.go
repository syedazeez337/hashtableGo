package main

import (
	"fmt"

	"github.com/syedazeez337/hashtableGo/hashtable"
)

func main() {
	ht := hashtable.NewHashTable(5) // Create a hash table with 5 buckets
	fmt.Println("Inserting keys into the hash table:")

	keys := []string{"apple", "banana", "grape", "orange", "kiwi"}
	for _, key := range keys {
		ht.Insert(key, key+"-value") // Insert each key with a simple value
	}

	// Display the content of each bucket
	for i, bucket := range ht.Buckets {
		if bucket.Key != "" {
			fmt.Printf("Bucket %d: [%s: %s]\n", i, bucket.Key, bucket.Value)
		}
	}

	// Delete a key
	keyToDelete := "banana"
	if deleted := ht.Delete(keyToDelete); deleted {
		fmt.Printf("Key '%s' deleted successfully.\n", keyToDelete)
	} else {
		fmt.Printf("Key '%s' not found.\n", keyToDelete)
	}

	// Display the content after deletion
	fmt.Println("\nContent after deletion:")
	for i, bucket := range ht.Buckets {
		if bucket.Key != "" {
			fmt.Printf("Bucket %d: [%s: %s]\n", i, bucket.Key, bucket.Value)
		}
	}

	// Search for the deleted key
	if value, found := ht.Search("banana"); found {
		fmt.Printf("Found banana: %s\n", value)
	} else {
		fmt.Println("Banana not found.")
	}
}
