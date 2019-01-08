package storage

// More info: http://www.cse.yorku.ca/~oz/hash.html
// https://stackoverflow.com/questions/1579721/why-are-5381-and-33-so-important-in-the-djb2-algorithm
func djb2(key string) uint64 {
	var hash uint64 = 5381

	for i := 0; i < len(key); i++ {
		hash = hash*33 + uint64(key[i])
	}
	return hash
}

func (s *Storage) getShard(key string) *shard {
	shardID := djb2(key) & uint64(len(s.shards)-1)
	return s.shards[shardID]
}
