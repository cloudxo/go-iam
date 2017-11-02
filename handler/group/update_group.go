package group

import (
	"github.com/bitly/go-simplejson"
	"github.com/go-iam/context"
	"github.com/go-iam/db"
	"github.com/go-iam/gerror"
	"github.com/go-iam/handler/util"
	"net/http"
)

type UpdateGroupApi struct {
	req      *http.Request
	status   int
	err      error
	group    Group
	newGroup string
}

func (uga *UpdateGroupApi) Parse() {
	params := util.ParseParameters(uga.req)
	uga.group.groupName = params["GroupName"]
	if params["NewGroupName"] != "" {
		uga.newGroup = params["NewGroupName"]
	}
	if params["NewComments"] != "" {
		uga.group.comments = params["NewComments"]
	}
}

func (uga *UpdateGroupApi) Validate() {
	if uga.group.groupName == "" {
		uga.err = MissGroupNameError
		uga.status = http.StatusBadRequest
		return
	}
}

func (uga *UpdateGroupApi) Auth() {
	uga.err = doAuth(uga.req)
	if uga.err != nil {
		uga.status = http.StatusForbidden
	}
}

func (uga *UpdateGroupApi) updateGroup() {
	bean := db.GroupBean{
		GroupName:  uga.group.groupName,
		Comments:   uga.group.comments,
		CreateDate: uga.group.createDate,
	}
	if uga.newGroup != "" {
		bean.GroupName = uga.newGroup
	}
	group, account := uga.group.groupName, uga.group.account
	uga.err = db.ActiveService().UpdateGroup(group, account, &bean)
	if uga.err == db.GroupNotExistError {
		uga.status = http.StatusNotFound
	} else if uga.err == db.GroupExistError {
		uga.status = http.StatusConflict
	} else {
		uga.status = http.StatusInternalServerError
	}
}

func (uga *UpdateGroupApi) Response() {
	json := simplejson.New()
	if uga.err == nil {
		j := uga.group.Json()
		json.Set("Group", j)
	} else {
		json.Set("ErrorMessage", uga.err.Error())
		context.Set(uga.req, "request_error", gerror.NewIAMError(uga.status, uga.err))
	}
	json.Set("RequestId", context.Get(uga.req, "request_id"))
	data, _ := json.Encode()
	context.Set(uga.req, "response", data)
}

func UpdateGroupHandler(w http.ResponseWriter, r *http.Request) {
	uga := UpdateGroupApi{req: r, status: http.StatusOK}
	defer uga.Response()

	if uga.Auth(); uga.err != nil {
		return
	}

	if uga.Parse(); uga.err != nil {
		return
	}

	if uga.Validate(); uga.err != nil {
		return
	}

	gga := GetGroupApi{}
	gga.group.groupName = uga.group.groupName
	gga.group.account = uga.group.account

	if gga.getGroup(); gga.err != nil {
		uga.err = gga.err
		return
	}

	if uga.group.comments == "" {
		uga.group.comments = gga.group.comments
	}

	if uga.updateGroup(); uga.err != nil {
		return
	}
}