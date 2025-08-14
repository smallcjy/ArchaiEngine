// Windows下优先级队列驱动定时器系统
#include <windows.h>
#include <functional>
#include <stdio.h>
#include <queue>
#include <vector>
#include <thread>
#include <event2/event.h>

// 定时器任务回调类型
using TimerCallback = std::function<void()>;

struct TimerTask
{
    size_t tid;
    size_t timestamp;
    TimerCallback callback;
    bool is_cancel_;
};

std::function<bool(const TimerTask &, const TimerTask &)> CMP = [](const TimerTask &a, const TimerTask &b)
{
    return a.timestamp > b.timestamp;
};

class QueueTimer
{
public:
    QueueTimer(size_t precision) : task_queue_(CMP), next_task_id_(1), precision_(precision) {}
    ~QueueTimer() = default;

    void start() {

    }

    void update()
    {
        auto now = GetTickCount();
        while (!task_queue_.empty())
        {
            TimerTask *task = task_queue_.top();
            task_queue_.pop();

            if (task->timestamp <= now)
            {
                if (task->is_cancel_)
                {
                    task_map_.erase(task->tid);
                    delete task;
                    break;
                }
                task->callback();
            }
            else
            {
                break;
            }
        }
    }

    // 添加定时任务 :  毫秒级
    size_t add_task(size_t delay_ms, TimerCallback cb)
    {
        size_t tid = next_task_id_++;
        size_t timestamp = GetTickCount() + delay_ms;
        TimerTask *task = new TimerTask({tid, timestamp, cb, false});
        task_queue_.push(task);
        task_map_[tid] = task;
        return tid;
    }

    void cancel_task(size_t tid)
    {
        auto it = task_map_.find(tid);
        if (it != task_map_.end())
        {
            TimerTask *task = it->second;
            task->is_cancel_ = true;
        }
    }

private:
    event* ev_;
    size_t next_task_id_;
    HANDLE timer_handle_;
    std::priority_queue<TimerTask, std::vector<TimerTask *>, decltype(CMP)> task_queue_;
    std::unordered_map<size_t, TimerTask *> task_map_;

    size_t precision_;
};