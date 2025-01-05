package hashtable

// Constants for prime numbers (same as before)
const (
	Prime1 = 31
	Prime2 = 37
)

// Entry represents a single key-value pair in the hash table
type Entry struct {
	Key   string
	Value string
}

// HashTable represents the entire hash table structure
type HashTable struct {
	Buckets []Entry // Array of entries (open addressing)
	Size    int     // Total size of the table
	Count   int     // Number of entries in the table
}

// NewHashTable initializes a new hash table with the given size
func NewHashTable(size int) *HashTable {
	return &HashTable{
		Buckets: make([]Entry, size), // Create an array of empty entries
		Size:    size,
		Count:   0,
	}
}

// Hash function that converts a string into an index for the hash table
func (ht *HashTable) Hash(key string, prime, numBuckets int) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash*prime + int(key[i])) % numBuckets
	}
	return hash
}

// Double Hashing function that handles collisions using two hash functions
func (ht *HashTable) GetHash(key string, attempt int) int {
	hashA := ht.Hash(key, Prime1, ht.Size)    // First hash function
	hashB := ht.Hash(key, Prime2, ht.Size)    // Second hash function
	return (hashA + attempt*(hashB+1)) % ht.Size
}

// Insert adds a new key-value pair to the hash table, resolving collisions
func (ht *HashTable) Insert(key string, value string) {
	for attempt := 0; attempt < ht.Size; attempt++ {
		index := ht.GetHash(key, attempt)
		if ht.Buckets[index].Key == "" { // If the bucket is empty, insert the entry
			ht.Buckets[index] = Entry{Key: key, Value: value}
			ht.Count++
			return
		}
	}
}

// Search retrieves the value for a given key from the hash table
func (ht *HashTable) Search(key string) (string, bool) {
	for attempt := 0; attempt < ht.Size; attempt++ {
		index := ht.GetHash(key, attempt)
		if ht.Buckets[index].Key == key {
			return ht.Buckets[index].Value, true
		}
		if ht.Buckets[index].Key == "" { // If an empty bucket is encountered, stop searching
			break
		}
	}
	return "", false
}

// Delete removes a key-value pair from the hash table
func (ht *HashTable) Delete(key string) bool {
	for attempt := 0; attempt < ht.Size; attempt++ {
		index := ht.GetHash(key, attempt)
		if ht.Buckets[index].Key == key {
			ht.Buckets[index] = Entry{} // Remove the entry by resetting it
			ht.Count--
			return true
		}
		if ht.Buckets[index].Key == "" { // Stop searching if an empty bucket is found
			break
		}
	}
	return false
}
