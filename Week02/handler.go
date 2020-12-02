package httpserver

import (
	"charge/grpcclient"
	"charge/zaplog"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type chargeCallback struct {
	amount    int32
	attach    string
	orderid   string
	productid string

	roleid    uint64
	timestamp uint64
	userid    string
}

func (h *HTTPSvr) Charge(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r != nil && r.Body != nil {
			_ = r.Body.Close()
		}
	}()

	by, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	c := &chargeCallback{}
	err = json.Unmarshal(by, c)
	if err == nil {
		order := &PayOrder{
			Platform:  1,
			OrderId:   c.orderid,
			UserId:    c.userid,
			RoleId:    c.roleid,
			Channel:   1,
			ProductId: "maple10",
			Amount:    c.amount,
			Timestamp: c.timestamp,
			Status:    1,
		}
		if err := order.Add(); err != nil {
			zaplog.Log.Error("", zap.Error(err))
			// ret failed TODO
			return
		}
		grpcclient.Charge(c.roleid, c.amount)
	}
}
