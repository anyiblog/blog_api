package api

import (
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"sort"
	"strings"
)

func WechatPush(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")
	selfSignature := WechatPushMakeSignature(timestamp, nonce)

	if signature == selfSignature {
		c.String(200, echostr)
	} else {
		c.String(200, "Wechat Message Service: this http request is not from Wechat platform!")
	}
}

// WechatPushMakeSignature 推送自签名
func WechatPushMakeSignature(timestamp, nonce string) string {
	sl := []string{os.Getenv("Wechat_Push_Token"), timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	_, _ = io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}
