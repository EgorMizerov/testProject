package handler

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func (h *Handler) GetSortedHackers(ctx *fasthttp.RequestCtx) {
	res, err := h.repo.Hacker.GetHackers()
	if err != nil {
		ctx.SetBodyString(err.Error())
		return
	}

	resp, err := json.Marshal(res)
	if err != nil {
		ctx.SetBodyString(err.Error())
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetBody(resp)
}
