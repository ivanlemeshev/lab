const std = @import("std");

pub fn isLeapYear(year: u32) bool {
    if (year % 100 == 0) {
        return year % 400 == 0;
    }
    return year % 4 == 0;
}

test "isLeapYear" {
    try std.testing.expectEqual(@as(bool, true), isLeapYear(2024));
    try std.testing.expectEqual(@as(bool, true), isLeapYear(2004));
    try std.testing.expectEqual(@as(bool, true), isLeapYear(2096));
    try std.testing.expectEqual(@as(bool, false), isLeapYear(2005));
    try std.testing.expectEqual(@as(bool, false), isLeapYear(2025));
    try std.testing.expectEqual(@as(bool, false), isLeapYear(1997));
    try std.testing.expectEqual(@as(bool, false), isLeapYear(1900));
    try std.testing.expectEqual(@as(bool, true), isLeapYear(2000));
}
