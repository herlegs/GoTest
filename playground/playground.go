package main

import (
	"fmt"
	"strings"
)

func main() {
	parts := strings.Fields(str)
	res := ""
	for i := 0; i < len(parts); i+=2 {
		res += parts[i]+"\t"+parts[i+1]+"\n"
	}
	fmt.Printf("%v\n",res)
}

func getProductFromPSM(psm string) string {
	if parts := strings.Split(psm, "."); len(parts) == 3 {
		return parts[0]
	}
	return ""
}

var str =`product
psm
toutiao
toutiao.shield.http2thrift_demo
life
life.api.poly
life
life.api.feed
life
life.api.item
maya
maya.share.api
toutiao
toutiao.novel.channel_goapi
toutiao
toutiao.shield.agw_thrift_api_service
h
h.api.client
ugc
toutiao.ugc.vote_publish
video
toutiao.video.biz_service_praise
ugc
toutiao.ugc.publish_goapi
life
vinna.api.app
internal
toutiao.fe.magic
temai
cmp.ecom.ken
toutiao
toutiao.imtoutiao.ceremony
h
h.api.in_api
ugc
toutiao.top.author_growth
toutiao
ttgame.assistant.api
toutiao
toutiao.reading.ugcapi
h
h.api.query
h
h.api.ugc
video
toutiao.video.user_gateway
video
toutiao.video.tv_platform
sandbox
toutiao.shield.agw_thrift_api_service
automobile
motor.pgcapi.content
h
h.api.operation
h
h.api.solution
hybrid
em.mars.api
h
h.api.system
h
h.api.study
h
h.api.guide
h
h.api.study_room
h
h.fe.wechat_node
h
h.api.reach
toutiao
toutiao.user_action.report
h
h.api.play
h
h.api.user
video
toutiao.video.dante
`