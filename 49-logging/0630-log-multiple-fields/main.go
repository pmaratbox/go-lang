package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"sort"
	"strings"
)

func main() {
	var buf bytes.Buffer
	h := slog.NewJSONHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(g []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	})
	log := slog.New(h)

	log.Info("request", "method", "GET", "status", 200)

	internal := map[string]bool{"level": true, "msg": true, "time": true}
	norm := map[string]string{"INFO": "info", "WARN": "warn", "ERROR": "error", "DEBUG": "debug"}
	for _, line := range strings.Split(strings.TrimSpace(buf.String()), "\n") {
		if line == "" {
			continue
		}
		var d map[string]any
		json.Unmarshal([]byte(line), &d)
		var keys []string
		for k := range d {
			if !internal[k] {
				keys = append(keys, k)
			}
		}
		sort.Strings(keys)
		var sb strings.Builder
		fmt.Fprintf(&sb, "%s|%v", norm[d["level"].(string)], d["msg"])
		for _, k := range keys {
			v := d[k]
			if f, ok := v.(float64); ok {
				fmt.Fprintf(&sb, "|%s=%d", k, int(f))
			} else {
				fmt.Fprintf(&sb, "|%s=%v", k, v)
			}
		}
		fmt.Println(sb.String())
	}
}
