package wow

type (
	RaidRank int
	RaidRole string
)

const (
	RaidRankMember    RaidRank = 0
	RaidRankAssistant RaidRank = 1
	RaidRankLeader    RaidRank = 2
)

func GetNumGroupMembers() int {
	return 0
}

func GetRaidRosterInfo(idx int) (name string, rank RaidRank, subgroup, level int, class, fileName, zone string, online, isDead bool, role RaidRole, masterLooter bool) {
	return "", 0, 0, 0, "", "", "", false, false, "", false
}
