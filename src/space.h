#pragma once

#include <memory>
#include <string>
#include <vector>
#include <unordered_map>

#include "DetourNavMesh.h"
#include "DetourNavMeshQuery.h"

#include "math_utils.h"

class Entity;
class AOI;

class Space {
public:
    Space(size_t w, size_t h);
    ~Space();

    void join(Entity* entity);
    void leave(Entity* entity);
    bool has_entity(Entity* entity);
    Entity* find_entity(int eid);

    void update_position(int eid, float x, float y, float z);

    void update();

    void call_all(int eid, const std::string& msg_name, const std::string& msg_bytes);
    void call_others(int eid, const std::string& msg_name, const std::string& msg_bytes);

    std::vector<Entity*> find_entities_in_sector(float cx, float cy, float ux, float uy, float r, float theta);
    std::vector<Entity*> find_entities_in_circle(float cx, float cy, float r);

    std::vector<Vector3f> find_path(const Vector3f& start_pos, const Vector3f& end_pos);
    std::vector<Vector3f> navigation_test(const Vector3f& start_pos, const Vector3f& end_pos);

private:
    size_t _width;
    size_t _height;

    int _update_timer;

    std::unordered_map<int, Entity*> _eid_2_entity;

    std::shared_ptr<AOI> _aoi;

    dtNavMesh* _nav_mesh = nullptr;
    dtNavMeshQuery* _nav_query = nullptr;
    dtQueryFilter _nav_filter;

    class Monster* _m = nullptr;
};
