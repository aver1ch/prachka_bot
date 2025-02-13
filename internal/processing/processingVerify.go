package processing

import (
	"fmt"
	"laundryBot/internal/errs"
	"sort"
	"strconv"
)

func ProcessRoomNumber(text string) error {
	rooms := []int{118, 119, 203, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 220, 221, 222, 223, 226, 227, 228, 229, 306, 307, 308, 309, 310, 311, 312, 313, 314, 315, 316, 320, 321, 322, 323, 324, 325, 326, 327, 328, 329, 330, 331, 332, 333, 403, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 420, 421, 422, 423, 424, 426, 427, 428, 429, 430, 431, 432, 433, 509, 510, 511, 512, 513, 514, 515, 516, 520, 521, 522, 523, 524, 526, 527, 528, 529, 530, 531, 532, 533}

	num, err := strconv.Atoi(text)
	if err != nil {
		return fmt.Errorf("%w: %v", errs.ErrRoomNumber, err)
	}

	sort.Ints(rooms)

	index := sort.SearchInts(rooms, num)

	if index > len(rooms) || rooms[index] != num {
		return fmt.Errorf("%w: комната %d не найдена", errs.ErrRoomNumber, num)
	}

	return nil
}
