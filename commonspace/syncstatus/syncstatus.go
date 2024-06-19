package syncstatus

import (
	"github.com/anyproto/any-sync/app"
)

const CName = "common.commonspace.syncstatus"

type StatusUpdater interface {
	app.Component

	HeadsChange(treeId string, heads []string)
	HeadsReceive(senderId, treeId string, heads []string)
}
