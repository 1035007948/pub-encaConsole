import { useEffect, useRef } from 'react';
import {
  Drawer,
  Stack,
  Group,
  Text,
  Badge,
  Divider,
  Button,
  ScrollArea,
  Timeline,
  Paper,
} from '@mantine/core';
import { IconCalendar, IconUser, IconClock } from '@tabler/icons-react';
import dayjs from 'dayjs';

interface DetailField {
  key: string;
  label: string;
  type?: 'text' | 'badge' | 'date' | 'number' | 'link';
  render?: (value: unknown, item: Record<string, unknown>) => React.ReactNode;
}

interface TimelineItem {
  time: string;
  action: string;
  user: string;
  description?: string;
}

interface DetailDrawerProps {
  opened: boolean;
  onClose: () => void;
  title: string;
  data: Record<string, unknown> | null;
  fields: DetailField[];
  timeline?: TimelineItem[];
  actions?: {
    label: string;
    onClick: () => void;
    color?: string;
    disabled?: boolean;
  }[];
}

export function DetailDrawer({
  opened,
  onClose,
  title,
  data,
  fields,
  timeline,
  actions,
}: DetailDrawerProps) {
  const scrollRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (opened && scrollRef.current) {
      scrollRef.current.scrollTop = 0;
    }
  }, [opened]);

  const renderField = (field: DetailField) => {
    if (!data) return null;

    const value = data[field.key];

    if (field.render) {
      return field.render(value, data);
    }

    switch (field.type) {
      case 'badge':
        return (
          <Badge
            color={
              value === '已完成' || value === 'completed'
                ? 'green'
                : value === '待处理' || value === 'pending'
                ? 'yellow'
                : value === '已驳回' || value === 'rejected'
                ? 'red'
                : 'blue'
            }
          >
            {String(value || '-')}
          </Badge>
        );

      case 'date':
        return value ? dayjs(value as string).format('YYYY-MM-DD HH:mm') : '-';

      case 'number':
        return value !== undefined && value !== null ? Number(value).toFixed(2) : '-';

      default:
        return String(value || '-');
    }
  };

  return (
    <Drawer
      opened={opened}
      onClose={onClose}
      position="right"
      size="lg"
      title={
        <Text fw={500} size="lg">
          {title}
        </Text>
      }
    >
      {data && (
        <ScrollArea h="calc(100vh - 120px)" ref={scrollRef}>
          <Stack gap="md">
            <Paper shadow="xs" p="md" withBorder>
              <Stack gap="sm">
                {fields.map((field, index) => (
                  <div key={field.key}>
                    <Group justify="space-between">
                      <Text size="sm" c="dimmed">
                        {field.label}
                      </Text>
                      <Text size="sm">{renderField(field)}</Text>
                    </Group>
                    {index < fields.length - 1 && <Divider my="xs" />}
                  </div>
                ))}
              </Stack>
            </Paper>

            {timeline && timeline.length > 0 && (
              <Paper shadow="xs" p="md" withBorder>
                <Text size="sm" fw={500} mb="md">
                  操作记录
                </Text>
                <Timeline active={timeline.length - 1} bulletSize={20} lineWidth={2}>
                  {timeline.map((item, index) => (
                    <Timeline.Item
                      key={index}
                      bullet={<IconClock size={12} />}
                      title={item.action}
                    >
                      <Text c="dimmed" size="xs">
                        <Group gap="xs">
                          <IconUser size={12} />
                          <span>{item.user}</span>
                          <IconCalendar size={12} />
                          <span>{dayjs(item.time).format('MM-DD HH:mm')}</span>
                        </Group>
                      </Text>
                      {item.description && (
                        <Text size="xs" mt={4}>
                          {item.description}
                        </Text>
                      )}
                    </Timeline.Item>
                  ))}
                </Timeline>
              </Paper>
            )}

            {actions && actions.length > 0 && (
              <Paper shadow="xs" p="md" withBorder>
                <Group grow>
                  {actions.map((action, index) => (
                    <Button
                      key={index}
                      color={action.color}
                      onClick={action.onClick}
                      disabled={action.disabled}
                    >
                      {action.label}
                    </Button>
                  ))}
                </Group>
              </Paper>
            )}
          </Stack>
        </ScrollArea>
      )}
    </Drawer>
  );
}
