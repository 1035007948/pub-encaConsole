import { useEffect, useRef, useState } from 'react';
import { Paper, Stack, Group, Text, Badge, ScrollArea } from '@mantine/core';
import { IconUser, IconClock, IconAlertCircle } from '@tabler/icons-react';
import dayjs from 'dayjs';

interface AuditLogItem {
  id: number;
  action: string;
  entity_type: string;
  entity_id: number;
  entity_no: string;
  operator: string;
  changes: Record<string, { old: unknown; new: unknown }>;
  ip_address: string;
  created_at: string;
}

interface AuditLogStreamProps {
  logs: AuditLogItem[];
  autoScroll?: boolean;
  maxHeight?: number;
}

export function AuditLogStream({
  logs,
  autoScroll = true,
  maxHeight = 400,
}: AuditLogStreamProps) {
  const scrollRef = useRef<HTMLDivElement>(null);
  const [isAtBottom, setIsAtBottom] = useState(true);

  useEffect(() => {
    if (autoScroll && isAtBottom && scrollRef.current) {
      scrollRef.current.scrollTop = scrollRef.current.scrollHeight;
    }
  }, [logs, autoScroll, isAtBottom]);

  const handleScroll = () => {
    if (scrollRef.current) {
      const { scrollTop, scrollHeight, clientHeight } = scrollRef.current;
      setIsAtBottom(scrollHeight - scrollTop - clientHeight < 50);
    }
  };

  const getActionColor = (action: string) => {
    switch (action) {
      case 'create':
        return 'green';
      case 'update':
        return 'blue';
      case 'delete':
        return 'red';
      case 'transition':
        return 'orange';
      default:
        return 'gray';
    }
  };

  const renderChanges = (changes: Record<string, { old: unknown; new: unknown }>) => {
    return Object.entries(changes).map(([field, values]) => (
      <Group key={field} gap="xs">
        <Text size="xs" c="dimmed">
          {field}:
        </Text>
        <Text size="xs" style={{ textDecoration: 'line-through' }} c="red">
          {String(values.old)}
        </Text>
        <Text size="xs">→</Text>
        <Text size="xs" c="green">
          {String(values.new)}
        </Text>
      </Group>
    ));
  };

  return (
    <Paper shadow="sm" withBorder>
      <Stack gap={0}>
        <Group justify="space-between" p="sm" style={{ borderBottom: '1px solid #e9ecef' }}>
          <Text fw={500}>审计日志</Text>
          <Badge size="sm">{logs.length} 条记录</Badge>
        </Group>

        <ScrollArea h={maxHeight} ref={scrollRef} onScroll={handleScroll}>
          <Stack gap={0}>
            {logs.map((log, index) => (
              <Paper
                key={log.id}
                p="sm"
                style={{
                  borderBottom: '1px solid #f1f3f5',
                  backgroundColor: index % 2 === 0 ? '#fff' : '#fafafa',
                }}
              >
                <Stack gap="xs">
                  <Group justify="space-between">
                    <Group gap="xs">
                      <Badge size="sm" color={getActionColor(log.action)}>
                        {log.action}
                      </Badge>
                      <Text size="sm" fw={500}>
                        {log.entity_type}
                      </Text>
                      <Text size="sm" c="dimmed">
                        #{log.entity_no}
                      </Text>
                    </Group>
                    <Group gap="xs">
                      <IconUser size={14} />
                      <Text size="xs">{log.operator}</Text>
                      <IconClock size={14} />
                      <Text size="xs" c="dimmed">
                        {dayjs(log.created_at).format('HH:mm:ss')}
                      </Text>
                    </Group>
                  </Group>

                  {Object.keys(log.changes).length > 0 && (
                    <Stack gap={2} pl="md">
                      {renderChanges(log.changes)}
                    </Stack>
                  )}
                </Stack>
              </Paper>
            ))}

            {logs.length === 0 && (
              <Group justify="center" p="xl">
                <IconAlertCircle size={20} />
                <Text c="dimmed">暂无审计日志</Text>
              </Group>
            )}
          </Stack>
        </ScrollArea>
      </Stack>
    </Paper>
  );
}
