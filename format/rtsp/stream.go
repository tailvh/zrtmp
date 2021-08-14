package rtsp

import (
	"github.com/tailvh/zrtmp/av"
	"github.com/tailvh/zrtmp/format/rtsp/sdp"
	"time"
)

type Stream struct {
	av.CodecData
	Sdp    sdp.Media
	client *Client

	// h264
	fuStarted  bool
	fuBuffer   []byte
	sps        []byte
	pps        []byte
	spsChanged bool
	ppsChanged bool

	gotpkt    bool
	pkt       av.Packet
	timestamp uint32
	firsttimestamp uint32

	lasttime time.Duration
}

