#pragma once

#include "icomponent.h"
#include "combat/skill/iskill.h"
#include "combat/attr_set.h"

#include "proto/space_service.pb.h"

#include "sync_array.h"
#include "property.h"

#include <vector>
#include <map>
#include <string_view>

class SkillInfo
{
public:
    void init(int skill_id, const std::string& anim, int cost_mana, int cool_down, bool local_predicated) {
        _skill_id = skill_id;
        _anmimator_state = anim;
        _cost_mana = cost_mana;
        _cool_down = cool_down;
        _local_predicated = local_predicated;
        _next_cast_time = 0;
    }

    bool instance_per_entity = true;

    void net_serialize(OutputBitStream& bs, bool to_self) const {
        bs.write(_skill_id);
        bs.write(_anmimator_state);
        bs.write(_cost_mana);
        bs.write(_cool_down);
        bs.write(_next_cast_time);
        bs.write(_local_predicated);
    }

    bool net_delta_serialize(OutputBitStream& bs, bool to_self) {
        bool dirty = false;
        bs.write(_dirty_flag);
        if (_dirty_flag) {
            dirty = true;

            WRITE_IF_DIRTY(skill_id);
            WRITE_IF_DIRTY(anmimator_state);
            WRITE_IF_DIRTY(cost_mana);
            WRITE_IF_DIRTY(cool_down);
            WRITE_IF_DIRTY(next_cast_time);
            WRITE_IF_DIRTY(local_predicated);

            _dirty_flag = 0;
        }
        return dirty;
    }

public:
    enum class DirtyFlag : uint8_t {
        skill_id = 1,
        anmimator_state = 2,
        cost_mana = 4,
        cool_down = 8,
        next_cast_time = 16,
        local_predicated = 32,
    };

private:
    uint8_t _dirty_flag = 0;

    INT_PROPERTY(skill_id);
    STR_PROPERTY(anmimator_state);
    INT_PROPERTY(cost_mana);
    // ms
    INT_PROPERTY(cool_down);
    INT_PROPERTY(next_cast_time);
    BOOL_PROPERTY(local_predicated);
};

class CombatComponent : public IComponent {
public:
    static inline const char* COMPONENT_NAME = "CombatComponent";

    CombatComponent(Entity* entity) : IComponent(entity) {}

    void init() override;
    void destroy() override;
    void net_serialize(OutputBitStream& bs, bool to_self) const override;
    bool net_delta_serialize(OutputBitStream& bs, bool to_self) override;
    void reset_dirty() override;

    void normal_attack(int combo_seq);
    void cast_skill(int skill_id);
    bool can_cast_skill(const SkillInfo& info);

    void take_damage(CombatComponent* attacker, int damage);

private:
    ISkill* get_or_create_skill_instance(int skill_id, bool instance_per_entity = true);
    ISkill* create_skill_instance(int skill_id);
    void stop_running_skill();

private:
    AttrSet _attr_set;

    TSyncArray<SkillInfo> _skill_infos;
    std::map<int, ISkill*> _running_skills;

    size_t _next_normal_attack_time = 0;
    int _normal_attack_timer = -1;
};
