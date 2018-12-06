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
	type info struct {
		bar   *pbar.Marks
		order int
	}
	infos := map[string]*info{}
	subName := reg.SubexpNames()

	keyIndex := 0
	for i, name := range subName {
		if name == titleKey {
			keyIndex = i
			break
		}
	}

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

		title := string(subData[keyIndex])

		if _, ok := infos[title]; !ok {
			infos[title] = &info{
				order: len(infos),
				bar:   deepclone.Clone(bar).(*pbar.Marks),
			}
			io.WriteString(writer, "\n")
		}

		info := infos[title]
		for i, name := range subName {
			if name != "" {
				val := string(subData[i])
				info.bar.Input(name, val)
			}
		}

		off := len(infos) - info.order
		io.WriteString(writer, "\r")
		io.WriteString(writer, cursor.RawMoveUp(uint64(off)))
		io.WriteString(writer, info.bar.String())
		io.WriteString(writer, cursor.RawClearLine())
		io.WriteString(writer, "\r")
		io.WriteString(writer, cursor.RawMoveDown(uint64(off)))
	}
}
