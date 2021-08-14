package format

import (
	"github.com/tailvh/zrtmp/format/mp4"
	"github.com/tailvh/zrtmp/format/ts"
	"github.com/tailvh/zrtmp/format/rtmp"
	"github.com/tailvh/zrtmp/format/rtsp"
	"github.com/tailvh/zrtmp/format/flv"
	"github.com/tailvh/zrtmp/format/aac"
	"github.com/tailvh/zrtmp/av/avutil"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}

