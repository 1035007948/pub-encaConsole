import { useState, useEffect, useRef } from 'react';
import {
  Modal,
  Stack,
  TextInput,
  Group,
  Text,
  Badge,
  Kbd,
  ScrollArea,
  Paper,
} from '@mantine/core';
import { IconSearch, IconArrowRight } from '@tabler/icons-react';

interface CommandItem {
  id: string;
  label: string;
  category: string;
  shortcut?: string;
  icon?: React.ReactNode;
  action: () => void;
}

interface CommandPaletteProps {
  opened: boolean;
  onClose: () => void;
  commands: CommandItem[];
}

export function CommandPalette({
  opened,
  onClose,
  commands,
}: CommandPaletteProps) {
  const [query, setQuery] = useState('');
  const [selectedIndex, setSelectedIndex] = useState(0);
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (opened) {
      setQuery('');
      setSelectedIndex(0);
      setTimeout(() => inputRef.current?.focus(), 100);
    }
  }, [opened]);

  const filteredCommands = commands.filter(
    (cmd) =>
      cmd.label.toLowerCase().includes(query.toLowerCase()) ||
      cmd.category.toLowerCase().includes(query.toLowerCase())
  );

  const groupedCommands = filteredCommands.reduce(
    (acc, cmd) => {
      if (!acc[cmd.category]) {
        acc[cmd.category] = [];
      }
      acc[cmd.category].push(cmd);
      return acc;
    },
    {} as Record<string, CommandItem[]>
  );

  const handleKeyDown = (e: React.KeyboardEvent) => {
    switch (e.key) {
      case 'ArrowDown':
        e.preventDefault();
        setSelectedIndex((prev) =>
          Math.min(prev + 1, filteredCommands.length - 1)
        );
        break;
      case 'ArrowUp':
        e.preventDefault();
        setSelectedIndex((prev) => Math.max(prev - 1, 0));
        break;
      case 'Enter':
        e.preventDefault();
        if (filteredCommands[selectedIndex]) {
          filteredCommands[selectedIndex].action();
          onClose();
        }
        break;
      case 'Escape':
        onClose();
        break;
    }
  };

  const handleSelect = (command: CommandItem) => {
    command.action();
    onClose();
  };

  return (
    <Modal
      opened={opened}
      onClose={onClose}
      withCloseButton={false}
      size="lg"
      padding={0}
    >
      <Stack gap={0}>
        <TextInput
          ref={inputRef}
          placeholder="搜索命令或输入快捷键..."
          leftSection={<IconSearch size={16} />}
          value={query}
          onChange={(e) => {
            setQuery(e.target.value);
            setSelectedIndex(0);
          }}
          onKeyDown={handleKeyDown}
          size="lg"
          styles={{
            input: {
              border: 'none',
              borderBottom: '1px solid #e9ecef',
            },
          }}
        />

        <ScrollArea h={400}>
          <Stack gap={0} p="sm">
            {Object.entries(groupedCommands).map(([category, items]) => (
              <div key={category}>
                <Text size="xs" c="dimmed" mb="xs">
                  {category}
                </Text>
                <Stack gap={4}>
                  {items.map((item) => {
                    const globalIndex = filteredCommands.indexOf(item);
                    const isSelected = globalIndex === selectedIndex;

                    return (
                      <Paper
                        key={item.id}
                        p="xs"
                        withBorder
                        style={{
                          cursor: 'pointer',
                          backgroundColor: isSelected ? '#f8f9fa' : undefined,
                        }}
                        onClick={() => handleSelect(item)}
                        onMouseEnter={() => setSelectedIndex(globalIndex)}
                      >
                        <Group justify="space-between">
                          <Group gap="xs">
                            {item.icon}
                            <Text size="sm">{item.label}</Text>
                          </Group>
                          {item.shortcut && (
                            <Group gap={4}>
                              {item.shortcut.split('+').map((key, i) => (
                                <span key={i}>
                                  <Kbd size="xs">{key}</Kbd>
                                  {i < item.shortcut.split('+').length - 1 && (
                                    <span> + </span>
                                  )}
                                </span>
                              ))}
                            </Group>
                          )}
                        </Group>
                      </Paper>
                    );
                  })}
                </Stack>
              </div>
            ))}

            {filteredCommands.length === 0 && (
              <Text c="dimmed" ta="center" py="xl">
                未找到匹配的命令
              </Text>
            )}
          </Stack>
        </ScrollArea>

        <Group justify="space-between" p="sm" style={{ borderTop: '1px solid #e9ecef' }}>
          <Group gap="xs">
            <Kbd size="xs">↑</Kbd>
            <Kbd size="xs">↓</Kbd>
            <Text size="xs" c="dimmed">
              选择
            </Text>
          </Group>
          <Group gap="xs">
            <Kbd size="xs">Enter</Kbd>
            <Text size="xs" c="dimmed">
              执行
            </Text>
          </Group>
          <Group gap="xs">
            <Kbd size="xs">Esc</Kbd>
            <Text size="xs" c="dimmed">
              关闭
            </Text>
          </Group>
        </Group>
      </Stack>
    </Modal>
  );
}
