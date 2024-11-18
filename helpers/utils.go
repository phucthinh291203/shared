package helpers

import (
	"context"
	"regexp"
	"strings"
	"time"
)

//chứa các hàm tiện ích được sử dụng trong suốt application

func NewCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 6*time.Second)
	// giới hạn 6 giây thực hiện một yêu cầu hoặc thao tác
}

func ContainsImageFormat(input string) bool {
	// Các định dạng ảnh phổ biến
	pattern := `(?i)\.(jpg|jpeg|png|gif|bmp|tiff|svg|webp)$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(strings.ToLower(input))
}

func Pagination(page int, limited int) (int, int) {
	if page <= 0 {
		page = 1
	}

	if limited <= 0 {
		limited = 5
	}

	//Bỏ qua số lượng documents -> trang 1 thì ko bỏ qua, trang 2 bỏ qua số documents ở trang 1
	offset := (page - 1) * limited
	return limited, offset
}
