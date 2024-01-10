package utils

import (
	"strconv"
	"strings"
	"time"
)

/*
这段Go语言代码的主要作用是将传入的时间字符串解析为`time.Duration`类型的值。
定义了一个函数`ParseDuration`，该函数接受一个表示时间的字符串参数`d`，
并返回一个`time.Duration`类型的值和一个错误（如果有）。
*/

func ParseDuration(d string) (time.Duration, error) {
	// 将参数字符串d去除空格，并重新赋值给d。
	d = strings.TrimSpace(d)

	// 使用标准的 time.ParseDuration 函数将字符串`d`解析为 time.Duration 类型的值。
	dr, err := time.ParseDuration(d)
	// 如果解析出错，则错误信息被存储在`err`变量中；否则，将结果存储在`dr`变量中。
	if err == nil {
		// 检查上一步中是否存在错误。如果没有，则直接返回解析结果`dr`和一个`nil`的错误。
		return dr, nil
	}

	// 如果上一步中存在错误，检查字符串`d`中是否包含字符"d"。
	if strings.Contains(d, "d") {

		// 如果包含字符"d"，则找到第一次出现"d"的位置，将其索引存储在`index`变量中。
		index := strings.Index(d, "d")

		// 将从字符串开头到"d"之间的子字符串转换为整数，并存储在`hour`变量中。由于不需要第二个返回值（该函数返回一个错误），因此使用下划线_blank标识符来忽略它。
		hour, _ := strconv.Atoi(d[:index])

		// 将小时数转换为天数，乘以`time.Hour`和24，并将结果存储在`dr`变量中。
		dr = time.Hour * 24 * time.Duration(hour)

		// 将"d"之后的子字符串(`index+1`表示从第一个字符开始，另一个冒号`:`表示只取后面的部分)解析为`time.Duration`类型的值，并存储在`ndr`变量中。如果有错误，则存储在`err`变量中。
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			// 如果解析出错，则直接返回上一步计算得到的`dr`值和一个`nil`的错误。
			return dr, nil
		}

		// 如果解析成功，则将前面计算得到的`dr`值和当前解析得到的`ndr`值相加，并返回该结果和一个`nil`的错误。
		return dr + ndr, nil
	}

	// 如果字符串中不包含"d"字符，则尝试将整个字符串解析为int64类型。如果解析出错，则错误信息存储在`err`变量中；否则，将结果存储在`dv`变量中。
	dv, err := strconv.ParseInt(d, 10, 64)

	// 最终，无论是通过"d"字符分割还是直接将整个字符串转换为int64类型，都将其转换为`time.Duration`类型的值，并返回该值和前面可能存在的错误信息。
	return time.Duration(dv), err
}
