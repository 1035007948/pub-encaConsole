import { Paper, Stack, Group, Text, Badge, Timeline, ScrollArea } from '@mantine/core';
import {
  IconCircleCheck,
  IconCircle,
  IconClock,
  IconAlertCircle,
  IconUser,
} from '@tabler/icons-react';
import dayjs from 'dayjs';

interface TimelineEvent {
  id: number;
  time: string;
  event_type: 'created' | 'updated' | 'transitioned' | 'commented' | 'error';
  title: string;
  description?: string;
  operator: string;
  from_status?: string;
  to_status?: string;
  details?: Record<string, unknown>;
}

interface TimelinePanelProps {
  events: TimelineEvent[];
  title?: string;
  maxHeight?: number;
}

export function TimelinePanel({
  events,
  title = '时间线',
  maxHeight = 500,
}: TimelinePanelProps) {
  const getEventIcon = (type: string) => {
    switch (type) {
      case 'created':
        return <IconCircleCheck size={16} />;
      case 'updated':
        return <IconCircle size={16} />;
      case 'transitioned':
        return <IconClock size={16} />;
      case 'error':
        return <IconAlertCircle size={16} />;
      default:
        return <IconCircle size={16} />;
    }
  };

  const getEventColor = (type: string) => {
    switch (type) {
      case 'created':
        return 'green';
      case 'updated':
        return 'blue';
      case 'transitioned':
        return 'orange';
      case 'error':
        return 'red';
      default:
        return 'gray';
    }
  };

  const getStatusBadge = (status?: string) => {
    if (!status) return null;

    const colorMap: Record<string, string> = {
      draft: 'gray',
      pending: 'yellow',
      in_progress: 'blue',
      completed: 'green',
      rejected: 'red',
      archived: 'cyan',
    };

    const labelMap: Record<string, string> = {
      draft: '草稿',
      pending: '待处理',
      in_progress: '进行中',
      completed: '已完成',
      rejected: '已驳回',
      archived: '已归档',
    };

    return (
      <Badge size="sm" color={colorMap[status] || 'gray'}>
        {labelMap[status] || status}
      </Badge>
    );
  };

  return (
    <Paper shadow="sm" withBorder>
      <Stack gap={0}>
        <Group justify="space-between" p="sm" style={{ borderBottom: '1px solid #e9ecef' }}>
          <Text fw={500}>{title}</Text>
          <Badge size="sm">{events.length} 个事件</Badge>
        </Group>

        <ScrollArea h={maxHeight}>
          <Stack p="md">
            <Timeline
              active={events.length}
              bulletSize={24}
              lineWidth={2}
              color="blue"
            >
              {events.map((event) => (
                <Timeline.Item
                  key={event.id}
                  bullet={getEventIcon(event.event_type)}
                  title={
                    <Group gap="xs">
                      <Text size="sm" fw={500}>
                        {event.title}
                      </Text>
                      {event.from_status && event.to_status && (
                        <Group gap={4}>
                          {getStatusBadge(event.from_status)}
                          <Text size="xs">→</Text>
                          {getStatusBadge(event.to_status)}
                        </Group>
                      )}
                    </Group>
                  }
                >
                  <Stack gap="xs">
                    <Group gap="xs">
                      <IconUser size={12} />
                      <Text size="xs" c="dimmed">
                        {event.operator}
                      </Text>
                      <Text size="xs" c="dimmed">
                        •
                      </Text>
                      <Text size="xs" c="dimmed">
                        {dayjs(event.time).format('YYYY-MM-DD HH:mm:ss')}
                      </Text>
                    </Group>

                    {event.description && (
                      <Text size="sm" c="dimmed">
                        {event.description}
                      </Text>
                    )}

                    {event.details && Object.keys(event.details).length > 0 && (
                      <Stack gap={2} pl="sm">
                        {Object.entries(event.details).map(([key, value]) => (
                          <Group key={key} gap="xs">
                            <Text size="xs" c="dimmed">
                              {key}:
                            </Text>
                            <Text size="xs">{String(value)}</Text>
                          </Group>
                        ))}
                      </Stack>
                    )}
                  </Stack>
                </Timeline.Item>
              ))}
            </Timeline>

            {events.length === 0 && (
              <Group justify="center" py="xl">
                <Text c="dimmed">暂无事件记录</Text>
              </Group>
            )}
          </Stack>
        </ScrollArea>
      </Stack>
    </Paper>
  );
}
