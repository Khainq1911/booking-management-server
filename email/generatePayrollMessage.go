package email

import (
	"fmt"
	"time"
)

func GeneratePayrollMessage(employeeName string, employeeId int, payPeriodStart, payPeriodEnd string,
	baseSalary, overtime, bonus, deductions, netSalary float64,
	paymentDate time.Time, notes string) string {

	return fmt.Sprintf(
		`Kính gửi %s,

Chúng tôi xin gửi đến bạn bảng lương chi tiết cho kỳ lương từ ngày %s đến ngày %s:

- 🆔 Mã nhân viên: %d
- 💰 Lương cơ bản: %.2f VND
- ⏳ Lương làm thêm giờ: %.2f VND
- 🎉 Thưởng: %.2f VND
- 📉 Các khoản khấu trừ: %.2f VND
- 💳 Lương thực nhận: **%.2f VND**

📅 Ngày thanh toán: %s

📌 Ghi chú: %s

Nếu có bất kỳ thắc mắc nào về bảng lương, vui lòng liên hệ với bộ phận nhân sự.

Trân trọng,
📍 Khách sạn XYZ
📞 0123-456-789
📧 hr@khachsanxyz.com`,
		employeeName, payPeriodStart, payPeriodEnd, employeeId,
		baseSalary, overtime, bonus, deductions, netSalary,
		paymentDate.Format("02-01-2006"),
		notes,
	)
}
