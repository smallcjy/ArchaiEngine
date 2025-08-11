#include<vector>
#include<string>
#include<sstream>
// string类操作方法集

// SplitStringByDelimiter
// 按照分隔符分割字符串 
inline std::vector<std::string> SplitStringByDelimiter(const std::string& str, char delimiter) {
    std::vector<std::string> ret;
    std::string token;
    std::istringstream tokenStream(str);
    while (std::getline(tokenStream, token, delimiter)) {
        ret.push_back(token);
    }
    return ret;
}

inline int String2Int(const std::string& str) {
    int result = 0;
    for (char c : str) {
        result = result * 10 + (c - '0');
    }
    return result;
}

inline int32_t String2Int32(const std::string& str) {
    int32_t result = 0;
    for (char c : str) {
        result = result * 10 + (c - '0');
    }
    return result;
}

inline int64_t String2Int64(const std::string& str) {
    int64_t result = 0;
    for (char c : str) {
        result = result * 10 + (c - '0');
    }
    return result;
}