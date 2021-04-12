package handler

import (
	"bufio"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func (h *Handler) GetSortedHackers(ctx *fasthttp.RequestCtx) {
	res, err := h.repo.Hacker.GetHackers()
	if err != nil {
		ctx.SetBodyString(err.Error())
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetBodyStreamWriter(func(w *bufio.Writer) {
		json.NewEncoder(w).Encode(&res)
	})
}
