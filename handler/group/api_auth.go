package group

import (
	"github.com/go-iam/db"
	"github.com/go-iam/handler/key"
	"github.com/go-iam/mux"
	"github.com/go-iam/security"
	"net/http"
)

func doAuth(r *http.Request) error {
	keyInfo, err := db.ActiveService().GetKey(mux.Vars(r)["AccessKeyId"])
	if err != nil {
		return err
	}

	owner, err := key.GetKeyOwner(keyInfo.KeyId, keyInfo.CreatorType)
	if err != nil {
		return err
	}

	resource := "ccs:iam:*:" + owner + ":user/*"
	return security.DoAuth(r.Method, "CreateIamUser", resource, mux.Vars(r))
}
