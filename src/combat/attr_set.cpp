#include "attr_set.h"
#include "math_utils.h"
#include "bit_stream.h"

#include <algorithm>
#include <spdlog/spdlog.h>

void AttrSet::net_serialize(OutputBitStream& bs, bool to_self) const {
    bs.write(_max_health);
    bs.write(_health);
    bs.write(_max_mana);
    bs.write(_mana);
}

bool AttrSet::net_delta_serialize(OutputBitStream& bs, bool to_self) {
    bool dirty = false;
    bs.write(_dirty_flag);
    if (_dirty_flag) {
        dirty = true;

        WRITE_IF_DIRTY(max_health);
        WRITE_IF_DIRTY(health);
        WRITE_IF_DIRTY(max_mana);
        WRITE_IF_DIRTY(mana);
    }
    return dirty;
}

void AttrSet::add_health(int val) {
    int n = _health + val;
    n = std::max(0, std::min(_max_health, n));
    set_health(n);
}

void AttrSet::add_mana(int val) {
    int n = _mana + val;
    n = std::max(0, std::min(_max_mana, n));
    set_mana(n);
}

int AttrSet::take_damage(AttrSet& attacker, int attack_damage) {
    // 1. 基础伤害计算（至少为1）
    int base_damage = (std::max)(1, attack_damage - this->defence);

    // 2. 暴击判断
    bool is_critical = random_01() < attacker.critical_rate;
    int damage = is_critical ?
        static_cast<int>(base_damage * (1.0f + attacker.critical_damage)) : base_damage;

    // 3. 命中与闪避判断
    float hit_chance = attacker.accuracy - this->dodge_rate;
    bool is_hit = random_01() < hit_chance;
    if (!is_hit) {
        damage = 0;
    }

    // 4. 护盾吸收
    int shield_absorb = (std::min)(damage, this->shield);
    this->shield -= shield_absorb;
    damage -= shield_absorb;

    // 5. 实际扣血（确保不小于0）
    add_health(-damage);

    // 6. 吸血效果（不超过最大生命值）
    if (damage > 0 && attacker.lifesteal > 0) {
        int heal_amount = static_cast<int>(damage * attacker.lifesteal);
        attacker.add_health(heal_amount);
    }

    return damage;
}

void AttrSet::clamp_all_attributes() {
    accuracy = std::clamp(accuracy, 0.0f, 1.0f);
    dodge_rate = std::clamp(dodge_rate, 0.0f, 1.0f);
    tenacity = std::clamp(tenacity, 0.0f, 1.0f);
}
