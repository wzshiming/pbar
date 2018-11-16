package run

import (
	"bufio"
	"io"
	"regexp"

	"github.com/wzshiming/cursor"
	"github.com/wzshiming/deepclone"
	"github.com/wzshiming/pbar"
)

func RunBar(reader io.Reader, writer io.Writer, bar *pbar.Marks, reg *regexp.Regexp, titleKey string) error {
	stdin := bufio.NewReader(reader)
	bars := map[string]*pbar.Marks{}
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
		if len(subData) == 0 {
			io.WriteString(writer, "\r")
			writer.Write(line)
			io.WriteString(writer, cursor.RawClearLine())
			io.WriteString(writer, "\r")
			continue
		}

		title := ""
		for i, name := range subName {
			if name != "" {
				val := string(subData[i])
				if name == titleKey {
					title = val
				}
			}
		}

		if _, ok := order[title]; !ok {
			order[title] = len(order)
			bars[title] = deepclone.Clone(bar).(*pbar.Marks)
			io.WriteString(writer, "\n")
		}

		for i, name := range subName {
			if name != "" {
				val := string(subData[i])
				bars[title].Input(name, val)
			}
		}

		off := len(order) - order[title]
		io.WriteString(writer, "\r")
		io.WriteString(writer, cursor.RawMoveUp(uint64(off)))
		io.WriteString(writer, bars[title].String())
		io.WriteString(writer, cursor.RawClearLine())
		io.WriteString(writer, "\r")
		io.WriteString(writer, cursor.RawMoveDown(uint64(off)))
	}
}
