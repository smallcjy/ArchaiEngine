#pragma once
#include <vector>
#include <cstdint>
#include <stdexcept>

using BYTES = std::vector<unsigned char>;

uint16_t bytes_to_uint16(const BYTES& bytes) {
    if (bytes.size() < 2) throw std::runtime_error("Not enough bytes");
    return (static_cast<uint16_t>(bytes[0]) << 8) | static_cast<uint16_t>(bytes[1]);
}

uint32_t bytes_to_uint32(const BYTES& bytes) {
    if (bytes.size() < 4) throw std::runtime_error("Not enough bytes");
    return (static_cast<uint32_t>(bytes[0]) << 24) |
           (static_cast<uint32_t>(bytes[1]) << 16) |
           (static_cast<uint32_t>(bytes[2]) << 8)  |
           static_cast<uint32_t>(bytes[3]);
}

uint64_t bytes_to_uint64(const BYTES& bytes) {
    if (bytes.size() < 8) throw std::runtime_error("Not enough bytes");
    return (static_cast<uint64_t>(bytes[0]) << 56) |
           (static_cast<uint64_t>(bytes[1]) << 48) |
           (static_cast<uint64_t>(bytes[2]) << 40) |
           (static_cast<uint64_t>(bytes[3]) << 32) |
           (static_cast<uint64_t>(bytes[4]) << 24) |
           (static_cast<uint64_t>(bytes[5]) << 16) |
           (static_cast<uint64_t>(bytes[6]) << 8)  |
           static_cast<uint64_t>(bytes[7]);
}