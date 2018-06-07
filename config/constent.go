package config

func CollectionNames (name string) string {
	Collection := map[string]string{
		"MEETINGROOM": "meetingroom", 
		"MEETINGS": "meetings",
		"USER": "users",
	}
	return Collection[name]
}
