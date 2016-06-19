package wow

type ChatType string

const (
	ChatTypeRaid ChatType = "RAID"
)

type AddonChatType string

const (
	AddonChatTypeRaid AddonChatType = "RAID"
)

func SendChatMessage(msg string, chatType ChatType, channel interface{}) {

}

func SendAddonMessage(prefix, msg string, chatType AddonChatType, target interface{}) {

}

func RegisterAddonMessagePrefix(prefix string) bool {
	return false
}
