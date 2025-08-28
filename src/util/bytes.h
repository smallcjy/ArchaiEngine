#pragma once
#include <vector>
#include <cstdint>
#include <stdexcept>

using BYTES = unsigned char*;

uint16_t read_uint16_bytes(const BYTES bytes, size_t n) {
    if (n < 2) throw std::runtime_error("Not enough bytes");
    return (static_cast<uint16_t>(bytes[0]) << 8) | static_cast<uint16_t>(bytes[1]);
}

uint32_t read_uint32_bytes(const BYTES bytes, size_t n) {
    if (n < 4) throw std::runtime_error("Not enough bytes");
    return (static_cast<uint32_t>(bytes[0]) << 24) |
           (static_cast<uint32_t>(bytes[1]) << 16) |
           (static_cast<uint32_t>(bytes[2]) << 8)  |
           static_cast<uint32_t>(bytes[3]);
}

uint64_t read_uint64_bytes(const BYTES bytes, size_t n) {
    if (n < 8) throw std::runtime_error("Not enough bytes");
    return (static_cast<uint64_t>(bytes[0]) << 56) |
           (static_cast<uint64_t>(bytes[1]) << 48) |
           (static_cast<uint64_t>(bytes[2]) << 40) |
           (static_cast<uint64_t>(bytes[3]) << 32) |
           (static_cast<uint64_t>(bytes[4]) << 24) |
           (static_cast<uint64_t>(bytes[5]) << 16) |
           (static_cast<uint64_t>(bytes[6]) << 8)  |
           static_cast<uint64_t>(bytes[7]);
}