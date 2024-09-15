const std = @import("std");

pub fn eggCount(number: usize) usize {
    if (number == 0) {
        return 0;
    }
    return number % 2 + eggCount(number / 2);
}

test "eggCount" {
    try std.testing.expectEqual(@as(usize, 0), eggCount(0));
    try std.testing.expectEqual(@as(usize, 1), eggCount(1));
    try std.testing.expectEqual(@as(usize, 1), eggCount(16));
    try std.testing.expectEqual(@as(usize, 4), eggCount(89));
}
