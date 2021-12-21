package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-template/cmd/global"
	"go-zero-template/cmd/internal/types"
	"go-zero-template/cmd/model"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
)

func OperationRecord(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body []byte
		var userId int
		if r.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(r.Body)
			if err != nil {
				global.GVA_LOG.Error("read body from request error:", zap.Any("err", err))
			} else {
				r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}
		str := r.Header.Get("claims")
		if str != "" {
			var waitUse types.CustomClaims
			err := json.Unmarshal([]byte(str), &waitUse)
			if err != nil {
				httpx.Error(w, err)
				userId = 0
			}
			userId = int(waitUse.ID)
		} else {
			id, err := strconv.Atoi(r.Header.Get("x-user-id"))
			if err != nil {
				userId = 0
			}
			userId = id
		}
		record := types.SysOperationRecord{
			Ip:     r.Host,
			Method: r.Method,
			Path:   r.RequestURI,
			Agent:  r.UserAgent(),
			Body:   string(body),
			UserID: userId,
		}
		// 存在某些未知错误 TODO
		//values := c.Request.Header.Values("content-type")
		//if len(values) >0 && strings.Contains(values[0], "boundary") {
		//	record.Body = "file"
		//}
		//writer := responseBodyWriter{
		//	ResponseWriter: c.Writer,
		//	body:           &bytes.Buffer{},
		//}
		//c.Writer = writer


		next(w, r)


		//record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = 1

		//record.Resp = writer.body.String()

		if err := model.CreateSysOperationRecord(record); err != nil {
			global.GVA_LOG.Error("create operation record error:", zap.Any("err", err))
		}
	}
}

