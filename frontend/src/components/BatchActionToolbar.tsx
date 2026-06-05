import { useState } from 'react';
import {
  Paper,
  Group,
  Button,
  Menu,
  Checkbox,
  Text,
  ActionIcon,
  Badge,
  Modal,
  Stack,
  Textarea,
} from '@mantine/core';
import {
  IconCheckbox,
  IconTrash,
  IconArrowRight,
  IconDotsVertical,
  IconAlertCircle,
} from '@tabler/icons-react';

interface BatchAction {
  key: string;
  label: string;
  icon?: React.ReactNode;
  color?: string;
  requiresReason?: boolean;
  onExecute: (ids: number[], reason?: string) => void;
}

interface BatchActionToolbarProps<T> {
  items: T[];
  selectedIds: number[];
  onSelectionChange: (ids: number[]) => void;
  getId: (item: T) => number;
  actions: BatchAction[];
  statusOptions?: { value: string; label: string }[];
}

export function BatchActionToolbar<T>({
  items,
  selectedIds,
  onSelectionChange,
  getId,
  actions,
  statusOptions,
}: BatchActionToolbarProps<T>) {
  const [reasonModalOpen, setReasonModalOpen] = useState(false);
  const [pendingAction, setPendingAction] = useState<BatchAction | null>(null);
  const [reason, setReason] = useState('');

  const allSelected =
    items.length > 0 && selectedIds.length === items.length;
  const someSelected = selectedIds.length > 0;

  const handleSelectAll = () => {
    if (allSelected) {
      onSelectionChange([]);
    } else {
      onSelectionChange(items.map(getId));
    }
  };

  const handleAction = (action: BatchAction) => {
    if (action.requiresReason) {
      setPendingAction(action);
      setReasonModalOpen(true);
    } else {
      action.onExecute(selectedIds);
    }
  };

  const handleConfirmWithReason = () => {
    if (pendingAction) {
      pendingAction.onExecute(selectedIds, reason);
      setReasonModalOpen(false);
      setPendingAction(null);
      setReason('');
    }
  };

  return (
    <>
      <Paper shadow="sm" p="sm" withBorder>
        <Group justify="space-between">
          <Group>
            <Checkbox
              checked={allSelected}
              indeterminate={someSelected && !allSelected}
              onChange={handleSelectAll}
              label="全选"
            />
            {someSelected && (
              <Badge variant="light">
                已选择 {selectedIds.length} 项
              </Badge>
            )}
          </Group>

          {someSelected && (
            <Group gap="xs">
              {actions.slice(0, 3).map((action) => (
                <Button
                  key={action.key}
                  size="xs"
                  variant="light"
                  color={action.color}
                  leftSection={action.icon}
                  onClick={() => handleAction(action)}
                >
                  {action.label}
                </Button>
              ))}

              {actions.length > 3 && (
                <Menu position="bottom-end">
                  <Menu.Target>
                    <ActionIcon variant="subtle">
                      <IconDotsVertical size={16} />
                    </ActionIcon>
                  </Menu.Target>
                  <Menu.Dropdown>
                    {actions.slice(3).map((action) => (
                      <Menu.Item
                        key={action.key}
                        color={action.color}
                        leftSection={action.icon}
                        onClick={() => handleAction(action)}
                      >
                        {action.label}
                      </Menu.Item>
                    ))}
                  </Menu.Dropdown>
                </Menu>
              )}
            </Group>
          )}
        </Group>
      </Paper>

      <Modal
        opened={reasonModalOpen}
        onClose={() => setReasonModalOpen(false)}
        title="请填写原因"
      >
        <Stack>
          <Text size="sm" c="dimmed">
            执行 "{pendingAction?.label}" 操作需要填写原因
          </Text>
          <Textarea
            label="原因"
            placeholder="请输入原因..."
            value={reason}
            onChange={(e) => setReason(e.target.value)}
            minRows={3}
          />
          <Group justify="flex-end">
            <Button
              variant="subtle"
              onClick={() => setReasonModalOpen(false)}
            >
              取消
            </Button>
            <Button onClick={handleConfirmWithReason} disabled={!reason.trim()}>
              确认
            </Button>
          </Group>
        </Stack>
      </Modal>
    </>
  );
}
