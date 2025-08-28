#include <random>
using Seed = std::mt19937;

namespace archai::rand
{
    // [min, max] 随机整数
    size_t NextUInt(size_t seed, size_t min, size_t max)
    {
        Seed rng(seed);
        std::uniform_int_distribution<size_t> dist(min, max);
        return dist(rng);
    }

    // [min, max) 随机 double
    double NextDouble(size_t seed, double min, double max)
    {
        Seed rng(seed);
        std::uniform_real_distribution<double> dist(min, max);
        return dist(rng);
    }

    // [min, max) 随机 float
    float NextFloat(size_t seed, float min, float max)
    {
        Seed rng(seed);
        std::uniform_real_distribution<float> dist(min, max);
        return dist(rng);
    }

    // 伯努利分布，返回 true 的概率为 p
    bool NextBool(size_t seed, double p = 0.5)
    {
        Seed rng(seed);
        std::bernoulli_distribution dist(p);
        return dist(rng);
    }

    // 正态分布（高斯分布），均值 mean，标准差 stddev
    double NextNormal(size_t seed, double mean = 0.0, double stddev = 1.0)
    {
        Seed rng(seed);
        std::normal_distribution<double> dist(mean, stddev);
        return dist(rng);
    }

    // 泊松分布，参数 lambda
    size_t NextPoisson(size_t seed, double lambda)
    {
        Seed rng(seed);
        std::poisson_distribution<size_t> dist(lambda);
        return dist(rng);
    }

    // 指数分布，参数 lambda
    double NextExponential(size_t seed, double lambda)
    {
        Seed rng(seed);
        std::exponential_distribution<double> dist(lambda);
        return dist(rng);
    }
}
