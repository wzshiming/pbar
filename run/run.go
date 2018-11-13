package run

import (
	"bufio"
	"io"
	"regexp"

	"github.com/wzshiming/cursor"
	"github.com/wzshiming/pbar"
)

func RunBar(reader io.Reader, writer io.Writer, bar *pbar.Marks, reg *regexp.Regexp, titleKey string) error {
	stdin := bufio.NewReader(reader)
	infos := map[string]*pbar.Info{}
	order := map[string]int{}
	subName := reg.SubexpNames()
	for {
		line, _, err := stdin.ReadLine()
		if err != nil {
			if err != io.EOF {
				return nil
			}
			return err
		}
		subData := reg.FindSubmatch(line)
		title := ""

		if len(subData) == 0 {
			io.WriteString(writer, "\r")
			io.WriteString(writer, cursor.RawClearLine())
			writer.Write(line)
			io.WriteString(writer, "\r")
			continue
		}
		for i, name := range subName {
			if name != "" {
				val := string(subData[i])
				bar.Input(name, val)
				if name == titleKey {
					title = val
				}
			}
		}

		info := infos[title]
		if info == nil {
			info = &pbar.Info{}
			infos[title] = info
			order[title] = len(order)
			io.WriteString(writer, "\r")
			io.WriteString(writer, cursor.RawClearLine())
			io.WriteString(writer, bar.MarkFormat(info))
			io.WriteString(writer, "\n")
		} else {
			off := len(order) - order[title]
			io.WriteString(writer, "\r")
			io.WriteString(writer, cursor.RawMoveUp(uint64(off)))
			io.WriteString(writer, cursor.RawClearLine())
			io.WriteString(writer, bar.MarkFormat(info))
			io.WriteString(writer, "\r")
			io.WriteString(writer, cursor.RawMoveDown(uint64(off)))
		}
	}
}
